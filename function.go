package everyfuckingnoun

import (
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

var indexTmpl = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	functions.HTTP("Index", index)
}

func index(w http.ResponseWriter, r *http.Request) {
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
