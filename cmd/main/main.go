package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fliqt/config"
	"fliqt/internal/api/interview"
	"github.com/gin-gonic/gin"

	"fliqt/internal/lib/db"
)

const (
	defReadTimeout    = 10 * time.Second
	defWriteTimeout   = 30 * time.Second
	defMaxHeaderBytes = 1 << 20
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Set Gin mode based on debug flag
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
		gin.DefaultErrorWriter = io.Discard
	}

	// Initialize database connection
	db.Init(cfg)

	// Setup router
	router := gin.Default()
	rg := router.Group("/api")
	interview.Route(rg)

	// Initialize HTTP server
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%s", "8080"),
		Handler:        router,
		ReadTimeout:    defReadTimeout,
		WriteTimeout:   defWriteTimeout,
		MaxHeaderBytes: defMaxHeaderBytes,
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("🚀 Server is starting on port %s in %s mode...\n", "8080", gin.Mode())
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("❌ Server listen error: %v\n", err)
			os.Exit(1) // Exit with non-zero code if server fails to start
		}
	}()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	fmt.Printf("\n🛑 Shutdown signal received: %v\n", sig)

	// Context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), defReadTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("❌ Server Shutdown error: %v\n", err)
	} else {
		fmt.Println("✅ Server gracefully stopped.")
	}

	fmt.Println("👋 Server exiting. Goodbye!")
}
