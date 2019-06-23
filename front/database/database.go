package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "ao6415012"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME   := "clippo"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
	  panic(err.Error())
	}
	return db
}
