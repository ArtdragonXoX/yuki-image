package handlers

import (
	"net/http"
	"yuki-image/model"
	"yuki-image/utils"

	"github.com/gin-gonic/gin"
)

func GetTmpInfo(ctx *gin.Context) {
	var tmpInof model.TmpInfo
	size, err := utils.GetDirSize("tmp")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "获取文件大小失败", Data: gin.H{"error": err}})
		return
	}
	count, err := utils.GetFileCount("tmp")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "获取文件数量失败", Data: gin.H{"error": err}})
		return
	}
	tmpInof.Size = size
	tmpInof.Count = count

	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "获取成功", Data: tmpInof})
}

func ClearTmp(ctx *gin.Context) {
	err := utils.DeleteDir("tmp")
	_ = utils.EnsureDir("tmp")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "清理失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "清理成功"})

}
