package models

import (
	"database/sql"
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type User struct {
	ID                 uuid.UUID
	Email              string
	Role               string
	FirstName          string
	LastName           string
	AccountId          uuid.UUID
	Phone              sql.NullString
	CreatedAt          time.Time
	UpdatedAt          time.Time
	EncryptedPassword  sql.NullString
	ResetPasswordToken sql.NullString
	ConfirmedAt        sql.NullTime
}
