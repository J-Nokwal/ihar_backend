package models

import (
	"fmt"
	"time"
)

type User struct {
	ID               string `json:"userId" gorm:"primaryKey;NOT NULL"`
	IsAnoymous       bool   `json:"is_anaoymous" gorm:"type:bool;default:true"`
	Email            string `json:"email" gorm:"type:varchar(100)"`
	ProfileLink      string `jsnon:"profile_link" gorm:"type:varchar(100)"`
	ProfilePhotoLink string `jsnon:"profile_link" gorm:"type:varchar(100)"`
	FirstName        string `jsnon:"first_name" gorm:"type:varchar(100)"`
	LastName         string `jsnon:"last_name" gorm:"type:varchar(100)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
}

func init() {
	// usr := User{Id: "#4"}
	// usr := CreditCard{Id: 3344}
	// var cred [2]CreditCard
	// cred[0] = CreditCard{Id: 3322, Number: "dfd"}
	// usr := User{ID: "dasd", FirstName: "asdadd"}

	// usr := Posts{Message: "sfdf", UserID: "dasdx"}

	// result := db.Create(&usr)
	// // fmt.Println("-----------", result, result.Error,), "\n")
	// if result.Error != nil {
	// fmt.Println("+++++++++++", result.Error)
	// }
}

func (user *User) CreateUser() (*User, error) {
	if err := db.NewRecord(user); err {
		return nil, fmt.Errorf("invalid primarykey source, userId  is null")
	}
	if err := db.Create(&user).Error; err != nil {

		return nil, fmt.Errorf("error while insertion")

	}
	return user, nil
}

func GetAllUser() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id string) (*User, error) {
	var GetUser User
	err := db.Where("id=?", Id).Find(&GetUser).Error
	if err != nil {
		return nil, fmt.Errorf("id not found")
	}
	return &GetUser, nil
}

func GetUsersForSearchQuery(searchQuery string) (*[]User, error) {
	searchQuery = "%" + searchQuery + "%"
	var GetUsers []User
	err := db.Where("first_name LIKE ? OR last_name LIKE ?  OR id LIKE ?", searchQuery, searchQuery, searchQuery).Find(&GetUsers).GetErrors()
	if len(err) != 0 {
		fmt.Println(err)
		return nil, fmt.Errorf("error while searching")
	}
	return &GetUsers, nil
}

func UpdateUser(user User) (*User, error) {
	// var user User
	d := db.Model(&user).Where("id=?", user.ID).Update(user)
	err := d.Error
	if err != nil {
		return nil, fmt.Errorf("error while patch query")
	}
	d.Find(&user)
	return &user, nil

}
