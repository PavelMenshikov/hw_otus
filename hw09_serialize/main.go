package main

import (
	"encoding/json"
	"fmt"
)

// Определяем структуру Book
type Book struct {
	ID     int     `json:"id" xml:"id" yaml:"id"`
	Title  string  `json:"title" xml:"title" yaml:"title"`
	Author string  `json:"author" xml:"author" yaml:"author"`
	Year   int     `json:"year" xml:"year" yaml:"year"`
	Size   int     `json:"size" xml:"size" yaml:"size"`
	Rate   float64 `json:"rate" xml:"rate" yaml:"rate"`
	Sample []byte  `json:"sample" xml:"sample" yaml:"sample"`
}

// Функция сериализации в JSON
func ToJSON(book Book) ([]byte, error) {
	data, err := json.Marshal(book) // ✅ Исправлено
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Функция десериализации из JSON
func FromJSON(data []byte) (Book, error) {
	var book Book
	err := json.Unmarshal(data, &book) // ✅ Исправлено
	if err != nil {
		return book, err
	}
	return book, nil
}

func main() {
	// Создаём тестовую книгу
	book := Book{
		ID:     1,
		Title:  "Грокаем алгоритмы",
		Author: "Адитья Бхаргава",
		Year:   2016,
		Size:   256,
		Rate:   4.8,
		Sample: []byte("Это фрагмент книги."),
	}

	// Тестируем сериализацию
	jsonData, err := ToJSON(book) // ✅ Исправлено
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("JSON:", string(jsonData))

	// Тестируем десериализацию
	newBook, err := FromJSON(jsonData) // ✅ Исправлено
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}
	fmt.Println("Десериализованная книга:", newBook)
}
