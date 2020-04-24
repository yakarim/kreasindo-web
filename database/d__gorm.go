package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	// database ...
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB ... gorm
var DB *gorm.DB
var once sync.Once

func init() {

	const (
		host     = "ec2-52-71-231-180.compute-1.amazonaws.com"
		port     = 5432
		database = "d4hq86psqtvd0e"
		user     = "kwpicwgoeksjds"
		pass     = "1ff0cd6a150c7486649d7c8ef2f7f12a5671567c8631374e8c742b75aabbcecc"
		sslmode  = "require"
	)
	var db *gorm.DB
	p := os.Getenv("PORT")
	if p == "8080" {
		db = pqsl("localhost", "postgres", "1234", "kreasindo", "disable", 5432)
	} else {
		db = pqsl(host, user, pass, database, sslmode, port)
	}
	db.AutoMigrate(&User{}, &Contact{}, &Images{}, &Abouth{}, &Specialist{}, &Gallery{})
	DB = db

}

func pqsl(host, user, pass, database, sslmode string, port int) *gorm.DB {
	var bb *gorm.DB

	once.Do(func() {
		psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s ", host, port, user, pass, database, sslmode)

		b, err := gorm.Open("postgres", psql)
		if err != nil {
			log.Println(err)
		}
		bb = b
	})
	return bb

}
