package middlewares

import (
	"net/http"
	"time"
	"yuki-image/internal/model"

	"yuki-image/utils"

	"github.com/gin-gonic/gin"
)

func tokenAutoRefresh(c *gin.Context) error {
	exp, err := utils.GetTokenExpire(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token无效，获取用户信息失败", nil))
		c.Abort()
		return err
	}

	// 计算token剩余时间
	timeLeft := exp - uint64(time.Now().Unix())
	//log.Println(timeLeft, config.Refresh)
	if timeLeft > utils.Refresh {
		return nil
	}

	name, err := utils.GetTokenName(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token无效，获取用户信息失败", nil))
		c.Abort()
		return err
	}

	// 生成新token
	token, err := utils.GenerateToken(name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token刷新失败", nil))
		c.Abort()
		return err
	}

	c.JSON(http.StatusUnauthorized, model.RespRetry("token已刷新，请重新发送请求", token))
	c.Abort()
	return nil
}

func tokenVerify(c *gin.Context) error {
	err := utils.VerifyToken(c)
	if err != nil {
		return err
	}
	return nil
}
