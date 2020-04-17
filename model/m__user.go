package model

import (
	"errors"

	"github.com/yakarim/kreasindo-web/database"
)

// Query user.
func (m *Model) Query() ([]database.User, error) {
	var user []database.User
	if err := db.Order("created_at desc").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Create user.
func (m *Model) Create(user database.User) error {
	if !db.Where("email = ?", user.Email).First(&user).RecordNotFound() {
		return errors.New("EMAIL_FOUND")
	}
	user.Password = m.HashAndSalt(m.GetPwd(user.Password))
	if err := db.Model(&m.User).Create(&user).Error; err != nil {
		return err
	}
	return errors.New("SUKSES_EMAIL")
}
