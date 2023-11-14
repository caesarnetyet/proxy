package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *NoteServer) handleGetNote(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	note, err := s.store.findNote(id)

	if err != nil {
		http.Error(w, "Note was not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Could not encode Note", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (s *NoteServer) handlePostNote(w http.ResponseWriter, r *http.Request) {
	var noteRequest NoteStoreRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&noteRequest); err != nil {
		http.Error(w, fmt.Sprint("Request decode failed", err.Error()), http.StatusBadRequest)
		return
	}
	note, err := s.store.storeNote(noteRequest.Title, noteRequest.Content)

	if err != nil {
		http.Error(w, fmt.Sprint("Error while trying to save into store", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := NoteStoreResponseDTO{Message: "Note created successfully", Data: note}

	json.NewEncoder(w).Encode(response)
}
