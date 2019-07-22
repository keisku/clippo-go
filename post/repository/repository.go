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
	PASS := "Root0000"
	PROTOCOL := "tcp(clippo-rds.cpciso94q1yy.ap-northeast-1.rds.amazonaws.com:3306)"
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
