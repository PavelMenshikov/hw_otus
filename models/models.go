package models

type UserAnswer struct {
	ID         int
	UserID     int
	QuestionID int
	Answer     string
}

type FinalCalculation struct {
	ID                     int
	UserID                 int
	TotalAnnualExpenses    float64
	AverageMonthlyExpenses float64
}
