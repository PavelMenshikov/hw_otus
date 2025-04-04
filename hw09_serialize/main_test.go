package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Тестируем JSON-сериализацию и десериализацию
func TestJSONSerialization(t *testing.T) {
	// Исходная книга
	originalBook := Book{
		ID:     1,
		Title:  "Грокаем алгоритмы",
		Author: "Адитья Бхаргава",
		Year:   2016,
		Size:   256,
		Rate:   4.8,
		Sample: []byte("Это фрагмент книги."),
	}

	// Сериализация в JSON
	jsonData, err := ToJSON(originalBook)
	if err != nil {
		t.Fatalf("Ошибка сериализации: %v", err)
	}

	// Десериализация из JSON
	newBook, err := FromJSON(jsonData)
	if err != nil {
		t.Fatalf("Ошибка десериализации: %v", err)
	}

	// Сравниваем книги (используем reflect.DeepEqual для сравнения структур)
	if !reflect.DeepEqual(originalBook, newBook) {
		t.Errorf("Ожидали %+v, а получили %+v", originalBook, newBook)
	}

	// Проверяем, что JSON можно корректно распарсить в карту
	var parsedData map[string]interface{}
	err = json.Unmarshal(jsonData, &parsedData)
	if err != nil {
		t.Fatalf("Ошибка парсинга JSON: %v", err)
	}

	// Проверяем, что в JSON есть ожидаемые поля
	expectedFields := []string{"id", "title", "author", "year", "size", "rate", "sample"}
	for _, field := range expectedFields {
		if _, ok := parsedData[field]; !ok {
			t.Errorf("Поле %s отсутствует в JSON", field)
		}
	}
}
