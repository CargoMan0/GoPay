package entity

import "github.com/google/uuid"

type Event struct {
	ID   uuid.UUID
	Type string
}
