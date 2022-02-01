package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"server/global"
	"server/router"
	"syscall"
)

func startWebServer() error {
	routersInit := router.InitRouters()

	addr := fmt.Sprintf("%s:%d", global.CONF.WebServer.Address, global.CONF.WebServer.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: routersInit,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			global.LOG.Fatal("ListenAndServe Failed", zap.Error(err))
		}
	}()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	err := server.Shutdown(nil)
	if err != nil {
		global.LOG.Warn("Shutdown Failed", zap.Error(err))
	}

	return err
}
