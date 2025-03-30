package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LogStats struct {
	Count int
}

func AnalyzeLogFile(filePath string, level string) (*LogStats, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close()

	stats := &LogStats{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, level) {
			stats.Count++
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return stats, nil
}
