package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// GormConnect mysqlとの接続
func GormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error in %v\n", err)
	}

	DBMS := os.Getenv("DBMS")
	USER := "root"
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Fatalln(err)
	}
	return db
}
