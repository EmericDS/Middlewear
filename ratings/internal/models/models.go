package models

import (
    "github.com/gofrs/uuid"
)

type Rating struct {
    ID       *uuid.UUID `json:"id"`
    UserID   *uuid.UUID `json:"userId"`
    SongID   *uuid.UUID `json:"songId"`
    Rating   int        `json:"rating"`
    Comments string     `json:"comments"`
}
