package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"finalProject/controllers"
	"finalProject/models"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("dist", false)))
	router.Use(cors.Default())

	// Connect to database
	models.ConnectDatabase()

	// Routes
	router.GET("/api/biodata/list", controllers.GetBiodatas)
	router.POST("/api/biodata/post", controllers.PostBiodata)

	router.Run(":5000")
}
