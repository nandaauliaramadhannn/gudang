package controllers

import (
	"gudang/config"
	"gudang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"role":     user.Role,
	})
}
