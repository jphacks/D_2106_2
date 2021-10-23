package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jphacks/D_2106_2/api"
	"github.com/jphacks/D_2106_2/config"
	"github.com/jphacks/D_2106_2/database"
	"github.com/jphacks/D_2106_2/handler"
	"github.com/jphacks/D_2106_2/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// application setup
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
	// db, err := sqlHandler.Conn.DB()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()

	r := gin.Default()

	// handlers
	userRepo := database.NewUserRepository(*sqlHandler)
	albumRepo := database.NewAlbumRepository(*sqlHandler)
	coordinateRepo := database.NewCoordinateRepository(*sqlHandler)

	authHandler := handler.NewAuthHandler(userRepo)
	userHandler := handler.NewUserHandler(userRepo)
	albumHandler := handler.NewAlbumHandler(albumRepo, coordinateRepo)
	imageHandler := handler.NewImageHandler()

	// auth middleware
	authMiddleware, err := middleware.GetAuthMiddleware(*authHandler)
	if err != nil {
		fmt.Println(err)
	}

	// routing
	r.Use(cors.Default()) // cors

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/fetch", func(c *gin.Context) {
		data := api.FetchSampleApi()
		c.JSON(http.StatusOK, data)
	})

	r.GET("/albums", func(c *gin.Context) { albumHandler.GetAllAlbums(c) })
	r.GET("/albums/user", func(c *gin.Context) { albumHandler.GetUserAlbums(c) })
	r.GET("/album", func(c *gin.Context) { albumHandler.GetAlbum(c) })
	r.GET("/album/detail", func(c *gin.Context) { albumHandler.GetAlbumDetail(c) })
	r.POST("/album", func(c *gin.Context) { albumHandler.PostAlbum(c) })

	r.POST("/upload/image", func(c *gin.Context) { imageHandler.UploadImages(c) })

	r.POST("/register", func(c *gin.Context) { userHandler.RegisterUser(c, authMiddleware) })
	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/users", func(c *gin.Context) { userHandler.GetAllUsers(c) })
		auth.GET("/user", func(c *gin.Context) { userHandler.GetUser(c) })
	}

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
