package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id       *uuid.UUID `json:"id"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	// Ajouter d'autres champs selon les besoins
}