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

const (
	host     = "ec2-52-71-231-180.compute-1.amazonaws.com"
	port     = 5432
	database = "d4hq86psqtvd0e"
	user     = "kwpicwgoeksjds"
	pass     = "1ff0cd6a150c7486649d7c8ef2f7f12a5671567c8631374e8c742b75aabbcecc"
)

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
			host, port, user, pass, database)

		b, err := gorm.Open("postgres", psql)
		if err != nil {
			log.Println(err)
		}
		db = b
	}
	db.AutoMigrate(&User{})
	DB = db

}
