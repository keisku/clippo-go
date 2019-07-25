package repository

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/kskumgk63/clippo-go/front/entity"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GormConnect mysqlとの接続
func GormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatalln(err)
		return nil
	}
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("USERNAME")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Printf("DBMS = %v\n", DBMS)
		log.Printf("CONNECT = %v\n", CONNECT)
		log.SetFlags(log.Lshortfile)
		log.Fatalln(err)
		return nil
	}
	return db
}

// CreateTable テーブル作成
func CreateTable() {
	db := GormConnect()
	if db.HasTable("users") {
		db.DropTable("users")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.User{})
	} else {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.User{})
	}
	if db.HasTable("posts") {
		db.DropTable("posts")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Post{})
	} else {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Post{})
	}
}
