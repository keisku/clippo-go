package repository

import (
	"log"

	"github.com/kskumgk63/clippo-go/post/entity"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func unique(array []*uint) []*uint {
	m := map[*uint]bool{}
	uniq := []*uint{}

	for _, element := range array {
		if !m[element] {
			m[element] = true
			uniq = append(uniq, element)
		}
	}
	return uniq
}

func queryText(words []string) string {
	var query string
	for i, word := range words {
		if i == len(words)-1 {
			w := "%" + word + "%"
			query += w
		} else {
			w := "%" + word
			query += w
		}
	}
	return query
}

// Create create new post
func Create(post *entity.Post) error {
	// connect with db
	db := gormConnect()
	defer db.Close()

	if err := db.Create(&post).Error; err != nil {
		return err
	}

	return nil
}

// Delete delete a post
func Delete(id uint) error {
	var post entity.Post
	post.ID = id
	// connect with DB
	db := gormConnect()
	defer db.Close()
	if err := db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}

// GetByUserID get posts by user_id
func GetByUserID(userID string) []*entity.Post {
	// connect with DB
	db := gormConnect()
	defer db.Close()

	// get posts related to tags
	posts := []*entity.Post{}
	err := db.Preload("Tags").Order("ID desc").Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		return nil
	}
	return posts
}

// Search search posts by title or tags
func Search(how, userID string, words []string) []*entity.Post {
	posts := []*entity.Post{}

	// connect with DB
	db := gormConnect()
	defer db.Close()

	// search by title
	if how == "title" {
		err := db.Where("user_id = ?", userID).Where("title LIKE ?", queryText(words)).Find(&posts).Error
		// if not found any posts in DB, return sample
		if posts == nil {
			log.SetFlags(log.Lshortfile)
			log.Println(err)
			return nil
		}
		return posts
	}

	// search by tag
	if how == "tag" {
		var postID *uint
		var postIDs []*uint

		// search tags by tag_name
		for _, word := range words {
			tag, err := GetTag(word)
			if err != nil {
				log.SetFlags(log.Lshortfile)
				log.Println(err)
			} else {
				postID = GetPostIDByTagID(tag.ID)
			}
			postIDs = append(postIDs, postID)
		}

		// delete repwated post_id
		uniqPostIDs := unique(postIDs)

		// search posts by user_id and post_id
		for _, postID := range uniqPostIDs {
			// preload tags because there is empty tags unless preload
			db.Preload("Tags").Where("id = ? AND user_id = ?", postID, userID).Find(&posts)
			if posts == nil {
				log.Println("Not found any posts")
				return nil
			}
		}
		return posts
	}
	return posts
}
