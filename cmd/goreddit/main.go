package main

import (
	"log"
	"net/http"

	"github.com/goweb/goreddit/postgres"
	"github.com/goweb/goreddit/web"
)

func main() {
	// connect to the database
	store, err := postgres.NewStore("postgres://postgres:secret@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// intantiate the handlre
	h := web.NewHandler(store)

	// listen for requests
	http.ListenAndServe(":3000", h)
}
