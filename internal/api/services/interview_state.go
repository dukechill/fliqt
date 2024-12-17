package services

import (
	"errors"

	"fliqt/internal/model"
)

// InterviewState defines the interface for state machine logic
type InterviewState interface {
	NextState() (InterviewState, error)
	GetStatus() model.InterviewStatus
}

// Pending State
type PendingState struct{}

func (s *PendingState) NextState() (InterviewState, error) { return &ScheduledState{}, nil }
func (s *PendingState) GetStatus() model.InterviewStatus   { return model.InterviewPending }

// Scheduled State
type ScheduledState struct{}

func (s *ScheduledState) NextState() (InterviewState, error) { return &InProgressState{}, nil }
func (s *ScheduledState) GetStatus() model.InterviewStatus   { return model.InterviewScheduled }

// InProgress State
type InProgressState struct{}

func (s *InProgressState) NextState() (InterviewState, error) { return &CompletedState{}, nil }
func (s *InProgressState) GetStatus() model.InterviewStatus   { return model.InterviewInProgress }

// Completed State
type CompletedState struct{}

func (s *CompletedState) NextState() (InterviewState, error) {
	return nil, errors.New("no next state from Completed")
}
func (s *CompletedState) GetStatus() model.InterviewStatus { return model.InterviewCompleted }

// InterviewContext manages state transitions
type InterviewContext struct {
	State InterviewState
}

func NewInterviewContext(currentStatus model.InterviewStatus) *InterviewContext {
	var state InterviewState
	switch currentStatus {
	case model.InterviewPending:
		state = &PendingState{}
	case model.InterviewScheduled:
		state = &ScheduledState{}
	case model.InterviewInProgress:
		state = &InProgressState{}
	case model.InterviewCompleted:
		state = &CompletedState{}
	default:
		state = &PendingState{} // default to pending
	}
	return &InterviewContext{State: state}
}

func (ctx *InterviewContext) Transition() error {
	nextState, err := ctx.State.NextState()
	if err != nil {
		return err
	}
	ctx.State = nextState
	return nil
}

func (ctx *InterviewContext) GetCurrentStatus() model.InterviewStatus {
	return ctx.State.GetStatus()
}
