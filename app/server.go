package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

// main sets up HTTP handlers for different endpoints and starts the server
// listening on port 8080. The server responds to three routes: "/segredo"
// and "/config" with the ConfigMap function, and "/" with the Hello function.
func main() {
	http.HandleFunc("/healthz", Heathz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/config", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

// Hello prints a friendly message with the name set on the environment variable
// NAME.
func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	w.Write([]byte("<h1>Alô, " + name + " é tão bom ouvir de novo a sua voz!</h1>"))
}

// ConfigMap reads the content of the file at /go/list/lista.txt and writes it
// to the HTTP response. If an error occurs while reading the file, it writes
// an error message to the response.
func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/list/lista.txt")
	if err != nil {
		fmt.Fprintf(w, "Error reading file: %s", err.Error())
	}
	fmt.Fprintf(w, "My Family: %s.", string(data))
}

// Secret prints a message with the content of the SECRET environment
// variable.
func Secret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I know your secret: %s.", os.Getenv("SECRET"))
}

func Heathz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)
	fmt.Fprintf(w, "Alive for %v", duration)
}
