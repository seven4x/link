package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/app"
	"github.com/seven4x/link/job"

	_ "github.com/mattn/go-sqlite3"
	flag "github.com/spf13/pflag"
)

func main() {

	flag.Parse()
	e := echo.New()
	err := BootApp(e)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	c := job.StartJob()

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
		app.DumpCuckooFilter()
		close(idleConnsClosed)
	}()
	go func(c *echo.Echo) {
		e.Logger.Fatal(e.Start(":8088"))
	}(e)

	<-idleConnsClosed
	//
	log.Printf("app shutdown")
}
