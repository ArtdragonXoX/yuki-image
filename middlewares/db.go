package middlewares

import (
	"net/http"
	"yuki-image/db"
	"yuki-image/model"

	"github.com/gin-gonic/gin"
)

func PingDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			if err := db.InitDataBase(); err != nil {
				c.JSON(http.StatusInternalServerError, model.Response{Code: 0, Msg: "数据库连接失败", Data: nil})
				return
			}
		}
		c.Next()
	}
}
