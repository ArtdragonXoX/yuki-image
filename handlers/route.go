package handlers

import (
	"yuki-image/conf"
)

func InitRoute() {
	server.Static("/i", conf.Conf.Server.Path)
	InitAlbum()
}

func InitAlbum() {
	albumRoute := server.Group("/album")
	{
		albumRoute.POST("/", InsertAlbum)
	}
}
