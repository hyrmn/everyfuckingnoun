package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var lines []string
var lineCount int

func main() {
	content, err := ioutil.ReadFile("./data/nouns.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines = strings.Split(string(content), "\n")
	lineCount = len(lines)

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

	randomIdx := rand.Intn(lineCount)
	noun := lines[randomIdx]
	fmt.Fprint(w, "Fuck ", noun)
}
