package middlewares

import (
	"net/http"
	"yuki-image/conf"
	"yuki-image/model"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// 通常Bearer token前面会有"Bearer "这个字符串，需要去掉它来比较
		if token != "Bearer "+conf.Conf.Server.Token {
			c.JSON(http.StatusUnauthorized, model.Response{Code: 0, Msg: "Unauthorized", Data: nil})
			c.Abort() // 终止请求
			return
		}
		c.Next() // 继续处理请求
	}
}
