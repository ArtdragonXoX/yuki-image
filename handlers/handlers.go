package handlers

import (
	"yuki-image/conf"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func InitHandlers() {
	server = gin.Default()
	server.Static("/i", conf.Conf.Server.Path)
	server.Run(":" + conf.Conf.Server.Port)
}
