package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers" // this should be highlighted as not present in the vendor directory
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gorilla!\n"))
	})).ServeHTTP)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
