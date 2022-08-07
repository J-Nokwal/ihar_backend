package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

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
			return
		}
		j.Liked = *liked
	}
	c.JSON(http.StatusOK, posts)

}

func GetPostByPageIdByUser(c *gin.Context) {
	var pageSize int64 = 10
	pageId, err := strconv.Atoi(c.Param("pageId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter is not int", "code": "1004"}) // error: "params are not int"
	}
	var beforeDateTime time.Time
	var noOfPages int64
	if pageId < 1 {
		count, err := models.CountPosts()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
			return
		}
		// noOfPages = math.Ceil(*count / pageSize)
		noOfPages = int64(math.Ceil(float64(*count) / float64(pageSize)))

		pageId = 0
		fmt.Println(noOfPages)
		beforeDateTime = time.Now()
	} else {
		var e error
		queryTime := c.Query("queryTime")
		if queryTime == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "queryTime must not be empty if pageId > 0", "code": "1004"}) // error: "queryTime not valid"
			return
		}
		// beforeDateTime, e = time.Parse(time.RFC3339, queryTime)
		beforeDateTime, e = time.Parse(time.RFC1123, queryTime)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid time", "code": "1004"})
			return
		}
	}
	byUser := c.Param("byUser")
	posts, err := models.GetPostByPageId(int(pageSize)*pageId, int(pageSize), beforeDateTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		return
	}

	for _, j := range posts {
		liked, err := models.CheckIfLiked(&j.ID, byUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
			return
		}
		j.Liked = *liked
	}
	c.JSON(http.StatusOK, gin.H{
		"posts":     posts,
		"queryTime": beforeDateTime.Format(time.RFC1123),
		"noOfPages": noOfPages,
	})

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
