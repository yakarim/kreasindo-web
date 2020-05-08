package model

import (
	"github.com/yakarim/kreasindo-web/config"
	"github.com/yakarim/kreasindo-web/database"
)

// Gallery model.
type Gallery struct {
	config.Config
}

// Query gallery.
func (g *Gallery) Query() ([]database.Gallery, error) {
	var glr []database.Gallery
	if err := db.Order("created_at desc").Find(&glr).Error; err != nil {
		return glr, err
	}
	return glr, nil
}

// Create user.
