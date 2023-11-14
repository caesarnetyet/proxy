package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxyServer(t *testing.T) {
	store := StubNoteStore{}
	server := NewNoteServer(&store)
	t.Run("Running Server", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		got := response.Body.String()
		want := "Hello World!"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func assertStatus(t testing.TB, responseCode, expectedStatusCode int) {
	t.Helper()

	if responseCode != expectedStatusCode {
		t.Errorf("got status %d, expected status code %d", responseCode, expectedStatusCode)
	}

}
