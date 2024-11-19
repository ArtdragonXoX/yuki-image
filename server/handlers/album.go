package handlers

import (
	"net/http"
	"strconv"
	ialbum "yuki-image/internal/album"
	"yuki-image/internal/conf"
	"yuki-image/internal/model"

	"github.com/gin-gonic/gin"
)

func InsertAlbum(ctx *gin.Context) {
	var album model.Album
	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	id, err := ialbum.Insert(album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Insert Failure"})
		return
	}
	album, err = ialbum.Select(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Insert Failure", Data: err})
		return
	}

	ctx.JSON(http.StatusCreated, model.Response{Code: 1, Msg: "插入成功", Data: album})
}

func UpdateAlbum(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	idHasValue := err == nil && id > 0

	var album model.Album
	err = ctx.ShouldBindJSON(&album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	if idHasValue {
		album.Id = (uint64(id))
	} else if param != "" {
		album.Name = param
	}
	err = ialbum.Update(album)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "更新失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "更新成功", Data: nil})
}

func SelectAlbum(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	idHasValue := err == nil && id > 0
	var album model.Album
	if !idHasValue {
		album, err = ialbum.SelectFromName(param)
	} else {
		album, err = ialbum.Select(uint64(id))
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: album})
}

func SelectAllAlbum(ctx *gin.Context) {
	albums, err := ialbum.SelectAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: albums})
}

func ClearAlbum(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	idHasValue := err == nil && id > 0
	if idHasValue {
		err = ialbum.Clear(uint64(id))
	} else {
		err = ialbum.ClearFromName(param)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "清空失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "清空成功", Data: nil})
}

func DeleteAlbum(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	idHasValue := err == nil && id > 0
	if idHasValue {
		err = ialbum.Delete(uint64(id))
	} else {
		err = ialbum.DeleteFromName(param)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "删除失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "删除成功", Data: nil})
}

func GetAllAlbumSize(ctx *gin.Context) {
	size, err := ialbum.GetAllSize()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: size})
}

func GetAllAlbumCount(ctx *gin.Context) {
	count, err := ialbum.GetAllCount()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: count})
}

func GetAlbumSize(ctx *gin.Context) {
	var size uint64
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	idHasValue := err == nil && id > 0
	if idHasValue {
		size, err = ialbum.GetSize(uint64(id))
	} else {
		size, err = ialbum.GetSizeFromName(param)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: size})
}

func GetAlbumCount(ctx *gin.Context) {
	var count uint64
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	idHasValue := err == nil && id > 0
	if idHasValue {
		count, err = ialbum.GetCount(uint64(id))
	} else {
		count, err = ialbum.GetCountFromName(param)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: count})
}

func InsertFormatSupport(ctx *gin.Context) {
	var formatSupport model.FormatSupport
	err := ctx.ShouldBindJSON(&formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	err = ialbum.InsertFormatSupport(formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Insert Failure", Data: err})
		return
	}
	ctx.JSON(http.StatusCreated, model.Response{Code: 1, Msg: "插入成功", Data: nil})
}

func SelectFormatSupport(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	idHasValue := err == nil && id > 0
	var format_support []model.Format
	if !idHasValue {
		format_support, err = ialbum.SelectFormatSupportFromName(param)
	} else {
		format_support, err = ialbum.SelectFormatSupport(uint64(id))
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: format_support})
}

func SelectImageFromAlbum(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	idHasValue := err == nil && id > 0
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "Query error"})
		return
	}
	upage := uint64(page)
	size, err := strconv.Atoi(ctx.Query("size"))
	if err != nil {
		size = conf.Conf.Image.ImageListDefalutSize
	}
	usize := uint64(size)
	var imageList model.ImageList
	if !idHasValue {
		imageList, err = ialbum.SelectImageFromName(param, upage, usize)
	} else {
		imageList, err = ialbum.SelectImage(uint64(id), upage, usize)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: imageList})
}

func DeleteFormatSupport(ctx *gin.Context) {
	var formatSupport model.FormatSupport
	err := ctx.ShouldBindJSON(&formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "JSON error"})
		return
	}
	err = ialbum.DeleteFormatSupport(formatSupport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "删除失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "删除成功", Data: nil})
}

func SelectCountStatistics(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	idHasValue := err == nil && id > 0
	var startDate model.CustomTime
	err = startDate.FromString(ctx.Query("start-time"))
	if err != nil {
		startDate.Now()
	}
	var endDate model.CustomTime
	err = endDate.FromString(ctx.Query("end-time"))
	if err != nil {
		endDate.Now()
	}
	var statistics = make(map[string]uint64)
	if !idHasValue {
		statistics, err = ialbum.GetCountStatisticsFromName(param, startDate.Time, endDate.Time)
	} else {
		statistics, err = ialbum.GetCountStatistics(uint64(id), startDate.Time, endDate.Time)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: statistics})
}

func SelectAllCountStatistics(ctx *gin.Context) {
	var startDate model.CustomTime
	err := startDate.FromString(ctx.Query("start-time"))
	if err != nil {
		startDate.Now()
	}
	var endDate model.CustomTime
	err = endDate.FromString(ctx.Query("end-time"))
	if err != nil {
		endDate.Now()
	}
	statistics, err := ialbum.GetAllCountStatistics(startDate.Time, endDate.Time)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: statistics})
}
