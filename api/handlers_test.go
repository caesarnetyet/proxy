package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNoteHandlers(t *testing.T) {
	server := NewNoteServer()

	t.Run("Get Note with Id 1", func(t *testing.T) {
		request := newNoteRequest(t, http.MethodGet, 1, nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		var got Note

		decodeNote(t, response.Body, &got)

		want := Note{ID: 1, Title: "TODO", Content: "Im working on these Notes"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got Note %v, want %v", got, want)
		}
	})
	t.Run("Get note with ID 23", func(t *testing.T) {
		request := newNoteRequest(t, http.MethodGet, 23, nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

	})

}

func decodeNote(t testing.TB, body io.Reader, note *Note) {
	t.Helper()

	if err := json.NewDecoder(body).Decode(note); err != nil {
		t.Fatal("Error trying to decode Note", err.Error())
	}

}

func newNoteRequest(t testing.TB, method string, id int, body io.Reader) *http.Request {
	t.Helper()
	request, err := http.NewRequest(method, fmt.Sprintf("/notes/%d", id), body)
	if err != nil {
		t.Fatalf("Error on request: %q", err.Error())
	}

	return request

}
