package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kskumgk63/clippo-go/post/entity"
	"github.com/kskumgk63/clippo-go/post/postpb"

	"github.com/speps/go-hashids"
)

// Create 投稿を作成
func Create(req *postpb.CreatePostRequest) error {
	// connect with db
	db := gormConnect()
	defer db.Close()

	// 文字列で受け取るのでuintへ変換
	resID := req.GetPost().GetUserId()
	id64, _ := strconv.ParseUint(resID, 10, 64)
	id := uint(id64)

	tags := entity.Tag{}
	tagNames := req.GetPost().GetTagId()
	var tagID string
	for i, tagName := range tagNames {
		// check if the tag_name is existed, if not create new
		if err := db.Where("tag_name = ?", tagName).Find(&tags).Error; err != nil {
			// generate short uuid
			hd := hashids.NewData()
			hd.MinLength = 8
			h, _ := hashids.NewWithData(hd)
			e, _ := h.Encode([]int{45, 434, 1313, 99})
			fmt.Println(e)

			// create new tag
			db.Create(&tags{
				TagName: tagName,
				TagID:   e,
			})
			log.Println(err)
		}
		if i == 0 {
			tagID += tags.TagID
		} else {
			tagID += "/" + tags.TagID
		}
	}

	post := entity.Post{
		URL:         req.GetPost().GetUrl(),
		Title:       req.GetPost().GetTitle(),
		Description: req.GetPost().GetDescription(),
		Image:       req.GetPost().GetImage(),
		TagID:       tagID,
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
func GetByUserID(req *postpb.GetAllPostsByUserIDRequest) []entity.Post {
	id := req.GetUserId()

	// データベースと接続
	db := gormConnect()
	defer db.Close()

	// 投稿一覧取得
	posts := []entity.Post{}
	if err := db.Order("ID desc").Where("user_id = ?", id).Find(&posts).Error; err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		return nil
	}
	return posts
}

// SearchByTitle 投稿のタイトル検索
func SearchByTitle(req *postpb.SearchPostsByTitleRequest) []entity.Post {
	// このユーザーIDを基にDB検索
	id := req.GetUserId()
	words := req.GetTitles()

	// 複数の検索ワードで検索できるように配列の文字列を加工
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

	// データベースと接続
	db := gormConnect()
	defer db.Close()

	// 投稿一覧取得
	posts := []entity.Post{}
	if err := db.Where("user_id = ?", id).Where("title LIKE ?", query).Find(&posts).Error; err != nil {
		log.SetFlags(log.Lshortfile)
		log.Println(err)
		return nil
	}
	return posts
}
