package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func main() {
	var path string = "data.json"
	var staff []types.Employee
	staff, _ = reader.ReadJSON(path, -1)

	for {
		fmt.Println("Введите команду: ")
		fmt.Println("1 - Вывести список сотрудников")
		fmt.Println("2 - Выполнить другую команду")
		fmt.Println("0 - Выход")

		var command int
		fmt.Scanln(&command)

		switch command {
		case 1:
			for _, employee := range staff {
				printer.PrintStaff(employee)
			}
		case 2:

		case 0:
			fmt.Println("До свидания!")
			return // Выход из программы
		default:
			fmt.Println("Неверная команда.")
		}
	}
}
