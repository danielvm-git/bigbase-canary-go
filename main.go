// story: e01s01
package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed index.html
var indexHTML string

// indexTmpl uses [[ ]] delimiters so the {{ }} in embedded GitHub Actions
// workflow snippets are passed through to the browser verbatim.
var indexTmpl = template.Must(
	template.New("index").Delims("[[", "]]").Parse(indexHTML),
)

type pageData struct {
	Version string
}

func handler(w http.ResponseWriter, r *http.Request) {
	v, err := os.ReadFile("VERSION")
	version := "unknown"
	if err == nil {
		version = strings.TrimSpace(string(v))
	} else {
		log.Printf("VERSION read error: %v", err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := indexTmpl.Execute(w, pageData{Version: version}); err != nil {
		log.Printf("template render error: %v", err)
	}
}

func listenAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listenAddr(), nil))
}
