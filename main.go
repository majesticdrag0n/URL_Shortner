package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Url struct {
	ID           string
	OrginalUrl   string
	ShortUrl     string
	CreationDate string
}

var urlDB = make(map[string]Url)

func generateShortUrl(orginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(orginalUrl))
	return hex.EncodeToString(hasher.Sum(nil))[:8]
}

func createUrl(orginalUrl string) Url {
	shortUrl := generateShortUrl(orginalUrl)
	url := Url{
		ID:           shortUrl,
		OrginalUrl:   orginalUrl,
		ShortUrl:     shortUrl,
		CreationDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	urlDB[shortUrl] = url
	return url
}

func getOriginalUrl(shortUrl string) (string, bool) {
	url, ok := urlDB[shortUrl]
	if !ok {
		return "", false
	}
	return url.OrginalUrl, true
}

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Url string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Println("Received URL:", data.Url)
	shortUrl_ := createUrl(data.Url)
	response := struct {
		ShortUrl string `json:"shortUrl"`
	}{
		ShortUrl: shortUrl_.ShortUrl,
	}
	json.NewEncoder(w).Encode(response)
}

func redirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[len("/redirect/"):]
	originalUrl, ok := getOriginalUrl(shortUrl)
	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalUrl, http.StatusMovedPermanently)
}

func main() {
	// fmt.Println(generateShortUrl("https://www.example.com"), "url shortner")
	// fmt.Println(createUrl("https://www.example.com"))

	// originalUrl, ok := getOriginalUrl("https://www.example.com")
	// if ok {
	// 	fmt.Println(originalUrl)
	// } else {
	// 	fmt.Println("Url not found")
	// }
	// fmt.Println(urlDB)

	http.HandleFunc("/shorten", ShortUrlHandler)
	http.HandleFunc("/redirect/", redirectUrlHandler)

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
