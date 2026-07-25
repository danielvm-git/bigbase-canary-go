// story: e01s01
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	v, err := os.ReadFile("VERSION")
	version := "unknown"
	if err == nil {
		version = strings.TrimSpace(string(v))
	}
	_, _ = fmt.Fprintf(w, "<h1>bigbase canary (Go)</h1><footer>v%s</footer>", version)
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
