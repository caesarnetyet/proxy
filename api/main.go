package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	store, err := NewPGNoteStore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	server := NewNoteServer(store)

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatal("Could not initialize server", err)
	}
}
