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
	server.Static("/i", conf.Conf.Image.Path)
	server.MaxMultipartMemory = int64(conf.Conf.Image.MaxSize) << 20
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
		albumRoute.GET("/all", handlers.SelectAllAlbum)
		albumRoute.GET("", handlers.SelectAlbum)
		albumRoute.POST("", handlers.InsertAlbum)
		albumRoute.POST("/format", handlers.InsertFormatSupport)
		albumRoute.GET("/format", handlers.SelectFormatSupport)
		albumRoute.DELETE("/format", handlers.DeleteFormatSupport)
		albumRoute.PUT("", handlers.UpdateAlbum)
		albumRoute.GET("/image", handlers.SelectImageFromAlbum)
	}
}

func InitFormat(api *gin.RouterGroup) {
	formatRoute := api.Group("/format")
	{
		formatRoute.GET("", handlers.SelectAllFormat)
		formatRoute.GET("/:id", handlers.SelectFormat)
	}
}

func InitImage(api *gin.RouterGroup) {
	imageRoute := api.Group("/image")
	{
		imageRoute.POST("", handlers.UploadImage)
		imageRoute.GET("/:id", handlers.SelectImage)
		imageRoute.DELETE("/:id", handlers.DeleteImage)
		imageRoute.GET("", handlers.SelectImageFromUrl)
	}
}

func InitSystem(api *gin.RouterGroup) {
	systemRoute := api.Group("/system")
	{
		systemRoute.GET("/tmp", handlers.GetTmpInfo)
		systemRoute.DELETE("/tmp", handlers.ClearTmp)
	}
}
