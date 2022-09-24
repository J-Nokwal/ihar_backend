package controllers

import (
	"net/http"

	"github.com/J-Nokwal/ihar_backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetSearchQueryResults(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	searchQuery := c.Param("searchQuery")
	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "searchQuery cant be empty or null", "code": "1004"}) // error:  "searchQuery cant be empty or null"
	}
	byUser := c.Param("byUser")
	posts, err := models.GetPostForSearchQuery(searchQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		return
	}

	for i, j := range posts {
		liked, err := models.CheckIfLiked(&j.ID, byUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
			return
		}
		posts[i].Liked = *liked
	}

	users, err := models.GetUsersForSearchQuery(searchQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"users": users,
	})

}
