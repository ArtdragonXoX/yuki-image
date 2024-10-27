package server

import (
	"yuki-image/internal/conf"
	"yuki-image/server/handlers"
	"yuki-image/server/middlewares"

	"github.com/gin-gonic/gin"
)

func NewAndInit() error {
	e := New()
	Init(e)
	return e.Run(":" + conf.Conf.Server.Port)
}

func Init(e *gin.Engine) {
	InitRoute(e)
}

func New() *gin.Engine {
	r := gin.Default()
	return r
}

func InitRoute(server *gin.Engine) {
	server.Static("/i", conf.Conf.Server.Path)
	api := server.Group("/api/v1")
	api.Use(middlewares.TokenAuthMiddleware())
	api.Use(middlewares.PingDB())
	InitAlbum(api)
	InitFormat(api)
	InitImage(api)
	InitSystem(api)
}

func InitAlbum(api *gin.RouterGroup) {
	albumRoute := api.Group("/album")
	{
		albumRoute.GET("/", handlers.SelectAllAlbum)
		albumRoute.GET("/:id", handlers.SelectAlbum)
		albumRoute.POST("/", handlers.InsertAlbum)
		albumRoute.POST("/format", handlers.InsertFormatSupport)
		albumRoute.GET("/format/:id", handlers.SelectFormatSupport)
		albumRoute.DELETE("/format", handlers.DeleteFormatSupport)
		albumRoute.POST("/:id", handlers.UpdateAlbum)
		albumRoute.GET("/image/:id", handlers.SelectImageFromAlbum)
	}
}

func InitFormat(api *gin.RouterGroup) {
	formatRoute := api.Group("/format")
	{
		formatRoute.GET("/", handlers.SelectAllFormat)
		formatRoute.GET("/:id", handlers.SelectFormat)
	}
}

func InitImage(api *gin.RouterGroup) {
	imageRoute := api.Group("/image")
	{
		imageRoute.POST("/", handlers.UploadImage)
		imageRoute.GET("/:id", handlers.SelectImage)
		imageRoute.DELETE("/:id", handlers.DeleteImage)
	}
}

func InitSystem(api *gin.RouterGroup) {
	systemRoute := api.Group("/system")
	{
		systemRoute.GET("/tmp", handlers.GetTmpInfo)
		systemRoute.DELETE("/tmp", handlers.ClearTmp)
	}
}