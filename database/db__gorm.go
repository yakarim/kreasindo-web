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
		db = b
	} else {
		psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
			"ec2-54-92-174-171.compute-1.amazonaws.com", 5432, "mqthtkprxhgcol", "cb6e1b8a1832614c8d86d28b7ae8e8128b382f4e112fc76be231fb61e9b9c0f0", "d2lvvitjtckt84")

		b, err := gorm.Open("postgres", psql)
		if err != nil {
			log.Println(err)
		}
		db = b
	}
	db.AutoMigrate(&User{})
	DB = db

}
