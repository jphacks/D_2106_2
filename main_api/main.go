package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jphacks/D_2106_2/api"
	"github.com/jphacks/D_2106_2/config"
	"github.com/jphacks/D_2106_2/database"
	"github.com/jphacks/D_2106_2/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// database setup
	config, err := config.GetConfig()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	sqlHandler, err := database.NewSqlClient(config)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	// s3 client
	s3client, err := api.NewS3client()
	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()

	// handlers
	userRepo := database.NewUserRepository(*sqlHandler)
	albumRepo := database.NewAlbumRepository(*sqlHandler)
	coordinateRepo := database.NewCoordinateRepository(*sqlHandler)
	imageRepo := database.NewImageRepository(*sqlHandler)
	s3service := api.NewS3service(*s3client)

	userHandler := handler.NewUserHandler(userRepo)
	albumHandler := handler.NewAlbumHandler(albumRepo, coordinateRepo, imageRepo)
	imageHandler := handler.NewImageHandler(imageRepo, s3service, coordinateRepo)

	// routing
	r.Use(cors.Default()) // cors

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	checkFlaskApi := r.Group("/check_flask_api")
	checkFlaskApi.GET("/get_sample", func(c *gin.Context) {
		data := api.GetSampleApi()
		c.JSON(http.StatusOK, data)
	})
	checkFlaskApi.GET("/clustering", func(c *gin.Context) {
		data, err := api.GetCheckClusteringApi()
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/albums", func(c *gin.Context) { albumHandler.GetAllAlbums(c) })
	r.GET("/albums/user", func(c *gin.Context) { albumHandler.GetUserAlbums(c) })
	r.GET("/album/detail", func(c *gin.Context) { albumHandler.GetAlbumDetail(c) })
	r.POST("/album", func(c *gin.Context) { albumHandler.PostAlbum(c) })

	r.POST("/upload/image", func(c *gin.Context) { imageHandler.UploadImages(c) })

	r.POST("/user", func(c *gin.Context) { userHandler.RegisterUser(c) })
	r.GET("/user", func(c *gin.Context) { userHandler.GetUser(c) })

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
