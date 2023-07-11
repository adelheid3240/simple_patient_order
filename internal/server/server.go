package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simplepatientorder/config"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(config *config.Config) {
	if config.Gin.Mode == gin.DebugMode || config.Gin.Mode == gin.TestMode || config.Gin.Mode == gin.ReleaseMode {
		gin.SetMode(config.Gin.Mode)
	}

	ginEngine := gin.Default()

	ginEngine.GET("/", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	ginEngine.NoRoute(notFound)

	srv := &http.Server{
		Addr:    ":" + config.Server.Port,
		Handler: ginEngine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	gracefulShutdown(srv, config.Server.ShutdownTimeoutSec)
}

func gracefulShutdown(srv *http.Server, shutdownTimeoutSec int) {
	quit := make(chan os.Signal, 3)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(shutdownTimeoutSec)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

func notFound(c *gin.Context) {
	if c.Request != nil && c.Request.Method == http.MethodPost {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err == nil {
			log.Printf("jsonData: %s\n", jsonData)
		}
	}

	c.AbortWithStatus(http.StatusNotFound)
}