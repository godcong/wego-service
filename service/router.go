package service

import (
	"github.com/gin-gonic/gin"
	_ "github.com/godcong/wego-manager-service/statik"
	"github.com/rakyll/statik/fs"
	"io"
	"log"
	"net/http"
	"path"
)

// Router ...
func Router(eng *gin.Engine) {
	verV0 := "v0"
	staticFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	eng.NoRoute(func(ctx *gin.Context) {
		_, file := path.Split(ctx.Request.URL.Path)
		open, err := staticFS.Open(ctx.Request.URL.Path)
		if file == "" || err != nil {
			open, err = staticFS.Open("/index.html")
			if err != nil {
				panic(err)
			}
		}
		_, err = io.Copy(ctx.Writer, open)
	})

	eng.Use(AccessControlAllow)
	g0 := eng.Group(verV0)

	g0.GET("inited", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Code": 0})
	})

	////登录
	//g0.POST("login", LoginPOST(verV0))
	////组织注册
	//g0.POST("register", RegisterPOST(verV0))
	//
	//g0.POST("genesis", GenesisGet(verV0))

}

func isInstalled() bool {
	return false
}

func AccessControlAllow(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("origin")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
		"Accept-Encoding, X-CSRF-Token, Authorization")
	if ctx.Request.Method == "OPTIONS" {
		ctx.String(200, "ok")
		return
	}
	ctx.Next()
}
