package repository

import (
	"log"
	"strconv"

	"github.com/kskumgk63/clippo-go/post/entity"
	"github.com/kskumgk63/clippo-go/post/postpb"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func makeSamplePost() *postpb.Post {
	var tagArray []string
	tagArray = append(tagArray, "sample")
	return &postpb.Post{
		Id:          "xxxx",
		Url:         "http://localhost:8080/",
		Title:       "まだ投稿されていないようなので、記事をクリップしてみてください",
		Description: "150文字以内で記事の簡単なサマリーを書いてください。この記事は何を目的としているか、ジャンルは何かひと目でわかるようになっています。できるだけシンプルにサマリーを書くことをおすすめします。",
		Image:       "http://designers-tips.com/wp-content/uploads/2015/03/paper-clip6.jpg",
		Tag:         tagArray,
		UserId:      "0000",
	}
}

func convertPost(post *entity.Post) *postpb.Post {
	// convert uint to string
	post64 := uint64(post.ID)
	postID := strconv.FormatUint(post64, 10)

	// convert uint to string
	user64 := uint64(post.UserID)
	userID := strconv.FormatUint(user64, 10)

	var tagArray []string
	tags := post.Tags

	// convert struct to []string for the view
	for _, tag := range tags {
		tagArray = append(tagArray, tag.TagName)
	}

	return &postpb.Post{
		Id:          postID,
		Url:         post.URL,
		Title:       post.Title,
		Description: post.Description,
		Image:       post.Image,
		Tag:         tagArray,
		UserId:      userID,
	}
}

// Create create new post
func Create(req *postpb.CreatePostRequest) error {
	var tags []entity.Tag

	// connect with db
	db := gormConnect()
	defer db.Close()

	// convert string to uint
	resID := req.GetPost().GetUserId()
	id64, _ := strconv.ParseUint(resID, 10, 64)
	id := uint(id64)

	tagNames := req.GetPost().GetTag()

	// array to string
	for _, tagName := range tagNames {
		tag := entity.Tag{}
		// check if the tag_name is existed, if not create new
		if err := db.Where("tag_name = ?", tagName).Find(&tag).Error; err != nil {
			// save tag in DB
			tag.TagName = tagName
			db.Create(&tag)
		}
		tags = append(tags, tag)
	}

	post := entity.Post{
		URL:         req.GetPost().GetUrl(),
		Title:       req.GetPost().GetTitle(),
		Description: req.GetPost().GetDescription(),
		Image:       req.GetPost().GetImage(),
		Tags:        tags,
		UserID:      id,
	}

	if err := db.Create(&post).Error; err != nil {
		return err
	}

	return nil
}

// Delete delete a post
func Delete(req *postpb.DeletePostRequest) error {
	var post entity.Post
	id64, _ := strconv.ParseUint(req.GetId(), 10, 64)
	post.ID = uint(id64)
	// connect with DB
	db := gormConnect()
	defer db.Close()
	if err := db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}

// GetByUserID get posts by user_id
func GetByUserID(req *postpb.GetAllPostsByUserIDRequest) []*postpb.Post {
	var pbs []*postpb.Post

	id := req.GetUserId()

	// connect with DB
	db := gormConnect()
	defer db.Close()

	// get posts related to tags
	posts := []entity.Post{}
	err := db.Preload("Tags").Order("ID desc").Where("user_id = ?", id).Find(&posts).Error
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
	}
	// if not found any posts in DB, return sample
	if len(posts) == 0 {
		pb := makeSamplePost()
		pbs = append(pbs, pb)
	} else {
		for _, post := range posts {
			pbs = append(pbs, convertPost(&post))
		}
	}
	return pbs
}

// Search search posts by title or tags
func Search(req *postpb.SearchPostsRequest) []*postpb.Post {
	var pbs []*postpb.Post

	id := req.GetUserId()
	how := req.GetHow()
	words := req.GetKeywords()

	// process keywords for query
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

	// connect with DB
	db := gormConnect()
	defer db.Close()

	// get posts
	posts := []entity.Post{}

	// check how to search
	if how == "title" {
		err := db.Where("user_id = ?", id).Where("title LIKE ?", query).Find(&posts).Error
		// if not found any posts in DB, return sample
		if len(posts) == 0 {
			log.SetFlags(log.Lshortfile)
			log.Println(err)
			pb := makeSamplePost()
			pbs = append(pbs, pb)
		} else {
			for _, post := range posts {
				pbs = append(pbs, convertPost(&post))
			}
		}
		return pbs
	}
	if how == "tag" {
		var tags []entity.Tag

		// search tags by tag_name
		for _, word := range words {
			var tag entity.Tag
			err := db.Where("tag_name = ?", word).Find(&tag).Error
			if err != nil {
				log.SetFlags(log.Lshortfile)
				log.Println(err)
			} else {
				tags = append(tags, tag)
			}
		}

		// search posts by tag_id
		err := db.Model(&posts).Related(&tags, "Tags").Error
		if err != nil {
			log.SetFlags(log.Lshortfile)
			log.Println(err)
		}
		if len(posts) == 0 {
			// if posts from DB are not found, return SAMPLE
			pb := makeSamplePost()
			pbs = append(pbs, pb)
		} else {
			for _, post := range posts {
				pbs = append(pbs, convertPost(&post))
			}
		}
		return pbs
	}
	return pbs
}
