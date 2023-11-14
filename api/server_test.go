package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxyServer(t *testing.T) {
	server := NewNoteServer()
	t.Run("Running Server", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		code := response.Code
		assertStatus(t, code, http.StatusOK)

		got := response.Body.String()
		want := "Hello World!"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}

}
