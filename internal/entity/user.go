package entity

import "github.com/gofrs/uuid"

type User struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Name string    `json:"name" binding:"required"`
}
