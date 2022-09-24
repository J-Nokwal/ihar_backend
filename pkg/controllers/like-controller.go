package controllers

import (
	"net/http"
	"strconv"

	"github.com/J-Nokwal/ihar_backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func TriggerLike(c *gin.Context) {

	var likes models.Likes
	if err := c.ShouldBindJSON(&likes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "code": "1000"})
		return
	}
	if likes.UserID == "" || likes.PostID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json, required likedBy and postId", "code": "1000"})
		return
	}
	liked, err := models.CheckIfLiked(likes.PostID, likes.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1001"}) // error: "error while insertion"
		return
	}
	// when post is already liked
	if *liked {
		if err := likes.UnLike(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1001"})
			return
		}
	} else { // when post is not liked
		if err := likes.Like(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1001"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"liked": !*liked,
	})
}

func GetUsersByPostLikes(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter is not int", "code": "1004"}) // error: "params are not int"

	}
	likes, err := models.GetUsersLikesByPost(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		return
	}
	c.JSON(http.StatusOK, likes)

}
