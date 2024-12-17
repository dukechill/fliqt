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
	"fliqt/internal/lib"
	"fliqt/internal/repository"
	"fliqt/internal/service"
	"github.com/gin-gonic/gin"

	"fliqt/internal/lib/db"
	"fliqt/internal/middleware"
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

	// Initialize repositories
	InterviewRepo := repository.NewInterviewRepository(db, logger)

	// Initialize services
	authService := service.NewAuthService(db)

	app.Use(middleware.Logger(logger))
	app.Use(middleware.ErrorMiddleware(logger))
	app.NoRoute(middleware.NotFoundHandler())

	middleware.NewRouter(
		cfg,
		app,
		logger,
		InterviewRepo,
		authService,
	)

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
			logs.Systemf("Server listen error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	logs.Systemf("Shutdown Server with signal %v", sig)

	ctx, cancel := context.WithTimeout(context.Background(), defReadTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logs.Systemf("Server Shutdown err: %v", err)
	}
	logs.Systemf("Server exiting")
}
