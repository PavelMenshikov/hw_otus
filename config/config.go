package config

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	BotToken = "7257094557:AAGy7dKJ3Oz0boJXPwSkxh2TwhmbbS4SiD4"
	DB       *gorm.DB
)

var PositiveQuestions = []string{
	"💰 Чтобы чувствовать себя уверенно на волнах экономических кризисов...",
	"📈 Прекрасно, что ты выбрал акции как инструмент...",
	"📊 На данном этапе облигации отличный инструмент...",
}

var NegativeQuestions = []string{
	"💼 Сколько вы тратите на нелюбимую работу?",
	"😟 Тревога о финансах – это сигнал к действию...",
	"💸 Долги – их надо отдавать. Напиши сумму своей долговой нагрузки...",
}
var Questions = map[string]string{
	// Позитивные вопросы с текстами
	"Финансовую подушку": "💰 Чтобы чувствовать себя уверенно на волнах экономических кризисов...",
	// Добавьте оставшиеся позитивные тексты...
	// Негативные вопросы с текстами
	"Тревогу о финансах": "😟 Тревога о финансах – это сигнал к действию...",
	// Добавьте оставшиеся негативные тексты...
}

func MigrateDB() {
	DB.AutoMigrate(&UserStep{}, &UserAnswer{}) // Добавьте другие структуры, если нужно
}

func FinalSummary(totalAnnual, avgMonthly float64) string {
	return fmt.Sprintf(`
Ты на шаг ближе к своим финансовым мечтам! Результаты твоего планирования на 2025 год – это не просто цифры, это живой документ, который может изменить твою жизнь. Каждая твоя инвестиция в себя, фондовый рынок и прочее – это шаг к твоей финансовой свободе.

<b>Итоговые расходы:</b> %.2f рублей в год (примерно %.2f рублей в месяц)

Понимание своих доходов и расходов – это ключ к управлению своей жизнью.

Если ты выбрал создать финансовую подушку, вложить деньги в акции и облигации, и даже рискнуть с криптовалютой, ты понимаешь, что это не просто увлечение – это твой путь к финансовой независимости! Ты знаешь, как важно чувствовать себя защищенным на волнах экономических изменений. 

Обрати внимание на свои расходы. Если они превышают твои доходы, это не повод для беспокойства – это сигнал к действию! Настало время переосмыслить свои финансовые привычки и научиться зарабатывать больше. 
Осознай: твои текущие установки могут сдерживать тебя! Что, если я скажу тебе, что ты способен на большее, чем думаешь?

Каждый шаг, который ты сделал в ходе планирования, – это возможность переосмыслить свои финансовые установки. Каждый рубль, который ты тратишь, должен работать на тебя. Ты можешь создать жизнь, о которой мечтаешь: с достатком, комфортом и свободными деньгами, просто изменяя свое восприятие денег и расходов.

Не ограничивай себя! Настало время преобразить твоё финансовое мышление и поменять свои действия, чтобы твои доходы росли, а заботы о нехватке денег исчезали. Представь, каково это – иметь возможность строить планы, не задумываясь о возможных трудностях. Осознай, что твоя решимость изменит твою жизнь.

Это не просто мечты – это твой шанс взять свою жизнь под контроль и стать хозяином своей судьбы. Начни действовать уже сейчас! Ты способен на это. Строй свое финансовое благополучие на 2025 год. Не упусти свою возможность!
`, totalAnnual, avgMonthly)
}
func UpdateUserStep(userID int64, step int, isPositive bool) error {
	return DB.Model(&UserStep{}).Where("user_id = ?", userID).Updates(UserStep{
		UserID:     userID,
		Step:       step,
		IsNegative: !isPositive,
	}).Error
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./financial_planner.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Миграции
	err = DB.AutoMigrate(&UserAnswer{}, &FinalCalculation{}, &UserStep{})
	if err != nil {
		log.Fatal("Ошибка миграции базы данных:", err)
	}
}

type UserAnswer struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     int64   `gorm:"not null"`
	Question   string  `gorm:"not null"`
	Answer     float64 `gorm:"not null"` // Оставляем тип float64
	Step       int     `gorm:"not null"`
	IsNegative bool    `gorm:"not null"`
}
type UserStep struct {
	ID         uint  `gorm:"primaryKey"`
	UserID     int64 `gorm:"not null"`
	Step       int   `gorm:"not null"`
	IsNegative bool  `gorm:"not null"`
}
type FinalCalculation struct {
	ID                     uint    `gorm:"primaryKey"`
	UserID                 int64   `gorm:"not null"`
	TotalAnnualExpenses    float64 `gorm:"not null"`
	AverageMonthlyExpenses float64 `gorm:"not null"`
}

// Создание или обновление ответа пользователя
func CreateOrUpdateUserAnswer(userID int64, question string, answer string, step int, isNegative bool) error {
	answerValue, err := strconv.ParseFloat(answer, 64)
	if err != nil {
		return fmt.Errorf("некорректный ввод: %v", err)
	}

	var userAnswer UserAnswer
	result := DB.Where("user_id = ? AND question = ?", userID, question).First(&userAnswer)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		userAnswer = UserAnswer{
			UserID:     userID,
			Question:   question,
			Answer:     answerValue,
			Step:       step,
			IsNegative: isNegative,
		}
		return DB.Create(&userAnswer).Error
	}

	userAnswer.Answer = answerValue
	userAnswer.Step = step
	userAnswer.IsNegative = isNegative
	return DB.Save(&userAnswer).Error
}

// Вычисление итогового результата
func CalculateFinalSummary(userID int64) (float64, float64) {
	// Получаем все ответы пользователя
	var answers []UserAnswer
	err := DB.Where("user_id = ?", userID).Find(&answers).Error
	if err != nil {
		log.Printf("Ошибка получения ответов пользователя: %v", err)
		return 0, 0
	}

	var totalAnnual float64
	var economicAnswer float64

	// Номер шага для вопроса "Экономить на всём не нужно"
	const EconomicQuestionStep = 30 // Указываем точный шаг для вопроса

	// Суммируем все ответы и находим ответ для экономии
	for _, answer := range answers {
		totalAnnual += answer.Answer
		if answer.Step == EconomicQuestionStep {
			economicAnswer = answer.Answer
		}
	}

	// Вычитаем ответ на вопрос об экономии, умноженный на 2
	totalAnnual -= economicAnswer * 2

	// Вычисляем среднемесячную сумму
	avgMonthly := totalAnnual / 12

	// Логируем ключевые данные
	log.Printf("Итоговая сумма перед вычитанием: %.2f", totalAnnual+economicAnswer*2)
	log.Printf("Ответ для экономии: %.2f, вычитается: %.2f", economicAnswer, economicAnswer*2)
	log.Printf("Итоговая сумма после вычитания: %.2f", totalAnnual)

	return totalAnnual, avgMonthly
}
