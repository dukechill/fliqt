package repository

import (
	"testing"

	"github.com/pkg/errors"
)

func TestCreateInterviewDTOValidate(t *testing.T) {
	dto := CreateInterviewDTO{
		CandidateName: "John Doe",
		Position:      "Backend Engineer",
		Status:        1,
	}

	if err := dto.Validate(); err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	dto = CreateInterviewDTO{
		CandidateName: "John Doe",
		Position:      "Backend Engineer",
		Status:        6, // Invalid status
	}

	if err := dto.Validate(); !errors.Is(err, ErrInvalidInterviewStatus) {
		t.Errorf("expected ErrInvalidInterviewStatus, got %v", err)
	}
}
