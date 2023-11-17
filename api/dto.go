package main

import "github.com/go-playground/validator/v10"

type NoteStoreRequestDTO struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type GetNotesResponseDTO struct {
	Message string `json:"message"`
	Data    []Note `json:"data"`
}

type NoteStoreResponseDTO struct {
	Message string `json:"message"`
	Data    Note   `json:"data"`
}

func ValidateNoteStoreRequest(dto NoteStoreRequestDTO) error {
	validate := validator.New()
	return validate.Struct(dto)
}
