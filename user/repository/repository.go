package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GormConnect mysqlとの接続
func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "ao6415012"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "clippo"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Printf("DBMS = %v\n", DBMS)
		log.Printf("CONNECT = %v\n", CONNECT)
		log.SetFlags(log.Lshortfile)
		log.Fatalln(err)
	}
	return db
}
