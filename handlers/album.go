package handlers

import (
	"net/http"
	"strconv"
	"yuki-image/db"
	"yuki-image/model"

	"github.com/gin-gonic/gin"
)

func InsertAlbum(ctx *gin.Context) {
	var album model.Album
	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误"})
		return
	}
	id, err := db.InsertAlbum(album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "插入失败"})
		return
	}
	album, err = db.SelectAlbum(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "插入失败"})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "插入成功", Data: gin.H{"album": album}})
}

func SelectAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误"})
		return
	}
	album, err := db.SelectAlbum(uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败"})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"album": album}})
}

func SelectAllAlbum(ctx *gin.Context) {
	albums, err := db.SelectAllAlbum()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败"})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"album": albums}})
}
