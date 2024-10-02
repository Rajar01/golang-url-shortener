package main

import (
	"errors"
	"fmt"
	"github.com/Rajar01/golang-url-shortener/src/database"
	"github.com/Rajar01/golang-url-shortener/src/models"
	"github.com/Rajar01/golang-url-shortener/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type originalURLRequestBody struct {
	OriginalUrl string `json:"original_url"`
}

// Function to initialize the database connection
func initDatabase() {
	if err := database.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return
	}

	log.Println("Connected to the database!")
}

// Function for database migration
func databaseMigration() {
	err := database.DB.AutoMigrate(&models.ShortLink{})
	if err != nil {
		log.Fatalf("failed to migrate database models: %v", err)
		return
	}

	log.Println("Database migration completed successfully.")
}

func main() {
	utils.InitEnv()

	baseUrl := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")

	initDatabase()
	databaseMigration()

	// Initialize the gin library
	r := gin.Default()

	// Http handler to create a new shortened link
	r.POST("/shorted-links", func(c *gin.Context) {
		var jsonData originalURLRequestBody

		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "Invalid request body. Please provide valid JSON data."},
			)

			log.Fatalf("Failed to bind JSON: %v", err)
			return
		}

		shortenedURL, err := utils.GenerateShortenedURL(fmt.Sprintf("%s:%s", baseUrl, port), 8)
		if err != nil {
			log.Printf("Error generating shortened URL: %v", err)
			return
		}

		database.DB.Create(&models.ShortLink{OriginalURL: jsonData.OriginalUrl, ShortenedURL: shortenedURL})

		c.JSON(http.StatusCreated, gin.H{})
	})

	// Http handler to delete a shortened link
	r.DELETE("/shorted-links/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := database.DB.First(&models.ShortLink{ID: uint(id)}).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}

		database.DB.Delete(&models.ShortLink{}, c.Param("id"))
		c.JSON(http.StatusOK, gin.H{})
	})

	// Http handler to update an original url for shortened link
	r.PUT("/shorted-links/:id", func(c *gin.Context) {
		var shortLink models.ShortLink

		// Checking if shortened link exist by id
		err := database.DB.First(&shortLink, c.Param("id")).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}

		// Update original url for shortened link
		var jsonData originalURLRequestBody

		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "Invalid request body. Please provide valid JSON data."},
			)

			log.Fatalf("Failed to bind JSON: %v", err)
			return
		}

		shortLink.OriginalURL = jsonData.OriginalUrl
		database.DB.Save(&shortLink)
		c.JSON(http.StatusOK, gin.H{})
	})

	// Http handler to get all shortened link
	r.GET("/shorted-links", func(c *gin.Context) {
		var shortLinks []models.ShortLink
		database.DB.Find(&shortLinks)
		c.JSON(http.StatusOK, gin.H{"short_links": shortLinks})
	})

	// Http handler to get specific shortened link by id
	r.GET("/shorted-links/:id", func(c *gin.Context) {
		var shortLink models.ShortLink
		err := database.DB.First(&shortLink, c.Param("id")).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"short_links": []models.ShortLink{}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"short_links": []models.ShortLink{shortLink}})
	})

	// Http handler to handle shortened url
	r.GET("/:shorted-link", func(c *gin.Context) {
		var shortLink models.ShortLink

		err := database.DB.Where("shortened_url = ?", fmt.Sprintf("%s:%s/%s", baseUrl, port, c.Param("shorted-link"))).First(&shortLink).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"short_links": []models.ShortLink{}})
			return
		}

		c.Redirect(http.StatusMovedPermanently, shortLink.OriginalURL)
	})

	// Listen and serve on 127.0.0.1:8080
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		return
	}
}
