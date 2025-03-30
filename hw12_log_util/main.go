package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/PavelMenshikov/hw_otus/hw12_log_util/analyzer"
)

func main() {
	logFilePath := flag.String("file", "", "Путь к лог-файлу (обязательный)")

	logLevel := flag.String("level", "info", "Уровень логов (необязательно)")

	outputFilePath := flag.String("output", "", "Путь к файлу для записи статистики (необязательно)")

	flag.Parse()

	fmt.Println("Открываю файл:", *logFilePath)

	if *logFilePath == "" {
		*logFilePath = os.Getenv("LOG_ANALYZER_FILE")
	}

	if *logLevel == "" {
		*logLevel = os.Getenv("LOG_ANALYZER_LEVEL")
	}

	if *outputFilePath == "" {
		*outputFilePath = os.Getenv("LOG_ANALYZER_OUTPUT")
	}

	if *logFilePath == "" {
		log.Fatal("Ошибка: не указан путь к лог-файлу (флаг -file или переменная окружения LOG_ANALYZER_FILE)")
	}

	stats, err := analyzer.AnalyzeLogFile(*logFilePath, *logLevel)
	if err != nil {
		log.Fatalf("Ошибка анализа файла: %v", err)
	}

	result := fmt.Sprintf("Статистика для уровня '%s':\nВсего записей: %d\n", *logLevel, stats.Count)

	if *outputFilePath != "" {
		err = os.WriteFile(*outputFilePath, []byte(result), 0o600)
		if err != nil {
			log.Fatalf("Ошибка записи в файл: %v", err)
		}

		fmt.Println("Статистика записана в", *outputFilePath)
	} else {
		fmt.Println(result)
	}
}
