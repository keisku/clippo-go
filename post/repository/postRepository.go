package repository

import (
	"log"
	"strconv"
	"strings"

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

func convertPost(post *entity.Post) (pb *postpb.Post) {
	// convert uint to string
	post64 := uint64(post.ID)
	postID := strconv.FormatUint(post64, 10)

	// convert uint to string
	user64 := uint64(post.UserID)
	userID := strconv.FormatUint(user64, 10)

	// convert string to string array
	tagArray := strings.Split(post.Tag, "/")

	pb = &postpb.Post{
		Id:          postID,
		Url:         post.URL,
		Title:       post.Title,
		Description: post.Description,
		Image:       post.Image,
		Tag:         tagArray,
		UserId:      userID,
	}
	return pb
}

// Create 投稿を作成
func Create(req *postpb.CreatePostRequest) error {
	// connect with db
	db := gormConnect()
	defer db.Close()

	// convert string to uint
	resID := req.GetPost().GetUserId()
	id64, _ := strconv.ParseUint(resID, 10, 64)
	id := uint(id64)

	tag := entity.Tag{}
	tagNames := req.GetPost().GetTag()
	var postTag string

	// array to string
	for i, tagName := range tagNames {
		// TODO: handle SQL problem
		// the last word of array never be found in DB, it necessarily save

		// check if the tag_name is existed, if not create new
		if err := db.Where("tag_name = ?", tagName).Find(&tag).Error; err != nil {
			// save tag in DB
			db.Create(&entity.Tag{
				TagName: tagName,
			})
		}
		if i == 0 {
			postTag += tagName
		} else {
			postTag += "/" + tagName
		}
	}

	post := entity.Post{
		URL:         req.GetPost().GetUrl(),
		Title:       req.GetPost().GetTitle(),
		Description: req.GetPost().GetDescription(),
		Image:       req.GetPost().GetImage(),
		Tag:         postTag,
		UserID:      id,
	}

	if err := db.Create(&post).Error; err != nil {
		return err
	}

	return nil
}

// Delete 投稿を削除
func Delete(req *postpb.DeletePostRequest) error {
	var post entity.Post
	id64, _ := strconv.ParseUint(req.GetId(), 10, 64)
	post.ID = uint(id64)
	// データベースと接続
	db := gormConnect()
	defer db.Close()
	if err := db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}

// GetByUserID ユーザーIDに紐づく投稿を全取得
func GetByUserID(req *postpb.GetAllPostsByUserIDRequest) []*postpb.Post {
	var pbs []*postpb.Post

	id := req.GetUserId()

	// connect with DB
	db := gormConnect()
	defer db.Close()

	// get posts
	posts := []entity.Post{}
	if err := db.Order("ID desc").Where("user_id = ?", id).Find(&posts).Error; err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		// if posts from DB are not found, return SAMPLE
		pb := makeSamplePost()
		pbs = append(pbs, pb)
		return pbs
	}
	for _, post := range posts {
		pbs = append(pbs, convertPost(&post))
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
		if err := db.Where("user_id = ?", id).Where("title LIKE ?", query).Find(&posts).Error; err != nil {
			log.SetFlags(log.Lshortfile)
			log.Println(err)
			// if posts from DB are not found, return SAMPLE
			pb := makeSamplePost()
			pbs = append(pbs, pb)
			return pbs
		}
		for _, post := range posts {
			pbs = append(pbs, convertPost(&post))
		}
	}
	if how == "tag" {
		if err := db.Where("user_id = ?", id).Where("tag LIKE ?", query).Find(&posts).Error; err != nil {
			log.SetFlags(log.Lshortfile)
			log.Println(err)
			// if posts from DB are not found, return SAMPLE
			pb := makeSamplePost()
			pbs = append(pbs, pb)
			return pbs
		}
		for _, post := range posts {
			pbs = append(pbs, convertPost(&post))
		}
	}
	return pbs
}
