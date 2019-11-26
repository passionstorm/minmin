package models

import uuid "github.com/iris-contrib/go.uuid"

type Category struct {
	ID   uuid.UUID
	Name string
}
