package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubNoteStore struct {
	store map[int]Note
}

func (s *StubNoteStore) findNote(id int) (Note, error) {
	note, ok := s.store[id]
	if ok {
		return note, nil
	}
	return Note{}, errors.New("Could not find Note by given ID")
}

func (s *StubNoteStore) storeNote(title string, content string) (Note, error) {
	incrementalID := len(s.store)
	incrementalID++

	newNote := Note{incrementalID, title, content}
	s.store[incrementalID] = newNote

	return s.store[incrementalID], nil
}

func TestNoteHandlers(t *testing.T) {
	store := StubNoteStore{map[int]Note{
		1:  {ID: 1, Title: "TODO", Content: "Im working on these Notes"},
		23: {ID: 23, Title: "Not me", Content: "This is not me."},
	}}

	server := NewNoteServer(&store)

	t.Run("Get Note with Id 1", func(t *testing.T) {
		request := newGETNoteRequest(http.MethodGet, 1, nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		var got Note

		decodeNote(t, response.Body, &got)

		want := Note{ID: 1, Title: "TODO", Content: "Im working on these Notes"}

		assertNote(t, got, want)

	})
	t.Run("Get note with ID 23", func(t *testing.T) {
		request := newGETNoteRequest(http.MethodGet, 23, nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var got Note
		decodeNote(t, response.Body, &got)

		want := Note{ID: 23, Title: "Not me", Content: "This is not me."}

		assertNote(t, got, want)
	})
	t.Run("Store new Note", func(t *testing.T) {
		noteDTO := NoteStoreRequestDTO{Title: "New Note", Content: "New content"}

		requestBody := new(bytes.Buffer)

		if err := json.NewEncoder(requestBody).Encode(noteDTO); err != nil {
			t.Error("Could not encode Note", err)
		}
		request, _ := http.NewRequest(http.MethodPost, "/notes", requestBody)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusCreated)

		var got NoteStoreResponseDTO
		decodeNoteResponseDTO(t, response.Body, &got)

		want := Note{len(store.store), noteDTO.Title, noteDTO.Content}

		assertNote(t, got.Data, want)
	})

}

func decodeNote(t testing.TB, body io.Reader, note *Note) {
	t.Helper()

	if err := json.NewDecoder(body).Decode(note); err != nil {
		t.Error("Error trying to decode Note", err.Error())
	}

}

func newGETNoteRequest(method string, id int, body io.Reader) *http.Request {
	request, _ := http.NewRequest(method, fmt.Sprintf("/notes/%d", id), body)
	return request
}

func decodeNoteResponseDTO(t testing.TB, body io.Reader, response *NoteStoreResponseDTO) {
	t.Helper()

	if err := json.NewDecoder(body).Decode(response); err != nil {
		t.Error("Error trying to decode Note Response", err.Error())
	}

}

func assertNote(t testing.TB, got, want Note) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got Note %v, want %v", got, want)

	}

}
