package models

import "github.com/google/uuid"


type Product struct {
	ID uuid.UUID
	Price int
	Name string
}