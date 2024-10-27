package handlers

import (
	"yuki-image/conf"
	"yuki-image/middlewares"

	"github.com/gin-gonic/gin"
)

var api *gin.RouterGroup

func InitRoute() {
	server.Static("/i", conf.Conf.Server.Path)
	api = server.Group("/api/v1")
	api.Use(middlewares.TokenAuthMiddleware())
	api.Use(middlewares.PingDB())
	InitAlbum()
	InitFormat()
	InitImage()
	InitSystem()
}

func InitAlbum() {
	albumRoute := api.Group("/album")
	{
		albumRoute.GET("/", SelectAllAlbum)
		albumRoute.GET("/:id", SelectAlbum)
		albumRoute.POST("/", InsertAlbum)
		albumRoute.POST("/format", InsertFormatSupport)
		albumRoute.GET("/format/:id", SelectFormatSupport)
		albumRoute.DELETE("/format", DeleteFormatSupport)
		albumRoute.POST("/:id", UpdateAlbum)
		albumRoute.GET("/image/:id", SelectImageFromAlbum)
	}
}

func InitFormat() {
	formatRoute := api.Group("/format")
	{
		formatRoute.GET("/", SelectAllFormat)
		formatRoute.GET("/:id", SelectFormat)
	}
}

func InitImage() {
	imageRoute := api.Group("/image")
	{
		imageRoute.POST("/", UploadImage)
		imageRoute.GET("/:id", SelectImage)
		imageRoute.DELETE("/:id", DeleteImage)
	}
}

func InitSystem() {
	systemRoute := api.Group("/system")
	{
		systemRoute.GET("/tmp", GetTmpInfo)
		systemRoute.DELETE("/tmp", ClearTmp)
	}
}
