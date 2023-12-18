package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	ID          *uuid.UUID `json:"id"`
	Title       string     `json:"title"`
	Artist      string     `json:"artist"`
	Album       string     `json:"album"`
	ReleaseYear int        `json:"release_year"`
}
