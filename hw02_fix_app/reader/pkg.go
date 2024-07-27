package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/PavelMenshikov/hw_otus/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Путь к файлу указан неверно: %v\n", err)
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		return nil, err
	}

	// var data []types.Employee

	// err = json.Unmarshal(data, &employeeData)

	// res := data

	// return res, nil
	var employees []types.Employee
	err = json.Unmarshal(data, &employees)
	if err != nil {
		fmt.Printf("Ошибка парсинга JSON: %v\n", err)
		return nil, err
	}
	return employees, nil
}
