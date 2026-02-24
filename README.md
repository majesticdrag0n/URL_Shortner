🔗 URL Shortener — Go (Golang)

A simple and efficient URL shortener service built using Go (Golang).
This application converts long URLs into short, shareable links and redirects users to the original URL when accessed.

🚀 Features

Shorten long URLs

Redirect to original URLs

Unique short code generation

Fast and lightweight Go backend

REST API support

Simple and scalable architecture

Easy deployment

🏗️ Tech Stack

Language: Go (Golang)

Framework: net/http 

Database: In-memory 

Architecture: REST API

⚙️ Installation & Setup
1. Clone the repository
git clone https://https://github.com/majesticdrag0n/URL_Shortner.git
cd url-shortener
2. Install dependencies
go mod tidy
3. Run the application
go run main.go

Server runs on:

http://localhost:8080
📡 API Endpoints
➜ Shorten URL

POST /shorten

Request Body:

{
  "url": "https://example.com"
}

Response:

{
  "short_url": "abc123"
}
➜ Redirect to Original URL
"http://localhost:8080/redirect/abc123"

Example:

http://localhost:8080/abc123
🧠 How It Works

User submits a long URL.

System generates a unique short code.

Mapping between short code and original URL is stored.

When short URL is accessed, user is redirected.

🔐 Future Improvements

User authentication

URL expiration

Custom short links

Analytics and tracking

Rate limiting

Distributed storage

Caching with Redis

🧪 Testing

Run tests using:

go test ./...
🤝 Contributing

Contributions are welcome.

Fork the repository

Create a new branch

Commit changes

Submit a pull request
