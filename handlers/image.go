package handlers

import (
	"fmt"
	"net/http"
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
