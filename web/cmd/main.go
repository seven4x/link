package main

import (
	"context"
	"github.com/Seven4X/link/web/library/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	account "github.com/Seven4X/link/web/app/account/server"
	link "github.com/Seven4X/link/web/app/link/server"
	topic "github.com/Seven4X/link/web/app/topic/server/http"
	setup "github.com/Seven4X/link/web/library/echo"
	"github.com/labstack/echo/v4"
)

func main() {
	e := setup.NewEcho()
	// Routes
	e.GET("/", hello)

	// 初始化模块
	topic.Router(e)
	account.Router(e)
	link.Router(e)

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

	if err := e.Start(":1323"); err != http.ErrServerClosed {
		// Error starting or closing listener:
		e.Logger.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-idleConnsClosed
	//
	util.DumpCuckooFilter()
	log.Printf("app shutdown")
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
