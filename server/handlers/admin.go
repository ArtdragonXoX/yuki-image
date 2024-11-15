package handlers

import (
	"net/http"
	"yuki-image/internal/admin"
	"yuki-image/internal/model"

	"github.com/gin-gonic/gin"
)

func AdminRegister(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	err = admin.UserRegister(u)
	if err != nil {
		c.JSON(http.StatusOK, model.RespError("注册失败", err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.RespOk("注册成功", nil))
}

func AdminLogin(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	token, err := admin.UserLogin(u)
	if err != nil {
		c.JSON(http.StatusOK, model.RespError("登录失败", err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.RespOk("登录成功", token))
}
