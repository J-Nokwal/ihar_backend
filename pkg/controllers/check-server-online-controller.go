package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsServerOnline(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	// fmt.Println()
	// c.Status(http.)
	c.JSON(http.StatusOK, gin.H{"status": "okkk"})
	// http.StatusNoContent

}
