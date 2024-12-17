package interview

import (
	"fmt"
	"net/http"

	"fliqt/internal/api/services"
	"fliqt/internal/lib/db"
	"github.com/gin-gonic/gin"
)

func ListInterviews(ctx *gin.Context) {
	var filterParams services.InterviewFilterParams
	if err := ctx.ShouldBindQuery(&filterParams); err != nil {
		ctx.Error(err)
		return
	}

	filterParams.Normalize()
	conn := db.DBGorm
	srv := services.NewInterviewService(conn)
	accounts, err := srv.ListInterviews(ctx, filterParams)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

// GetInterviews is for getting interview details.
func GetInterviews(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.Error(fmt.Errorf("invalid id"))
		return
	}
	conn := db.DBGorm
	srv := services.NewInterviewService(conn)
	interviewer, err := srv.GetInterviewByID(ctx, id)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, interviewer)
}

func CreateInterviews(ctx *gin.Context) {
	var req services.CreateInterviewDTO
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	if err := req.Validate(); err != nil {
		ctx.Error(err)
		return
	}
	conn := db.DBGorm
	srv := services.NewInterviewService(conn)

	interviewer, err := srv.CreateInterview(ctx, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, interviewer)
}

func UpdateInterviews(ctx *gin.Context) {
	var req services.UpdateInterviewDTO
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ID := ctx.Param("id")
	if ID == "" {
		ctx.Error(fmt.Errorf("invalid id"))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.Error(err)
		return
	}

	conn := db.DBGorm
	srv := services.NewInterviewService(conn)

	job, err := srv.UpdateInterview(ctx, ID, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, job)
}

func DeleteInterviews(ctx *gin.Context) {
	ID := ctx.Param("id")
	if ID == "" {
		ctx.Error(fmt.Errorf("invalid id"))
		return
	}
	conn := db.DBGorm
	srv := services.NewInterviewService(conn)
	err := srv.DeleteInterview(ctx, ID)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
