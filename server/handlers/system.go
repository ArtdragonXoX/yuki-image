package handlers

import (
	"net/http"
	"yuki-image/internal/model"
	"yuki-image/internal/tmp"

	"github.com/gin-gonic/gin"
)

func GetTmpInfo(ctx *gin.Context) {
	tmpInof, err := tmp.GetInfo()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Failed to obtain temporary file information.", Data: gin.H{"error": err}})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "Obtain Success", Data: tmpInof})
}

func ClearTmp(ctx *gin.Context) {
	err := tmp.Clear()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Cleanup Failure", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "Cleanup Successfully"})
}
