package handlers

import (
	"log"
	"planner2/config"
)

func CalculateTotalExpenses(userID int64) (float64, float64) {
	var totalAnnualExpenses float64

	rows, err := config.DB.Raw("SELECT answer FROM user_answers WHERE user_id = ?", userID).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var answer float64
		if err := rows.Scan(&answer); err != nil {
			log.Fatal(err)
		}
		totalAnnualExpenses += answer
	}

	averageMonthlyExpenses := totalAnnualExpenses / 12
	return totalAnnualExpenses, averageMonthlyExpenses
}
