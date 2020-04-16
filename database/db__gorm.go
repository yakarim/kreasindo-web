package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// database ...
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB ... gorm
var DB *gorm.DB

func init() {
	var db *gorm.DB
	p := os.Getenv("PORT")
	if p == "8080" {
		psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			"localhost", 5432, "postgres", "1234", "kreasindo")

		b, err := gorm.Open("postgres", psql)
		if err != nil {
			log.Println(err)
		}
		db.AutoMigrate(&User{})
		db = b
	} else {

	}
	DB = db
}
