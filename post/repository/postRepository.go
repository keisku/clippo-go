package repository

import (
	"strconv"
	"time"

	"github.com/kskumgk63/clippo-go/post/entity"
	"github.com/kskumgk63/clippo-go/post/postpb"
)

// Create 投稿を作成
func Create(req *postpb.CreatePostRequest) error {
	// 文字列で受け取るのでuintへ変換
	resID := req.GetPost().GetUserId()
	id64, _ := strconv.ParseUint(resID, 10, 64)
	id := uint(id64)

	post := entity.Post{
		URL:         req.GetPost().GetUrl(),
		Title:       req.GetPost().GetTitle(),
		Description: req.GetPost().GetDescription(),
		Image:       req.GetPost().GetImage(),
		Usecase:     req.GetPost().GetUsecase(),
		Genre:       req.GetPost().GetGenre(),
		UserID:      id,
	}

	db := GormConnect()
	err := db.Create(&post).Error
	if err != nil {
		return err
	}
	db.Model(&post).Update("CreatedAt", time.Now().Add(9*time.Hour))
	db.Model(&post).Update("UpdatedAt", time.Now().Add(9*time.Hour))

	return nil
}

// Delete 投稿を削除
func Delete(req *postpb.DeletePostRequest) error {
	var post entity.Post
	id64, _ := strconv.ParseUint(req.GetId(), 10, 64)
	post.ID = uint(id64)
	// データベースと接続
	db := GormConnect()
	defer db.Close()
	err := db.Delete(&post).Error
	if err != nil {
		return err
	}
	return nil
}

// GetByUserID ユーザーIDに紐づく投稿を全取得
func GetByUserID(req *postpb.GetAllPostsByUserIDRequest) []entity.Post {
	id := req.GetUserId()

	// データベースと接続
	db := GormConnect()
	defer db.Close()

	// 投稿一覧取得
	Posts := []entity.Post{}
	db.Find(&Posts)
	db.Where("user_id = ?", id).Find(&Posts)
	return Posts
}

// SearchByTitle 投稿のタイトル検索
func SearchByTitle(req *postpb.SearchPostsByTitleRequest) []entity.Post {
	// このユーザーIDを基にDB検索
	id := req.GetUserId()
	title := req.GetTitle()

	// データベースと接続
	db := GormConnect()
	defer db.Close()
	// 投稿一覧取得
	Posts := []entity.Post{}
	db.Where("user_id = ? AND title LIKE ?", id, "%"+title+"%").Find(&Posts)
	return Posts
}

// SearchByUsecase 投稿のタイトル検索
func SearchByUsecase(req *postpb.SearchPostsByUsecaseRequest) []entity.Post {
	// このユーザーIDを基にDB検索
	id := req.GetUserId()
	usecase := req.GetUsecase()

	// データベースと接続
	db := GormConnect()
	defer db.Close()
	// 投稿一覧取得
	Posts := []entity.Post{}
	db.Where("user_id = ? AND usecase LIKE ?", id, "%"+usecase+"%").Find(&Posts)
	return Posts
}

// SearchByGenre 投稿のタイトル検索
func SearchByGenre(req *postpb.SearchPostsByGenreRequest) []entity.Post {
	// このユーザーIDを基にDB検索
	id := req.GetUserId()
	genre := req.GetGenre()

	// データベースと接続
	db := GormConnect()
	defer db.Close()
	// 投稿一覧取得
	Posts := []entity.Post{}
	db.Where("user_id = ? AND genre LIKE ?", id, "%"+genre+"%").Find(&Posts)
	return Posts
}
