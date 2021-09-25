package main

import (
	"context"
	"errors"
	api "gopher_lee/file/internal/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	//create Logger
	logger := logrus.NewEntry(logrus.New())

	// create handler instance
	uploadHandler := api.NewUpload(logger)
	downloadHandler := api.NewDownload(logger)

	// mux
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadHandler.Single)
	mux.HandleFunc("/download", downloadHandler.Single)

	// fileserver
	fs := http.FileServer(http.Dir("./file/internal/web/static"))
	mux.Handle("/", fs)

	// create Server Instace
	srv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("could not open Server With 3000 port error:%v", err)
		}
	}()
	logger.Infoln("starting server with 3000 port...")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Infoln("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalln("Server forced to shutdown:", err)
	}

	logger.Infoln("Server exiting")
}
