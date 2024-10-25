package handlers

import (
	"yuki-image/conf"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func InitHandlers() {
	server = gin.Default()
	InitRoute()
	server.Run(":" + conf.Conf.Server.Port)
}
