package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateFinalSummary(t *testing.T) {
	InitDB()
	userID := int64(123)

	DB.Create(&UserAnswer{
		UserID: userID,
		Answer: 1000,
		Step:   30,
	})
	DB.Create(&UserAnswer{
		UserID: userID,
		Answer: 500,
		Step:   10,
	})

	total, monthly := CalculateFinalSummary(userID)
	assert.Equal(t, 1000.0, total)
	assert.Equal(t, -500.0/12, monthly)
}

func TestCreateOrUpdateUserAnswer(t *testing.T) {
	userID := int64(456)
	err := CreateOrUpdateUserAnswer(userID, "Test question", "150", 1, false)
	assert.NoError(t, err)

	var answer UserAnswer
	DB.Where("user_id = ? AND question = ?", userID, "Test question").First(&answer)
	assert.Equal(t, 150.0, answer.Answer)
}
func TestMain(m *testing.M) {
	config.InitConfig()
	config.InitDB()
	exitVal := m.Run()
	os.Exit(exitVal)
}
