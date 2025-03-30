package analyzer

import (
	"os"
	"testing"
)

func TestAnalyzeLogFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "log_test")
	if err != nil {
		t.Fatalf("Ошибка создания временного файла: %v", err)
	}

	defer os.Remove(tempFile.Name())

	logData := `INFO: Это информационное сообщение

ERROR: Это сообщение об ошибке

INFO: Еще одно информационное сообщение

DEBUG: Это отладочное сообщение

INFO: Последнее информационное сообщение`

	if _, err = tempFile.WriteString(logData); err != nil {
		t.Fatalf("Ошибка записи в файл: %v", err)
	}

	tempFile.Close()

	stats, err := AnalyzeLogFile(tempFile.Name(), "INFO")
	if err != nil {
		t.Fatalf("Ошибка анализа файла: %v", err)
	}

	expected := 3

	if stats.Count != expected {
		t.Errorf("Ожидалось %d записей, получено %d", expected, stats.Count)
	}
}
