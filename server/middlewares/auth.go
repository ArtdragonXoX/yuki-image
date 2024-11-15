package middlewares

import (
	"net/http"
	"yuki-image/internal/conf"
	"yuki-image/internal/model"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		v := false
		// 通常Bearer token前面会有"Bearer "这个字符串，需要去掉它来比较
		if token == "Bearer "+conf.Conf.Server.Token {
			v = true
		}
		if !v {
			if err := tokenVerify(c); err == nil {
				v = true
			}
		}
		if !v {
			c.JSON(http.StatusUnauthorized, model.RespError("token验证失败", nil))
			return
		}
		c.Next() // 继续处理请求
	}
}
