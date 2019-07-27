package repository

import (
	"log"
	"os"

	"github.com/kskumgk63/clippo-go/post/entity"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gormConnect mysqlとの接続
func gormConnect() *gorm.DB {
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

// CreatePostsAndTagsTable create posts and tags table
func CreatePostsAndTagsTable() {
	db := gormConnect()
	if db.HasTable("posts") {
		log.Println("*** REcreate POSTS table ***")
		db.DropTable("posts")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Post{})
	} else {
		log.Println("*** Create POSTS table ***")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Post{})
	}
	if db.HasTable("tags") {
		log.Println("*** REcreate TAGS table ***")
		db.DropTable("tags")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Tag{})
	} else {
		log.Println("*** Create TAGS table ***")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Tag{})
	}
}
