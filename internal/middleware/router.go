package middleware

import (
	"fliqt/internal/model"
	"fliqt/internal/repository"
	"fliqt/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"fliqt/config"
)

func NewRouter(
	cfg *config.Config,
	app *gin.Engine,
	logger *zerolog.Logger,
	interviewRepo *repository.InterviewRepository,
	authService service.AuthServiceInterface,
) {
	r := app.Group("/api")

	r.Use(AuthMiddleware(authService, []model.UserRole{model.RoleHR}))

}
