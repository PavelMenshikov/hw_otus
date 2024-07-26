package main

import (
	"fmt"

	"github.com/PavelMenshikov/hw_otus/hw02_fix_app/printer"
	"github.com/PavelMenshikov/hw_otus/hw02_fix_app/reader"
)

func main() {
	var path string
	fmt.Printf("Введите путь к файлу JSON:")
	fmt.Scanln(&path)

	if path == "" {
		path = "data.json"
	}

	staff, err := reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Printf("Ошибка при чтении данных: %v\n", err)
		return
	}
	fmt.Println(err)

	printer.PrintStaff(staff)
	fmt.Println("Нажмите Enter для выхода...")
	fmt.Scanln()
}
