package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"yuki-image/conf"
	"yuki-image/db"
	"yuki-image/model"
	"yuki-image/utils"

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
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "插入失败", Data: gin.H{"error": err}})
		return
	}
	pathname := fmt.Sprintf("%s/%s", conf.Conf.Server.Path, album.Name)
	
	err = utils.EnsureDir(pathname)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "创建目录失败", Data: gin.H{"error": err}})
		return
	}

	ctx.JSON(http.StatusCreated, model.Response{Code: 1, Msg: "插入成功", Data: gin.H{"album": album}})
}

func UpdateAlbum(ctx *gin.Context) {
	var album model.Album
	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误"})
		return
	}
	err = db.UpdateAlbum(album, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "更新失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "更新成功", Data: nil})
}

func SelectAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误"})
		return
	}
	album, err := db.SelectAlbum(uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"album": album}})
}

func SelectAllAlbum(ctx *gin.Context) {
	albums, err := db.SelectAllAlbum()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"album": albums}})
}

func InsertFormatSupport(ctx *gin.Context) {
	var formatSupport model.FormatSupport
	err := ctx.ShouldBindJSON(&formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误"})
		return
	}
	err = db.InsertFormatSupport(formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "插入失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusCreated, model.Response{Code: 1, Msg: "插入成功", Data: nil})
}

func SelectFormatSupport(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误"})
		return
	}
	format_support, err := db.SelectFormatSupport(uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"format_support": format_support}})
}

func DeleteFormatSupport(ctx *gin.Context) {
	var formatSupport model.FormatSupport
	err := ctx.ShouldBindJSON(&formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误"})
		return
	}
	err = db.DeleteFormatSupport(formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "删除失败", Data: gin.H{"error": err}})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "删除成功", Data: nil})
}
