package controllers

import (
	"net/http"
	"strconv"

	"github.com/J-Nokwal/ihar_backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "code": "1000"})
		return
	}
	if comment.UserID == "" || comment.PostID == nil || comment.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while insertion, userId or postId or message is null", "code": "1001"})
		return
	}
	if _, err := comment.CreateComment(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1001"}) // error: "error while insertion"
		return
	}
	c.JSON(http.StatusOK, comment)

}
func GetAllCommentFromPost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter is not int", "code": "1004"}) // error: "params are not int"
		return
	}
	comments, err := models.GetAllCommentFromPost(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1002"}) // error: "exctraction error"
		return
	}
	c.JSON(http.StatusOK, comments)
}
func DeleteComment(c *gin.Context) {

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter is not int", "code": "1004"}) // error: "params are not int"
		return
	}
	err = models.DeletePost(commentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "code": "1006"}) // error: "params are not int"
		return
	}
	c.Status(http.StatusOK)
}
func PatchComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "code": "1000"})
		return
	}
	u, err := comment.UpdateComment()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1003"}) // error: "error while patching"
		return
	}
	c.JSON(http.StatusOK, u)
}
