package handlers

import (
	"planner2-копия/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBot struct {
	mock.Mock
}

func (m *MockBot) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	args := m.Called(c)
	return args.Get(0).(tgbotapi.Message), args.Error(1)
}

func TestHandleStartCommand(t *testing.T) {
	config.InitDB()
	mockBot := new(MockBot)

	mockBot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{ID: 123},
			Text: "/start",
		},
	}

	HandleStartCommand(mockBot, update)

	var user config.UserAnswer
	config.DB.Where("user_id = ?", 123).First(&user)
	assert.Equal(t, int64(123), user.UserID)
	mockBot.AssertExpectations(t)
}

func TestHandleMessage(t *testing.T) {
	config.InitDB()
	mockBot := new(MockBot)

	config.DB.Create(&config.UserAnswer{
		UserID: 456,
		Step:   10,
	})

	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{ID: 456},
			Text: "5000",
		},
	}

	mockBot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)
	HandleMessage(mockBot, update)

	var answer config.UserAnswer
	config.DB.Where("user_id = ? AND step = ?", 456, 10).First(&answer)
	assert.Equal(t, 5000.0, answer.Answer)
	mockBot.AssertExpectations(t)
}
