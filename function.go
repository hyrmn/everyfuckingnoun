package everyfuckingnoun

import (
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	_ "embed"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

//go:embed templates/index.html
var templateContent string

var indexTmpl = template.Must(template.New("Index").Parse(templateContent))

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
