package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const connectionString = "postgres://user:password@host:port/database"

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
	return Note{}, nil
}

func NewConnection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connectionString)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
