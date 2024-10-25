package handlers

import (
	"net/http"
	"yuki-image/db"
	"yuki-image/model"

	"github.com/gin-gonic/gin"
)

func InsertAlbum(ctx *gin.Context) {
	var album model.Album
	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 400, Msg: "参数错误"})
		return
	}
	id, err := db.InsertAlbum(album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 401, Msg: "插入失败"})
		return
	}
	album, err = db.SelectAlbum(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 401, Msg: "插入失败"})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 201, Msg: "插入成功", Data: album})
}
