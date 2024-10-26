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
	InitAlbum()
	InitFormat()
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

	}
}

func InitFormat() {
	formatRoute := api.Group("/format")
	{
		formatRoute.GET("/", SelectAllFormat)
		formatRoute.GET("/:id", SelectFormat)
	}
}
