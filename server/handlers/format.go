package handlers

import (
	"log"
	"net/http"
	"strconv"
	iformat "yuki-image/internal/format"
	"yuki-image/internal/model"

	"github.com/gin-gonic/gin"
)

// func InsertFormat(ctx *gin.Context) {
// 	var format model.Format
// 	err := ctx.ShouldBindJSON(&format)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, model.Response{Code: 400, Msg: "JSON error"})
// 		return
// 	}
// 	id, err := db.InsertFormat(&format)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, model.Response{Code: 401, Msg: "Insert Failure"})
// 		return
// 	}
// 	format, err = db.SelectFormat(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, model.Response{Code: 401, Msg: "Insert Failure"})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, model.Response{Code: 201, Msg: "插入成功", Data: gin.H{"format": format}})
// }

func SelectFormat(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	log.Println(id)
	format, err := iformat.Select(uint64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败"})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"format": format}})
}

func SelectAllFormat(ctx *gin.Context) {
	formats, err := iformat.SelectAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "查询失败"})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 1, Msg: "查询成功", Data: gin.H{"formats": formats}})
}
