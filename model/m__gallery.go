package model

import (
	"errors"

	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/protobuf"
)

// Gallery model.
type Gallery struct {
	config.Config
}

// Query gallery.
func (g *Gallery) Query() (*protobuf.Galleries, error) {
	var glr []*protobuf.Gallery

	if err := db.Find(&glr).Error; err != nil {
		return nil, err
	}
	data := protobuf.Galleries{
		Galleries: glr,
	}
	return &data, nil
}

// Create user.
func (g *Gallery) Create(sp *protobuf.Gallery) error {
	var glr *protobuf.Gallery
	if !db.Where("name = ?", sp.Name).First(&glr).RecordNotFound() {
		return errors.New("TITLE_FOUND")
	}

	if err := db.Model(&glr).Create(&sp).Error; err != nil {
		return err
	}
	return nil
}
