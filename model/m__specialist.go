package model

import (
	"errors"

	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/database"
)

type Specialist struct {
	config.Config
	database.Specialist
}

// Query specialist.
func (m *Specialist) Query() ([]database.Specialist, error) {
	var sp []database.Specialist
	if err := db.Order("created_at desc").Find(&sp).Error; err != nil {
		return sp, err
	}
	return sp, nil
}

// Create user.
func (m *Specialist) Create(sp database.Specialist) error {
	if !db.Where("title = ?", sp.Title).First(&m.Specialist).RecordNotFound() {
		return errors.New("TITLE_FOUND")
	}

	if err := db.Model(&m.Specialist).Create(&sp).Error; err != nil {
		return err
	}
	return nil
}

// Update user.
func (m *Specialist) Update(sp database.Specialist) error {
	if !db.Where("title = ?", sp.Title).First(&m.Specialist).RecordNotFound() {
		sp.Title = m.Specialist.Title
		if err := db.Model(&m.Specialist).Update(&sp).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&m.Specialist).Update(&sp).Error; err != nil {
		return err
	}
	return nil
}

// Delete specilaist...
func (m *Specialist) Delete(id string) error {
	if err := db.Delete(&database.Specialist{}, database.Specialist{Base: database.Base{ID: id}}).Error; err != nil {
		return err
	}
	return nil
}
