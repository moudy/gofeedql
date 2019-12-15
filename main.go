package main

import (
	"log"
	"net/http"
	"os"
)

var port = "3000"

func main() {
	val, ok := os.LookupEnv("PORT")

	if ok {
		port = val
	}

	mux := http.NewServeMux()
	mux.Handle("/graphql", SchemaHandler())

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
