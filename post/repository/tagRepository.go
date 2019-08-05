package repository

import (
	"log"

	"github.com/kskumgk63/clippo-go/post/entity"
)

// GetTag get a tag
func GetTag(name string) (*entity.Tag, error) {
	var tag entity.Tag
	db := gormConnect()
	defer db.Close()

	if err := db.Where("tag_name = ?", name).Find(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

// CreateTag create a new tag
func CreateTag(name string) *entity.Tag {
	var tag entity.Tag

	db := gormConnect()
	defer db.Close()

	tag.TagName = name

	if err := db.Create(&tag).Error; err != nil {
		return nil
	}
	if err := db.Where("tag_name = ?", name).Find(&tag).Error; err != nil {
		return &tag
	}
	return &tag
}

// GetPostIDByTagID get post_id by tag_id
func GetPostIDByTagID(id uint) *uint {
	var postsContactsTags entity.PostsContactsTags
	db := gormConnect()
	defer db.Close()
	if err := db.Where("tag_id = ?", id).Find(&postsContactsTags).Error; err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		return nil
	}
	return &postsContactsTags.PostID
}
