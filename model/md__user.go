package model

import (
	"errors"

	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/database"
)

// User ...
type User struct {
	config.Config
	database.User
}

// Query user.
func (m *User) Query() ([]database.User, error) {
	var user []database.User
	if err := db.Order("created_at desc").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Create user.
func (m *User) Create(user database.User) error {
	if !db.Where("email = ?", user.Email).First(&user).RecordNotFound() {
		return errors.New("EMAIL_FOUND")
	}
	user.Password = m.HashAndSalt(m.GetPwd(user.Password))
	if err := db.Model(&m.User).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Update user.
func (m *User) Update(user database.User) error {
	if user.Password == "" {
		if err := db.Model(&m.User).Update(&user).Error; err != nil {
			return err
		}
	} else {
		user.Password = m.HashAndSalt(m.GetPwd(user.Password))
		if err := db.Model(&m.User).Update(&user).Error; err != nil {
			return err
		}
		return errors.New("SUKSES_EMAIL")
	}
	return nil
}

// Delete ...
func (m *User) Delete(id string) error {
	if err := db.Delete(&database.User{}, database.User{Base: database.Base{ID: id}}).Error; err != nil {
		return err
	}
	return errors.New("SUKSES_EMAIL")

}
