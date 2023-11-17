package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *NoteServer) getRoutes() http.Handler {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(contentTypeJSONMiddleware)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler((cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})))

	router.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	router.Get("/api/notes", s.handleGetNotes)
	router.Get("/api/notes/{id:[1-9]+}", s.handleGetNote)
	router.Post("/api/notes", s.handlePostNote)

	return router

}
