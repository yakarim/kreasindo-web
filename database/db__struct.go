package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Base ...
type Base struct {
	ID        string `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	return scope.SetColumn("ID", uuid.String())
}

// User struct.
type User struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
