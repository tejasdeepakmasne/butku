package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tejasdeepakmasne/butku/resolver"
	"github.com/tejasdeepakmasne/butku/shortner"
)

func main() {
	loadenv()
	StartServer()

}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortner.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", resolver.ResolveURL).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error Loading env file %v", err)
	}
}
