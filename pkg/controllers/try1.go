package controllers

import (
	"fmt"
	"math/rand"

	"github.com/J-Nokwal/ihar_backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func Try_this1(c *gin.Context) {
	user := models.User{ID: RandStringBytes(5)}
	// user := models.User{UserID: 4}
	// user := models.User{ID: "4444"}

	user.CreateUser()
	c.JSON(200, &user)
	fmt.Println(user)

}

func Try_this2(c *gin.Context) {
	allUsers := models.GetAllUser()
	c.JSON(200, &allUsers)
	fmt.Println(allUsers)

	// user := User{"aaa", 3}
	// c.JSON(200, &us)
	// c.ShouldBind(&user)
}

func Try_this3(c *gin.Context) {
	// user := models.UpdateUser()
	// c.JSON(200, &user)
	// fmt.Println(user)

	// user := User{"aaa", 3}
	// c.JSON(200, &us)
	// c.ShouldBind(&user)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
