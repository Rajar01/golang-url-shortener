# Golang URL Shortener

This project is a simple URL shortener built using Golang, Gin, and GORM. It provides RESTful APIs to create, update, delete, and retrieve shortened URLs. The project uses MySQL (or any GORM-supported database) to store the original and shortened URLs.

## Features

- Shorten a URL with a randomly generated code.
- Retrieve all shortened URLs.
- Update or delete existing shortened URLs.
- Redirect to the original URL using the shortened URL.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Golang](https://golang.org/)
- [MySQL](https://www.mysql.com/) (or any GORM-compatible database)
- [Git](https://git-scm.com/)
- [Postman](https://www.postman.com/) or any API testing tool (for testing the API)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Rajar01/golang-url-shortener.git
    cd golang-url-shortener
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run database migrations and start the server:
   This will create the necessary tables in your database.
    ```bash
    go run main.go
    ```
   
4. The server will be running at `http://127.0.0.1:8080`.

## API Endpoints

### 1. Shorten a URL
- **Endpoint:** `/shorted-links`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "original_url": "https://example.com"
  }
  ```
- **Response:**
  ```json
  {}
  ```

### 2. Retrieve All Shortened URLs
- **Endpoint:** `/shorted-links`
- **Method:** `GET`
- **Response:**
  ```json
  {
    "short_links": [
      {
        "ID": 1,
        "OriginalURL": "https://example.com",
        "ShortenedURL": "http://localhost:8080/abc123",
        "CreatedAt": "2024-10-02T16:17:46.268+08:00",
        "UpdatedAt": "2024-10-02T16:20:18.851+08:00"
      }
    ]
  }
  ```

### 3. Retrieve a Shortened URL by ID
- **Endpoint:** `/shorted-links/:id`
- **Method:** `GET`
- **Response:**
  ```json
  {
    "short_links": [
      {
        "ID": 1,
        "OriginalURL": "https://example.com",
        "ShortenedURL": "http://localhost:8080/abc123",
        "CreatedAt": "2024-10-02T16:17:46.268+08:00",
        "UpdatedAt": "2024-10-02T16:20:18.851+08:00"
      }
    ]
  }
  ```

### 4. Update a Shortened URL
- **Endpoint:** `/shorted-links/:id`
- **Method:** `PUT`
- **Request Body:**
  ```json
  {
    "original_url": "https://newexample.com"
  }
  ```
- **Response:** `200 OK`

### 5. Delete a Shortened URL
- **Endpoint:** `/shorted-links/:id`
- **Method:** `DELETE`
- **Response:** `200 OK`

### 6. Redirect to Original URL
- **Endpoint:** `/:shorted-link`
- **Method:** `GET`
- **Response:** HTTP 301 Redirect to the original URL.

## Database

This project uses GORM for database operations. You can change the database connection in the `database.Connect()` function to use any other database supported by GORM, such as PostgreSQL or SQLite.

## Usage

Use Postman or Curl to test the API endpoints listed above. You can shorten URLs and retrieve the shortened versions using the API, then test the redirect functionality by accessing the shortened URL in your browser.

## GitHub Repository

For more details and to contribute, visit the repository: [Golang URL Shortener](https://github.com/Rajar01/golang-url-shortener).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Author

This project is maintained by [Rajar01](https://github.com/Rajar01).
