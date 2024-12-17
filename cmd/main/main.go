package main

import (
	"io"

	"fliqt/config"
	"fliqt/internal/lib"
	"fliqt/internal/repository"
	"fliqt/internal/service"
	"github.com/gin-gonic/gin"

	"fliqt/internal/middleware"
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

	db, err := lib.NewGormDB(cfg)
	if err != nil {
		panic(err)
	}

	//redisClient, err := lib.NewClient(cfg)
	//if err != nil {
	//	panic(err)
	//}

	// Initialize repositories
	InterviewRepo := repository.NewInterviewRepository(db, logger)

	// Initialize services
	authService := service.NewAuthService(db)

	// OpenTelemetry tracing, can be ignored when there's no setup for tracing when developing locally.
	if err := lib.InitTracer(cfg); err != nil {
		logger.Info().Msgf("Failed to initialize tracer: %v", err)
	}

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

	if err := app.Run(); err != nil {
		panic(err)
	}
}
