package migration

import (
	"time"

	"fliqt/internal/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate01() *gormigrate.Migration {
	type Interview struct {
		model.Base
		CandidateName string                `gorm:"not null"`
		Position      string                `gorm:"not null"`
		Status        model.InterviewStatus `gorm:"default:0;comment:'0:Pending, 1:Scheduled, 2:In Progress, 3:Completed, 4:Rejected, 5:Offered'"`
		ScheduledTime *time.Time            `gorm:"index:idx_interview_scheduled_time"`
		Notes         string                `gorm:"type:text"`
	}

	// Seed data for interviews
	seedData := []Interview{
		{
			Base:          model.Base{ID: "interview-1"},
			CandidateName: "John Doe",
			Position:      "Backend Engineer",
			Status:        model.InterviewScheduled,
			ScheduledTime: func() *time.Time { t := time.Now().Add(24 * time.Hour); return &t }(),
			Notes:         "First round technical interview",
		},
		{
			Base:          model.Base{ID: "interview-2"},
			CandidateName: "Jane Smith",
			Position:      "Frontend Engineer",
			Status:        model.InterviewPending,
			Notes:         "Resume reviewed, pending schedule",
		},
		{
			Base:          model.Base{ID: "interview-3"},
			CandidateName: "Michael Brown",
			Position:      "DevOps Engineer",
			Status:        model.InterviewCompleted,
			Notes:         "Final round completed",
		},
	}

	return &gormigrate.Migration{
		ID: "0001",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.Migrator().CreateTable(&Interview{}); err != nil {
				return err
			}

			if err := tx.Create(&seedData).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable(&Interview{}); err != nil {
				return err
			}
			return nil
		},
	}
}
