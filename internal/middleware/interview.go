package middleware

import (
	"net/http"

	"fliqt/internal/lib"
	"fliqt/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type InterviewMiddleware struct {
	repo   *repository.InterviewRepository
	logger *zerolog.Logger
}

func NewInterviewMiddleware(
	repo *repository.InterviewRepository,
	logger *zerolog.Logger,
) *InterviewMiddleware {
	return &InterviewMiddleware{
		repo:   repo,
		logger: logger,
	}
}

// ListInterviews is a middleware for listing all jobs.
func (h *InterviewMiddleware) ListInterviews(ctx *gin.Context) {
	var filterParams repository.InterviewFilterParams
	if err := ctx.ShouldBindQuery(&filterParams); err != nil {
		ctx.Error(err)
		return
	}

	filterParams.Normalize()

	accounts, err := h.repo.ListInterviews(ctx, filterParams)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

// GetInterviews is a middleware for getting job details.
func (h *InterviewMiddleware) GetInterviews(ctx *gin.Context) {
	tracerCtx, span := tracer.Start(ctx.Request.Context(), lib.GetSpanNameFromCaller())
	defer span.End()

	id := ctx.Param("id")
	if id == "" {
		ctx.Error(ErrNotFound)
		return
	}
	account, err := h.repo.GetInterviewByID(tracerCtx, id)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (h *InterviewMiddleware) CreateInterviews(ctx *gin.Context) {
	var req repository.CreateInterviewDTO
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	if err := req.Validate(); err != nil {
		ctx.Error(err)
		return
	}

	job, err := h.repo.CreateInterview(ctx, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, job)
}

func (h *InterviewMiddleware) UpdateInterviews(ctx *gin.Context) {
	var req repository.UpdateInterviewDTO
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ID := ctx.Param("id")
	if ID == "" {
		ctx.Error(ErrNotFound)
		return
	}

	if err := req.Validate(); err != nil {
		ctx.Error(err)
		return
	}

	job, err := h.repo.UpdateInterview(ctx, ID, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, job)
}

func (h *InterviewMiddleware) DeleteInterviews(ctx *gin.Context) {
	ID := ctx.Param("id")
	if ID == "" {
		ctx.Error(ErrNotFound)
		return
	}

	err := h.repo.DeleteInterview(ctx, ID)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
