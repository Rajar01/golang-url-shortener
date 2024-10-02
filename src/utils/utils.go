package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
)

func GenerateRandomString(length int) (string, error) {
	// Create a byte slice of the desired length
	randomBytes := make([]byte, length)

	// Fill the byte slice with random data
	if _, err := io.ReadFull(rand.Reader, randomBytes); err != nil {
		return "", err
	}

	// Encode the random bytes as a URL-safe base64 string
	return base64.URLEncoding.EncodeToString(randomBytes)[:length], nil
}

func GenerateShortenedURL(baseURL string, length int) (string, error) {
	randomString, err := GenerateRandomString(length)
	if err != nil {
		return "", err
	}

	// Concatenate the base URL with the random string
	return fmt.Sprintf("%s/%s", baseURL, randomString), nil
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
