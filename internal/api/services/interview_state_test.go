package services

import (
	"testing"

	"fliqt/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestInterviewStateTransitions(t *testing.T) {
	ctx := NewInterviewContext(model.InterviewPending)
	assert.Equal(t, model.InterviewPending, ctx.GetCurrentStatus())

	err := ctx.Transition()
	assert.NoError(t, err)
	assert.Equal(t, model.InterviewScheduled, ctx.GetCurrentStatus())

	err = ctx.Transition()
	assert.NoError(t, err)
	assert.Equal(t, model.InterviewInProgress, ctx.GetCurrentStatus())

	err = ctx.Transition()
	assert.NoError(t, err)
	assert.Equal(t, model.InterviewCompleted, ctx.GetCurrentStatus())

	err = ctx.Transition()
	assert.Error(t, err)
	assert.Equal(t, "no next state from Completed", err.Error())
}
