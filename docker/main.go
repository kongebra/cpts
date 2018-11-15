package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(GetPort(), router))
}

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return ":" + port
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
