package main

type NoteStoreRequestDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteStoreResponseDTO struct {
	Message string `json:"message"`
	Data    Note   `json:"Message"`
}
