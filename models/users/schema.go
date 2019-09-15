package users

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Base struct {
	ID        uint64    `json:"id" gorm:"AUTO_INCREMENT"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	// Roles     []string
}
