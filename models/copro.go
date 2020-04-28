package models

import (
	"github.com/google/uuid"
)

//Copro model
type Copro struct {
	UUID uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
