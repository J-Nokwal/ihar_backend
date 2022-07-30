package controllers

import (
	"net/http"
	"strconv"

	"github.com/J-Nokwal/ihar_backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "code": "1000"})
		return
	}
	if post.Likes != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while insertion, likes not 0 while posting", "code": "1001"})
		return
	}
	if _, err := post.CreatePost(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1001"}) // error: "error while insertion"
		return
	}
	c.JSON(http.StatusOK, post)
}
func GetPostForUser(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	byUser := c.Param("byUser")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter is not int", "code": "1004"}) // error: "params are not int"

	}
	post, err := models.GetPostById(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		return
	}
	a := uint(postId)
	liked, err := models.CheckIfLiked(&a, byUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"

	}
	post.Liked = *liked
	c.JSON(http.StatusOK, post)
}
func GetAllPostOfUserByUserId(c *gin.Context) {
	byUser := c.Param("byUser")
	ofUser := c.Param("ofUser")

	posts, err := models.GetAllPostFromUser(ofUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		return
	}

	for _, j := range posts {
		liked, err := models.CheckIfLiked(&j.ID, byUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		}
		j.Liked = *liked
	}
	c.JSON(http.StatusOK, posts)

}
func GetAllPostByUserId(c *gin.Context) {
	byUser := c.Param("byUser")

	posts, err := models.GetAllPost()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		return
	}

	for _, j := range posts {
		liked, err := models.CheckIfLiked(&j.ID, byUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		}
		j.Liked = *liked
	}
	c.JSON(http.StatusOK, posts)

}
func PatchPost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "code": "1000"})
		return
	}
	u, err := models.UpdatePost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1003"}) // error: "error while patching"
		return
	}
	c.JSON(http.StatusOK, u)
}
func DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter is not int", "code": "1004"}) // error: "params are not int"
		return
	}
	err = models.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "code": "1006"}) // error: "params are not int"
		return
	}

	c.Status(http.StatusOK)
}
