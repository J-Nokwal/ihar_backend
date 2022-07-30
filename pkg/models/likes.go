package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Likes struct {
	gorm.Model
	Post   Post   `json:"-" `
	PostID *uint  `json:"postId" `
	User   User   `json:"liked_from_user"`
	UserID string `json:"likedBy" `
}

func CheckIfLiked(postId *uint, userId string) (*bool, error) {
	like := Likes{}
	var liked bool
	d := db.Where("post_id = ? AND  user_id = ?", &postId, userId).Find(&like)
	if d.Error != nil {
		if d.Error == gorm.ErrRecordNotFound {
			liked = false
		} else {
			return nil, fmt.Errorf("exctraction error")
		}
	} else {
		liked = true
	}

	return &liked, nil
}

func (like Likes) Like() error {
	if errList := db.Create(&like).GetErrors(); len(errList) != 0 {
		fmt.Println(errList)
		return fmt.Errorf("error while insertion")
	}
	post := Post{}
	db.Model(&like).Related(&post)
	db.Model(&post).Update("likes", post.Likes+1)
	return nil

}
func (like Likes) UnLike() error {
	post := Post{}
	errList := db.Model(&like).Related(&post).GetErrors()
	if len(errList) != 0 {
		fmt.Println("111111", errList)
		return fmt.Errorf("internal error while like")
	}
	errList = db.Unscoped().Where("post_id=? AND user_id=?", like.PostID, like.UserID).Delete(&like).GetErrors()
	if len(errList) != 0 {
		fmt.Println("111111", errList)
		return fmt.Errorf("internal error while like")
	}
	db.Model(&post).Update("likes", post.Likes-1)
	return nil
}

func GetUsersLikesByPost(postId int) (*[]Likes, error) {
	likes := []Likes{}
	if errList := db.Preload("User").Where("post_id=?", postId).Find(&likes).GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("exctraction error")
	}
	return &likes, nil

}
