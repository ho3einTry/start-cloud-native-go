package main

import (
	"fmt"
	"github/ho3eintry/start-cloud-native-go/api"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/book/", api.BookHandlerFunc)

	http.ListenAndServe(port(), nil)

}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 1 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go.")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
