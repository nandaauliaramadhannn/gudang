package controllers

import (
	"gudang/config"
	"gudang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageProduct(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Preload("Kategori").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "product/product_list.html", gin.H{
		"Title":    "Product List",
		"Products": products,
	})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var kategori models.Kategori
	if err := config.DB.First(&kategori, product.KategoriID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/products")
}
