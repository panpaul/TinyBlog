package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"server/global"
	"server/router"
	"syscall"
	"time"
)

func startWebServer() error {
	routersInit := router.InitRouters()

	addr := fmt.Sprintf("%s:%d", global.CONF.WebServer.Address, global.CONF.WebServer.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: routersInit,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOG.Fatal("ListenAndServe Failed", zap.Error(err))
		}
	}()

	// Handle SIGINT and SIGTERM.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.LOG.Info("Shutting down Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		global.LOG.Fatal("Server Shutdown Failed", zap.Error(err))
	}

	return err
}
