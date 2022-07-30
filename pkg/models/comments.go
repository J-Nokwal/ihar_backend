package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Message string `json:"message" gorm:"type:MEDIUMTEXT"`
	Post    Post   `json:"-"`
	PostID  *uint  `json:"postId"`
	User    User   `json:"commentFrom"`
	UserID  string `json:"userId"`
}

func (comment *Comment) CreateComment() (*Comment, error) {
	if errList := db.Create(&comment).GetErrors(); len(errList) != 0 {
		fmt.Println(errList)
		return nil, fmt.Errorf("error while insertion")
	}
	user := User{}
	fmt.Println("11111111", db.Model(&comment).Related(&user, "user_id").GetErrors())
	return comment, nil
}

func GetAllCommentFromPost(ID int) (*[]Comment, error) {
	var comments []Comment
	if errList := db.Preload("User").Where("post_id=?", ID).Find(&comments).GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("exctraction error")
	}
	// for i,j := range comments{

	// }
	return &comments, nil
}

func GetCommentById(Id int) (*Comment, error) {
	var GetComment Comment
	if errList := db.Where("id=?", Id).Find(&GetComment).GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("exctraction error")
	}
	return &GetComment, nil
}

func (comment Comment) UpdateComment() (*Comment, error) {
	d := db.Model(&comment).Where("id=?", comment.ID).Update(comment)
	if errList := d.GetErrors(); len(errList) != 0 {
		fmt.Println("11111111", errList)
		return nil, fmt.Errorf("error while patch query")
	}
	d.Find(&comment)
	return &comment, nil

}

func DeleteComment(Id int) error {
	comment := Comment{}
	errList := db.Where("id=?", Id).Delete(&comment).GetErrors()
	if len(errList) != 0 {
		fmt.Println("111111", errList)
		return fmt.Errorf("error while deleting")
	}
	return nil
}
