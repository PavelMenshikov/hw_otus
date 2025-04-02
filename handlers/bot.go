package handlers

import (
	"fmt"
	"log"
	"planner2-копия/config"
	"planner2-копия/texts"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleStartCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	userID := message.From.ID
	userName := message.From.FirstName

	config.DB.Delete(&config.UserAnswer{}, "user_id = ?", userID)

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
		"Привет, <b>%s</b>! Мы начнём с позитивных вопросов, затем перейдём к негативным. Отвечайте последовательно.",
		userName,
	))
	msg.ParseMode = "HTML"
	bot.Send(msg)

	askNextQuestion(bot, message.Chat.ID, userID, false, 0)
}

func askNextQuestion(bot *tgbotapi.BotAPI, chatID int64, userID int64, isNegative bool, step int) {
	var questions map[string]string
	if isNegative {
		questions = texts.NegativeButtons
	} else {
		questions = texts.PositiveButtons
	}

	keys := make([]string, 0, len(questions))
	for k := range questions {
		keys = append(keys, k)
	}

	if step >= len(keys) {
		if isNegative {
			handleCombinedSummary(bot, chatID, userID)
		} else {
			msg := tgbotapi.NewMessage(chatID, "Теперь давайте поговорим о том, от чего вы хотите избавиться.")
			bot.Send(msg)
			askNextQuestion(bot, chatID, userID, true, 0)
		}
		return
	}

	currentQuestion := keys[step]
	currentText := questions[currentQuestion]

	msg := tgbotapi.NewMessage(chatID, currentText)
	bot.Send(msg)

	err := config.CreateOrUpdateUserAnswer(userID, currentQuestion, 0, step, isNegative)
	if err != nil {
		log.Printf("Ошибка сохранения шага: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при сохранении вопроса. Попробуйте снова.")
		bot.Send(msg)
	}
}

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	userID := message.From.ID

	if message.IsCommand() {
		switch message.Command() {
		case "start":
			HandleStartCommand(bot, message)
		default:
			msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Попробуйте ещё раз.")
			bot.Send(msg)
		}
		return
	}

	amount, err := strconv.ParseFloat(message.Text, 64)
	if err != nil || amount < 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, введите корректное число.")
		bot.Send(msg)
		return
	}

	userAnswer, err := config.GetCurrentStep(userID)
	if err != nil {
		log.Printf("Ошибка получения текущего шага: %v", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка. Попробуйте снова.")
		bot.Send(msg)
		return
	}

	userAnswer.Answer = amount
	if err := config.DB.Save(&userAnswer).Error; err != nil {
		log.Printf("Ошибка сохранения ответа: %v", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка при сохранении ответа. Попробуйте снова.")
		bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Ваш ответ сохранён!")
	bot.Send(msg)

	askNextQuestion(bot, message.Chat.ID, userID, userAnswer.IsNegative, userAnswer.Step+1)
}

func handleCombinedSummary(bot *tgbotapi.BotAPI, chatID int64, userID int64) {
	var answers []config.UserAnswer
	config.DB.Where("user_id = ?", userID).Find(&answers)

	positiveSummary := generateFinalSummary(answers, false)
	negativeSummary := generateFinalSummary(answers, true)
	totalAnnual, totalMonthly := calculateTotalExpenses(answers)

	finalMessage := fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		positiveSummary,
		negativeSummary,
		texts.FinalSummary(totalAnnual, totalMonthly),
	)

	msg := tgbotapi.NewMessage(chatID, finalMessage)
	msg.ParseMode = "HTML"
	bot.Send(msg)
}

func generateFinalSummary(answers []config.UserAnswer, isNegative bool) string {
	var details []string
	var total float64

	for _, answer := range answers {
		if isNegative == texts.IsNegative(answer.Question) {
			details = append(details, fmt.Sprintf("- %s: %.2f рублей", answer.Question, answer.Answer))
			total += answer.Answer
		}
	}

	if isNegative {
		return fmt.Sprintf(
			"Ты решил оставить в уходящем году следующие расходы:\n\n%s\n<b>Итоговая сумма:</b> %.2f рублей",
			strings.Join(details, "\n"), total,
		)
	}
	return fmt.Sprintf(
		"Ты на шаг ближе к своим финансовым мечтам!\n\n<b>Позитивные расходы:</b>\n%s\n<b>Итоговая сумма:</b> %.2f рублей",
		strings.Join(details, "\n"), total,
	)
}

func calculateTotalExpenses(answers []config.UserAnswer) (float64, float64) {
	var total float64
	for _, answer := range answers {
		total += answer.Answer
	}
	return total, total / 12
}
