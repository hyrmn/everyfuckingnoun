package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//go:generate go run gen.go

var indexTmpl = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8181"
		log.Printf("Defaulting to port %s", port)
	}

	http.HandleFunc("/", indexHandler)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	randomIdx := rand.Intn(WordCount)
	noun := Words[randomIdx]

	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	if err := indexTmpl.Execute(w, noun); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
