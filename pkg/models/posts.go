package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Message         string `json:"message" gorm:"type:MEDIUMTEXT"`
	Post_photo_link string `json:"post_photo_link" gorm:"type:MEDIUMTEXT"`
	Likes           int    `json:"likes" gorm:"type:INT"`
	// Posting_User    User   `json:"posting_user" `
	// Posting_User User `json:"posting_user" binding:"required"`
	User   User   `json:"postFrom"`
	UserID string `json:"userId"`
	Liked  bool   `json:"liked" gorm:"-"`
	// Posting_User    User   `json:"posting_user" gorm:"references:ID"`
}

func (post *Post) CreatePost() (*Post, error) {
	if errList := db.Create(&post).GetErrors(); len(errList) != 0 {
		fmt.Println(errList)
		return nil, fmt.Errorf("error while insertion")

	}
	user := User{}
	fmt.Println("11111111", db.Model(&post).Related(&user, "user_id").GetErrors())

	// post4.Posting_User = user
	return post, nil
}

func GetAllPost() ([]Post, error) {
	var posts []Post
	if errList := db.Preload("User").Order("created_at DESC").Find(&posts).GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("exctraction error")
	}
	return posts, nil
}
func GetPostByPageId(offset int, pageSize int, beforeDateTime time.Time) ([]Post, error) {
	var posts []Post
	if errList := db.Preload("User").Limit(pageSize).Offset(offset).Order("created_at DESC").Where("created_at < ?", beforeDateTime).Find(&posts).GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("exctraction error")
	}
	return posts, nil
}
func GetPostForSearchQuery(searchQuery string) ([]Post, error) {
	var posts []Post
	if errList := db.Preload("User").Order("created_at DESC").Where("message  LIKE ?", "%"+searchQuery+"%").Find(&posts).GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("exctraction error")
	}
	return posts, nil
}
func GetAllPostFromUser(id string) ([]Post, error) {
	var posts []Post
	if errList := db.Where("user_id=?", id).Find(&posts).GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("exctraction error")
	}
	return posts, nil
}

func GetPostById(Id int) (*Post, error) {
	var GetPost Post
	if errList := db.Where("id=?", Id).Find(&GetPost).GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("exctraction error")
	}
	return &GetPost, nil
}

func CountPosts() (*int64, error) {
	var count int64
	if errList := db.Model(&Post{}).Count(&count).GetErrors(); len(errList) != 0 {
		fmt.Println("error while counting", errList)
		return nil, fmt.Errorf("exctraction error")
	}
	return &count, nil
}

func UpdatePost(post Post) (*Post, error) {
	d := db.Model(&post).Where("id=?", post.ID).Update(post)
	if errList := d.GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("error while patch query")
	}
	d.Find(&post)
	return &post, nil
}

func DeletePost(Id int) error {
	post := Post{}
	errList := db.Where("id=?", Id).Delete(&post).GetErrors()
	if len(errList) != 0 {
		fmt.Println("111111", errList)
		return fmt.Errorf("error while deleting")

	}
	return nil
}
