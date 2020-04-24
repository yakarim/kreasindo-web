package model

import "github.com/yakarim/kreasindo-web/database"

// Abouth ...
type Abouth struct {
	database.Abouth
}

// Query contact.
func (c *Abouth) Query() (Abouth, error) {
	var abouth Abouth
	if err := db.Preload("Image").First(&abouth).Error; err != nil {
		return abouth, err
	}
	return abouth, nil
}

// Create contact.
func (c *Abouth) Create(a database.Abouth) error {
	if err := db.Model(&c.Abouth).Save(&a).Error; err != nil {
		return err
	}
	return nil
}
