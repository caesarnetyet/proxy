package main

import (
	"net/http"
)

type NoteStore interface {
	findNote(id int) (Note, error)
	storeNote(title string, content string) (Note, error)
}

type NoteServer struct {
	store NoteStore
	http.Handler
}

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewNoteServer(store NoteStore) *NoteServer {
	server := new(NoteServer)
	server.store = store
	server.Handler = server.getRoutes()

	return server
}
