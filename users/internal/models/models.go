package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id       *uuid.UUID `json:"id"`
	Username string     `json:"username"`
	Password    string     `json:"password"`
}
