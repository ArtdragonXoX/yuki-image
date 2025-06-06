package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	iablum "yuki-image/internal/album"
	iimage "yuki-image/internal/image"
	"yuki-image/internal/model"
	"yuki-image/internal/tmp"
	"yuki-image/utils"

	"github.com/gin-gonic/gin"
)

func UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: err})
		return
	}

	album_id := ctx.PostForm("album_id")
	idHasValue := album_id != ""
	album_name := ctx.PostForm("album_name")
	nameHasValue := album_name != ""
	if !idHasValue && !nameHasValue {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "缺少相册信息", Data: err})
		return
	}
	var album_uid uint64
	if !idHasValue {
		album_uid, err = iablum.SelectIdFromName(album_name)
	} else {
		album_uid, err = strconv.ParseUint(album_id, 10, 64)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: err})
		return
	}
	dst := fmt.Sprintf("%s/%s.%s", tmp.TmpPath, utils.GetRandKey(), utils.GetFileExt(file.Filename))
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: err})
		return
	}
	image, err := iimage.Upload(dst, file.Filename, album_uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: err})
		return
	}
	ctx.JSON(http.StatusCreated, model.Response{Code: 1, Msg: "文件上传成功", Data: image})
}

func SelectImage(ctx *gin.Context) {
	imageId := ctx.Param("id")
	if imageId == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error", Data: nil})
		return
	}
	image, err := iimage.Select(imageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: image})
}

func SelectImageFromUrl(ctx *gin.Context) {
	url := ctx.Query("url")
	if url == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Query error", Data: nil})
		return
	}
	image, err := iimage.SelectFromUrl(url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: image})
}

func DeleteImage(ctx *gin.Context) {
	imageId := ctx.Param("id")
	if imageId == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error", Data: nil})
		return
	}

	err := iimage.Delete(imageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "删除失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "删除成功", Data: nil})
}
