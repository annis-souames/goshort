package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/annis-souames/goshort/utils"
)

var ctx = context.Background()

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(writer, nil)
	})
	http.HandleFunc("/shorten", func(writer http.ResponseWriter, req *http.Request) {
		// Get the URL to shorten from the request
		url := req.URL.Query().Get("url")
		// Shorten the URL
		shortURL := utils.GetShortCode()
		fullShortURL := fmt.Sprintf("http://localhost:8080/r/%s", shortURL)

		// Generated short URL
		fmt.Printf("Generated short URL: %s\n", shortURL) // Log to console

		// Connect to redis server
		dbClient := utils.NewRedisClient()

		// Set the key in Redis
		utils.SetKey(&ctx, dbClient, shortURL, url, 0)

		// Ideally, we would return some html tags
		fmt.Fprintf(writer,
			`<p class="mt-4 text-green-600">Shortened URL: <a href="/r/%s" class="underline">%s</a></p>`, shortURL, fullShortURL)
	})

	// This handler will redirect the user to the long URL based on the short url
	http.HandleFunc("/r", func(writer http.ResponseWriter, req *http.Request) {
		// Handle redirect
	})
	http.ListenAndServe(":8080", nil)
}
