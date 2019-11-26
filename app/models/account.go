package models

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type Account struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
