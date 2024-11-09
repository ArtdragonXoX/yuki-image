package static

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var stic embed.FS

func InitStatic(e *gin.Engine) {
	e.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/web/")
	})

	web := e.Group("/web")
	SiglePageAppFS(web, os.DirFS("static/dist"))
}

func newFSHandler(fileSys fs.FS) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fp := strings.Trim(ctx.Param("filepath"), "/")
		f, err := fileSys.Open(fp)
		if err != nil {
			fp = ""
		} else {
			f.Close()
		}
		log.Println(fp)
		ctx.FileFromFS(fp, http.FS(fileSys))
	}
}

func SiglePageAppFS(r *gin.RouterGroup, fileSys fs.FS) {
	h := newFSHandler(fileSys)
	r.GET("/*filepath", h)
	r.HEAD("/*filepath", h)
}
