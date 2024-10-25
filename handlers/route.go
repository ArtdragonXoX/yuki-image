package handlers

import (
	"yuki-image/conf"

	"github.com/gin-gonic/gin"
)

var api *gin.RouterGroup

func InitRoute() {
	server.Static("/i", conf.Conf.Server.Path)
	api = server.Group("/api/v1")
	InitAlbum()
	InitFormat()
}

func InitAlbum() {
	albumRoute := api.Group("/album")
	{
		albumRoute.POST("/", InsertAlbum)
	}
}

func InitFormat() {
	formatRoute := api.Group("/format")
	{
		formatRoute.GET("/", SelectAllFormat)
		formatRoute.GET("/:id", SelectFormat)
	}
}
