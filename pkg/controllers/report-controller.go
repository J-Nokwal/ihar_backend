package controllers

import (
	"net/http"

	"github.com/J-Nokwal/ihar_backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func ReportUserOrPost(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var report models.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "code": "1000"})
		return
	}

	if report.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID not provided in report", "code": "1001"})
		return
	}
	if report.FromUserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": " fromUser not provided in report", "code": "1001"})
		return
	}
	if err := report.ReportUserOrPost(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "1001"})
	}
	// c.JSON(http.StatusOK, nil)
	c.Status(http.StatusAccepted)
}
