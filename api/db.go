package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type PGNoteStore struct {
	db *pgx.Conn
}

func NewPGNoteStore(ctx context.Context) (*PGNoteStore, error) {
	conn, err := NewConnection(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	return &PGNoteStore{conn}, nil
}

func (s *PGNoteStore) findNote(ctx context.Context, id int) (Note, error) {
	var note Note
	query := "SELECT id, title, content FROM notes WHERE id = $1"

	err := s.db.QueryRow(ctx, query, id).Scan(&note.ID, &note.Title, &note.Content)

	if err != nil {
		return Note{}, err
	}
	return note, nil
}

func (s *PGNoteStore) storeNote(ctx context.Context, title, content string) (Note, error) {
	var note Note
	query := "INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id, title, content"

	err := s.db.QueryRow(ctx, query, title, content).Scan(&note.ID, &note.Title, &note.Content)
	if err != nil {
		return Note{}, err
	}

	return note, nil
}

func (s *PGNoteStore) allNotes(ctx context.Context) ([]Note, error) {
	var notes []Note
	query := "SELECT id, title, content FROM notes"

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil

}

func NewConnection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_CONNECTION_STRING"))

	if err != nil {
		return nil, err
	}

	return conn, nil
}
