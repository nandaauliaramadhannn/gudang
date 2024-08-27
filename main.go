package main

import (
	"fmt"
	"gudang/config"
	"gudang/controllers"
	"gudang/middleware"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Menghubungkan ke database dan melakukan migrasi
	config.ConnectDatabase()
	config.AutoMigrate()

	// Menampilkan direktori kerja saat ini
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	}
	fmt.Println("Current working directory:", cwd)

	// Membuat router Gin
	router := gin.Default()

	// Melayani file statis (CSS, JS, Gambar)
	router.Static("/assets", "./assets")

	// Menggunakan middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // URL frontend
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	router.LoadHTMLGlob("templates/**/*")

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	// Rute yang memerlukan otentikasi
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/profile", controllers.UserProfile)
		authRoutes.GET("/products", controllers.PageProduct)
		authRoutes.GET("/dashboard", controllers.DashboardPage)
	}

	// Menjalankan server di port 8080
	router.Run(":8080")
}
