package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/dashboard.html", nil)
}
