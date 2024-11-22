package handlers

import (
	"net/http"
	"yuki-image/internal/conf"
	"yuki-image/internal/db"
	"yuki-image/internal/model"
	"yuki-image/internal/tmp"

	"github.com/gin-gonic/gin"
)

func GetTmpInfo(ctx *gin.Context) {
	tmpInof, err := tmp.GetInfo()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Failed to obtain temporary file information.", Data: err})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "Obtain Success", Data: tmpInof})
}

func ClearTmp(ctx *gin.Context) {
	err := tmp.Clear()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Cleanup Failure", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "Cleanup Successfully"})
}

func GetConf(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.RespOk("Obtain Success", conf.Conf))
}

func UpdateConf(ctx *gin.Context) {
	var confv conf.Config
	if err := ctx.ShouldBindJSON(&confv); err != nil {
		ctx.JSON(http.StatusBadRequest, model.RespError("Failed to obtain configuration", err.Error()))
		return
	}
	if err := conf.Conf.Update(confv); err != nil {
		ctx.JSON(http.StatusBadRequest, model.RespError("Failed to update configuration", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, model.RespOk("Update Successfully", conf.Conf))
}

func GetToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.RespOk("Obtain Success", conf.Conf.Server.Token))
}

func UpdateToken(ctx *gin.Context) {
	token, err := conf.GenerateToken()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.RespError("Failed to generate token", err))
		return
	}
	conf.UpdateToken(token)
	ctx.JSON(http.StatusOK, model.RespOk("Update Successfully", token))
}

func ResetDB(ctx *gin.Context) {
	err := db.ResetDB()
	db.InitDataBase()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.RespError("Failed to reset database", err))
		return
	}
	ctx.JSON(http.StatusOK, model.RespOk("Reset Successfully", nil))
}
