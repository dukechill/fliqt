package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fliqt/config"
	"fliqt/internal/api/interview"
	"fliqt/internal/lib"
	"github.com/gin-gonic/gin"

	"fliqt/internal/lib/db"
)

const (
	defReadTimeout    = 10 * time.Second
	defWriteTimeout   = 30 * time.Second
	defMaxHeaderBytes = 1 << 20
)

func main() {
	cfg := config.NewConfig()
	logger := lib.NewLogger(cfg)

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
		gin.DefaultErrorWriter = io.Discard
	}
	app := gin.Default()

	db.Init()

	router := gin.Default()

	rg := router.Group("/api")

	interview.Route(rg)

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", "8080"),
		Handler:        router,
		ReadTimeout:    defReadTimeout,
		WriteTimeout:   defWriteTimeout,
		MaxHeaderBytes: defMaxHeaderBytes,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server listen error: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	fmt.Printf("Shutdown Server with signal %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), defReadTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown err: %v\n", err)
	}
	fmt.Println("Server exiting")
}
