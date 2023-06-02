//go:build ignore

package main

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"text/template"
	"time"
)

func main() {
	lineBreak := "\n"

	if runtime.GOOS == "windows" {
		lineBreak = "\r\n"
	}

	content, err := ioutil.ReadFile("./data/nouns.txt")
	if err != nil {
		log.Fatal(err)
	}

	nouns := strings.Split(string(content), lineBreak)

	f, err := os.Create("dictionary.go")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	packageTemplate.Execute(f, struct {
		Timestamp time.Time
		WordCount int
		Nouns     []string
	}{
		Timestamp: time.Now(),
		WordCount: len(nouns),
		Nouns:     nouns,
	})
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package everyfuckingnoun

var WordCount = {{ .WordCount }}
var Words = []string{
{{- range .Nouns }}
	{{ printf "%q" . }},
{{- end }}
}
`))