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

// Specialist ...
type Specialist struct {
	Base
	Title string `json:"title"`
	Image Images `gorm:"foreignkey:uid;association_foreignkey:id"`
	Desc  string `json:"desc"`
}

// Gallery ...
type Gallery struct {
	Base
	Name  string `json:"name"`
	Image Images `gorm:"foreignkey:uid;association_foreignkey:id"`
	Desc  string `json:"desc"`
}

// Images ...
type Images struct {
	UID  string `gorm:"primary_key"`
	Name string `json:"name"`
	Body []byte `json:"body"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}

// Abouth ...
type Abouth struct {
	Base
	Title string `json:"title"`
	Image Images `gorm:"foreignkey:uid;association_foreignkey:id"`
	Desc  string `json:"desc"`
}

// Contact ...
type Contact struct {
	Base
	Alamat string `json:"alamat"`
	Image  Images `gorm:"foreignkey:uid;association_foreignkey:id"`
	Telp   string `json:"telp"`
}

// Tentang ...
type Tentang struct {
	Base
	Title string `json:"title"`
	Image Images `gorm:"foreignkey:uid;association_foreignkey:id"`
	Desc  string `json:"desc"`
}
