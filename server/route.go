package server

import (
	"yuki-image/internal/conf"
	"yuki-image/server/handlers"
	"yuki-image/server/middlewares"
	"yuki-image/static"

	"github.com/gin-gonic/gin"
)

func NewAndInit() error {
	e := New()
	Init(e)
	return e.Run(":" + conf.Conf.Server.Port)
}

func Init(e *gin.Engine) {
	InitStatic(e)
	InitAPIRoutes(e)
	InitAdminRoutes(e)
}

func New() *gin.Engine {
	r := gin.Default()
	return r
}

func InitStatic(server *gin.Engine) {
	server.Static("/i/", conf.Conf.Image.Path)
	server.MaxMultipartMemory = int64(conf.Conf.Image.MaxSize) << 20
	static.InitStatic(server)
}

func InitAdminRoutes(server *gin.Engine) {
	adminRoute := server.Group("/admin")
	{
		adminRoute.POST("/login", handlers.AdminLogin)
		adminRoute.POST("/register", handlers.AdminRegister)
	}
}

func InitAPIRoutes(server *gin.Engine) {
	api := server.Group("/api/v1")
	api.Use(middlewares.TokenAuthMiddleware())
	InitAlbum(api)
	InitFormat(api)
	InitImage(api)
	InitSystem(api)
}

func InitAlbum(api *gin.RouterGroup) {
	albumRoute := api.Group("/album")
	{
		albumRoute.GET("", handlers.SelectAllAlbum)
		albumRoute.GET("/:id", handlers.SelectAlbum)
		albumRoute.POST("", handlers.InsertAlbum)
		albumRoute.PUT("/:id", handlers.UpdateAlbum)
		albumRoute.DELETE("/:id", handlers.DeleteAlbum)
		albumRoute.DELETE("/image/:id", handlers.ClearAlbum)
		albumRoute.POST("/format", handlers.InsertFormatSupport)
		albumRoute.GET("/format/:id", handlers.SelectFormatSupport)
		albumRoute.DELETE("/format", handlers.DeleteFormatSupport)
		albumRoute.GET("/image/:id", handlers.SelectImageFromAlbum)
		albumRoute.GET("/size", handlers.GetAlbumSize)
		albumRoute.GET("/count", handlers.GetAlbumCount)
		albumRoute.GET("/size/:id", handlers.GetAlbumSize)
		albumRoute.GET("/count/:id", handlers.GetAlbumCount)
		albumRoute.GET("/statistics/:id", handlers.SelectStatistics)
		albumRoute.GET("/statistics", handlers.SelectAllStatistics)

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
