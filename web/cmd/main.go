package main

import (
	"context"
	"github.com/Seven4X/link/web/app/comment"
	"github.com/Seven4X/link/web/app/topic"
	"github.com/Seven4X/link/web/lib/config"
	"github.com/Seven4X/link/web/lib/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Seven4X/link/web/app/link"
	"github.com/Seven4X/link/web/app/user"
	"github.com/Seven4X/link/web/app/vote"
	setup "github.com/Seven4X/link/web/lib/echo"
)

func main() {
	e := setup.NewEcho()
	//站点静态文件build输出目录
	path := config.GetString("site.path")
	e.File("/*", path+"/index.html")
	e.Static("/static", path+"/static")
	e.File("/favicon.ico", path+"/favicon.ico")
	e.File("/manifest.json", path+"/manifest.json")
	//用于证书认证 https://letsencrypt.org/zh-cn/docs/challenge-types/
	e.Static("/wellknown", path+"/wellknown")
	// 初始化模块
	topic.Router(e)
	user.Router(e)
	link.Router(e)
	vote.Router(e)
	comment.Router(e)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGTSTP)
		<-sigint
		// We received an interrupt signal, shut down.
		if err := e.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()
	go func(c *echo.Echo) {
		e.Logger.Fatal(e.Start(":80"))
	}(e)
	e.AutoTLSManager.Cache = autocert.DirCache("/root/www/.cache")
	if err := e.StartAutoTLS(":443"); err != http.ErrServerClosed {
		// Error starting or closing listener:
		e.Logger.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-idleConnsClosed
	//
	util.DumpCuckooFilter()
	log.Printf("app shutdown")
}
