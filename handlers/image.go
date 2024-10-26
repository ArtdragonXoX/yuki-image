package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"yuki-image/conf"
	"yuki-image/db"
	"yuki-image/handlers/image"
	"yuki-image/model"
	"yuki-image/utils"

	"github.com/gin-gonic/gin"
)

func UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: gin.H{"error": err}})
		return
	}

	album_id := ctx.PostForm("album_id")
	album_uid, err := strconv.ParseUint(album_id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: gin.H{"error": err}})
		return
	}

	dst := fmt.Sprintf("tmp/%s", file.Filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: gin.H{"error": err}})
		return
	}

	format := utils.GetImageFormat(dst)
	album, err := db.SelectAlbum(album_uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: gin.H{"error": err}})
		return
	}

	if !utils.ContainsFormatSupport(album.FormatSupport, format) {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件格式不支持", Data: nil})
		return
	}

	hash, err := utils.GetImageHash(dst)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: gin.H{"error": err}})
		return
	}
	newFileName := fmt.Sprintf("%s.%s", hash, utils.GetImageFormatName(format))
	newFilePath := fmt.Sprintf("%s/%s", album.Name, newFileName)
	localFilePath := fmt.Sprintf("%s/%s", conf.Conf.Server.Path, newFilePath)
	switch format {
	case model.JPEG:
		image.ManipulateJPEG(dst, localFilePath, int(album.MaxHeight), int(album.MaxWidth))
	case model.PNG:
		image.ManipulatePNG(dst, localFilePath, int(album.MaxHeight), int(album.MaxWidth))
	case model.GIF:
		image.ManipulateGIF(dst, localFilePath, int(album.MaxHeight), int(album.MaxWidth))
	default:
	}

	size, err := utils.GetImageSize(localFilePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: gin.H{"error": err}})
		return
	}

	image := model.Image{
		Id:         utils.GetRandKey(),
		Name:       newFileName,
		AlbumId:    album_uid,
		Pathname:   newFilePath,
		OriginName: file.Filename,
		Size:       size,
		Mimetype:   utils.GetImageFormatName(format),
	}
	err = db.InsertImage(image)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: gin.H{"error": err}})
		return
	}

	image, err = db.SelectImage(image.Id)

	image.Url = utils.GetImageUrl(image)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: gin.H{"error": err}})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "文件上传成功", Data: gin.H{"image": image}})
}

func SelectImage(ctx *gin.Context) {
	imageId := ctx.Param("id")
	if imageId == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误", Data: nil})
		return
	}
	image, err := db.SelectImage(imageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: gin.H{"error": err}})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: image})
}

func DeleteImage(ctx *gin.Context) {
	imageId := ctx.Param("id")
	if imageId == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "参数错误", Data: nil})
		return
	}
	image, err := db.SelectImage(imageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: gin.H{"error": err}})
		return
	}

	err = db.DeleteImage(imageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "删除失败", Data: gin.H{"error": err}})
		return
	}

	pathname := fmt.Sprintf("%s/%s", conf.Conf.Server.Path, image.Pathname)
	err = os.Remove(pathname)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "删除失败", Data: gin.H{"error": err}})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "删除成功", Data: nil})
}
