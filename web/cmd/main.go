package main

import (
	"context"
	"github.com/Seven4X/link/web/app"
	"github.com/Seven4X/link/web/app/util"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	e := app.SetupEcho()

	c := app.StartJob()

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
		c.Stop()
		util.DumpCuckooFilter()
		close(idleConnsClosed)
	}()
	go func(c *echo.Echo) {
		e.Logger.Fatal(e.Start(":8080"))
	}(e)
	//e.AutoTLSManager.Cache = autocert.DirCache("/root/www/.cache")
	//if err := e.StartAutoTLS(":443"); err != http.ErrServerClosed {
	//	// Error starting or closing listener:
	//	e.Logger.Fatalf("HTTP server ListenAndServe: %v", err)
	//}
	<-idleConnsClosed
	//
	log.Printf("app shutdown")
}
