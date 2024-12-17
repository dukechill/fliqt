package services

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	dialector := mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true, // 跳過版本檢查，適用於測試
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	return db, mock
}

func TestListInterviews(t *testing.T) {
	db, mock := setupMockDB()
	service := NewInterviewService(db)

	mockRows := sqlmock.NewRows([]string{"id", "candidate_name", "position", "status", "scheduled_time", "notes"}).
		AddRow("1", "John Doe", "Backend Engineer", 1, time.Now(), "First round interview").
		AddRow("2", "Jane Doe", "Frontend Engineer", 2, time.Now(), "Second round interview")

	mock.ExpectQuery(`SELECT (.+) FROM "interviews" WHERE`).WillReturnRows(mockRows)

	filterParams := InterviewFilterParams{
		CandidateName: "John",
	}
	result, err := service.ListInterviews(filterParams)

	assert.NoError(t, err)
	assert.Len(t, result.Items, 2)
	assert.Equal(t, "John Doe", result.Items[0].CandidateName)
}

func TestGetInterviewByID(t *testing.T) {
	db, mock := setupMockDB()
	service := NewInterviewService(db)

	mockRow := sqlmock.NewRows([]string{"id", "candidate_name", "position", "status"}).
		AddRow("1", "John Doe", "Backend Engineer", 1)

	mock.ExpectQuery(`SELECT (.+) FROM "interviews" WHERE id = ?`).WithArgs("1").WillReturnRows(mockRow)

	result, err := service.GetInterviewByID("1")

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", result.CandidateName)
}

func TestCreateInterview(t *testing.T) {
	db, mock := setupMockDB()
	service := NewInterviewService(db)

	dto := CreateInterviewDTO{
		CandidateName: "Jane Doe",
		Position:      "Frontend Engineer",
		Status:        1,
		Notes:         "Scheduled for next week",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "interviews"`).
		WithArgs(sqlmock.AnyArg(), dto.CandidateName, dto.Position, dto.Status, dto.Notes).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := service.CreateInterview(dto)

	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", result.CandidateName)
}

func TestUpdateInterview(t *testing.T) {
	db, mock := setupMockDB()
	service := NewInterviewService(db)

	dto := UpdateInterviewDTO{
		Status:        2, // 將狀態設置為 "In Progress"
		ScheduledTime: "2024-01-01T10:00:00Z",
		Notes:         "Interview in progress",
	}

	// 模擬 SQL 更新邏輯
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "interviews"`).
		WithArgs(dto.Status, dto.ScheduledTime, dto.Notes, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := service.UpdateInterview(nil, "1", dto)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, int(result.Status)) // 確認狀態正確更新為 "In Progress"
}

func TestDeleteInterview(t *testing.T) {
	db, mock := setupMockDB()
	service := NewInterviewService(db)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "interviews" WHERE id = ?`).WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.DeleteInterview("1")

	assert.NoError(t, err)
}
