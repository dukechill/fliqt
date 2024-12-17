package model

import "time"

type InterviewStatus int

const (
	InterviewPending    InterviewStatus = iota // Pending
	InterviewScheduled                         // Scheduled
	InterviewInProgress                        // In Progress
	InterviewCompleted                         // Completed
	InterviewRejected                          // Rejected
	InterviewOffered                           // Offered
)

func (s InterviewStatus) String() string {
	return [...]string{
		"Pending",
		"Scheduled",
		"In Progress",
		"Completed",
		"Rejected",
		"Offered",
	}[s]
}

type Interview struct {
	Base

	CandidateName string          `gorm:"not null"`
	Position      string          `gorm:"not null"`
	Status        InterviewStatus `gorm:"default:0;comment:'0:Pending, 1:Scheduled, 2:In Progress, 3:Completed, 4:Rejected, 5:Offered'"`
	ScheduledTime *time.Time      `gorm:"index:idx_interview_scheduled_time"`
	Notes         string          `gorm:"type:text"`
}
