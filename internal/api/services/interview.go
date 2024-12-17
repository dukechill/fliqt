package services

import (
	"context"
	"errors"

	"fliqt/internal/model"
	"github.com/jinzhu/gorm"
)

type InterviewService struct {
	db *gorm.DB
}

func NewInterviewService(
	db *gorm.DB,
) *InterviewService {
	return &InterviewService{
		db: db,
	}
}

var (
	ErrInvalidInterviewStatus = errors.New("invalid interview status")
)

type InterviewFilterParams struct {
	model.PaginationParams

	CandidateName string `form:"candidate_name,omitempty"`
	Position      string `form:"position,omitempty"`
	Status        int    `form:"status,omitempty"`
}

type InterviewResponseDTO struct {
	ID            string `json:"id"`
	CandidateName string `json:"candidate_name"`
	Position      string `json:"position"`
	Status        string `json:"status"`
	ScheduledTime string `json:"scheduled_time,omitempty"`
	Notes         string `json:"notes,omitempty"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

// ListInterviews returns a list of interviews
func (r *InterviewService) ListInterviews(filterParams InterviewFilterParams) (model.PaginationResponse[InterviewResponseDTO], error) {
	var interviews []InterviewResponseDTO
	query := r.db.Model(&model.Interview{}).Order("id DESC")

	if filterParams.CandidateName != "" {
		query = query.Where("candidate_name LIKE ?", "%"+filterParams.CandidateName+"%")
	}
	if filterParams.Position != "" {
		query = query.Where("position = ?", filterParams.Position)
	}
	if filterParams.Status != 0 {
		query = query.Where("status = ?", filterParams.Status)
	}

	var total int64
	var result model.PaginationResponse[InterviewResponseDTO]

	if err := query.Count(&total).Error; err != nil {
		return result, err
	}

	if filterParams.NextToken != "" {
		query = query.Where("id < ?", filterParams.NextToken)
	}

	query = query.Limit(filterParams.PageSize)

	if err := query.Find(&interviews).Error; err != nil {
		return result, err
	}

	result.Total = total
	result.Items = interviews

	if len(interviews) > 0 && len(interviews) == filterParams.PageSize {
		result.NextToken = interviews[len(interviews)-1].ID
	}

	return result, nil
}

// GetInterviewByID returns an interview by its ID
func (r *InterviewService) GetInterviewByID(ID string) (*model.Interview, error) {
	var interview model.Interview
	if err := r.db.Where("id = ?", ID).First(&interview).Error; err != nil {
		return nil, err
	}

	return &interview, nil
}

type InterviewValidator interface {
	Validate() error
}

type CreateInterviewDTO struct {
	CandidateName string `json:"candidate_name" binding:"required"`
	Position      string `json:"position" binding:"required"`
	Status        int    `json:"status" binding:"required"`
	ScheduledTime string `json:"scheduled_time,omitempty"`
	Notes         string `json:"notes,omitempty"`
}

func (dto CreateInterviewDTO) Validate() error {
	if dto.Status < 0 || dto.Status > 5 {
		return ErrInvalidInterviewStatus
	}
	return nil
}

// CreateInterview creates a new interview
func (r *InterviewService) CreateInterview(dto CreateInterviewDTO) (*model.Interview, error) {
	interview := model.Interview{
		CandidateName: dto.CandidateName,
		Position:      dto.Position,
		Status:        model.InterviewStatus(dto.Status),
		Notes:         dto.Notes,
	}

	if err := r.db.Create(&interview).Error; err != nil {
		return nil, err
	}

	return &interview, nil
}

type UpdateInterviewDTO struct {
	Status        int    `json:"status" binding:"required"`
	ScheduledTime string `json:"scheduled_time,omitempty"`
	Notes         string `json:"notes,omitempty"`
}

func (dto UpdateInterviewDTO) Validate() error {
	if dto.Status < 0 || dto.Status > 5 {
		return ErrInvalidInterviewStatus
	}
	return nil
}

// UpdateInterview updates an interview
func (r *InterviewService) UpdateInterview(ctx context.Context, ID string, dto UpdateInterviewDTO) (*model.Interview, error) {
	if err := r.db.Model(&model.Interview{}).Where("id = ?", ID).UpdateColumns(dto).Error; err != nil {
		return nil, err
	}

	return r.GetInterviewByID(ID)
}

// DeleteInterview deletes an interview
func (r *InterviewService) DeleteInterview(ID string) error {
	if err := r.db.Where("id = ?", ID).Delete(&model.Interview{}).Error; err != nil {
		return err
	}

	return nil
}
