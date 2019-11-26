package models

import (
	"database/sql"
	uuid "github.com/iris-contrib/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Auth struct {
	ID             uuid.UUID
	Token          string
	Username       string `json:"username"`
	Password       string
	RefreshToken   sql.NullString
	ExpiresIn      sql.NullTime
	HashedPassword []byte `json:"-"`
	Scopes         sql.NullString
	CreatedAt      time.Time
}

// GeneratePassword will generate a hashed password for us based on the
// user's input.
func  GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword will check if passwords are matched.
func ValidatePassword(password string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}
