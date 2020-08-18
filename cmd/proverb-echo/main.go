package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/stevedesilva/my-proverb-app/pkg/handler"

	"github.com/labstack/gommon/log"

	"github.com/stevedesilva/my-proverb-app/api/server"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Info("My Proverb application!")
	// TODO Metric (e.g. NewRelic)
	// TODO Database
	e := echo.New()
	startServer(e)

	wg := &sync.WaitGroup{}
	listenForShutdownSignal(wg)
	// wait until shutdown signal sent the
	wg.Wait()

	gracefulShutDown(e)

	log.Info("Exiting application!")
}

func startServer(e *echo.Echo) {
	log.Info("Starting server ...")
	// serve static content
	e.Static("/swaggerui", "api/swaggerui")
	e.Static("/redoc", "api/redoc")

	// open api codegen used to generate client and server stubs
	server.RegisterHandlers(e, handler.NewProverbsService())

	// TODO port from environment var
	httpPort := ":8081"
	go func() {
		if err := e.Start(httpPort); err != nil {
			e.Logger.Fatal("Echo error: ", err) // TODO
			log.Error("Proverb App Error: ", err)
		}
	}()
}

// need to pass pointer else func will only receive a copy
func listenForShutdownSignal(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM)
		s := <-sig
		log.Info("Signal received, ", s)
	}()
}

func gracefulShutDown(e *echo.Echo) {
	log.Info("Gracefully shutting down server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Errorf("Graceful shutdown failed.", err)
		return
	}
	log.Info("Graceful shutdown complete")
}
