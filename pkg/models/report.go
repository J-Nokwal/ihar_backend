package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	Message    string `json:"message" gorm:"type:MEDIUMTEXT"`
	Post       Post   `json:"-" `
	PostID     *uint  `json:"postId" `
	User       User   `json:"-"`
	UserID     string `json:"userId" `
	FromUser   User   `json:"-"`
	FromUserID string `json:"fromUser" `
}

func (reportModal Report) ReportUserOrPost() error {
	if errList := db.Create(&reportModal).GetErrors(); len(errList) != 0 {
		fmt.Println(errList)
		return fmt.Errorf("error while insertion")
	}
	// user := User{}
	// reportModal2 := Report{}
	// fmt.Println("11111111", db.Preload("User").Where("id = ?", reportModal.ID).Find(&reportModal2).GetErrors())
	// fmt.Println("11111111", db.Model(&reportModal).Related(&user, "from_user_id").GetErrors())
	// k, _ := json.Marshal(&reportModal2)
	// fmt.Println(string(k))
	return nil
}
