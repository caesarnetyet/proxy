package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type NoteServer struct {
	http.Handler
}

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewNoteServer() *NoteServer {
	server := new(NoteServer)
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	router.Get("/notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		note := Note{1, "TODO", "Im working on these Notes"}

		json.NewEncoder(w).Encode(note)
	})

	server.Handler = router

	return server
}
