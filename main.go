package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	v := os.Getenv("APP_VERSION")
	if v == "" {
		v = "v1"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from %s\n", v)
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	http.ListenAndServe(":8080", nil)
}
