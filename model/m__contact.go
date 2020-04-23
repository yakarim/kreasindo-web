package model

import "github.com/yakarim/kreasindo-web/database"

// Contact ...
type Contact struct {
	database.Contact
}

// Query contact.
func (c *Contact) Query() (Contact, error) {
	var contact Contact
	if err := db.Preload("Image").First(&contact).Error; err != nil {
		return contact, err
	}
	return contact, nil
}

// Create contact.
func (c *Contact) Create(contact Contact) error {
	if err := db.Model(&c.Contact).Create(&contact).Error; err != nil {
		return err
	}
	return nil
}

// Update contact.
func (c *Contact) Update(contact Contact) error {
	if err := db.Model(&c.Contact).Update(&contact).Error; err != nil {
		return err
	}
	return nil
}
