package main

import (
	"testing"
	"time"
)

func TestSensorDataGenerator(t *testing.T) {
	dataChan := make(chan int)

	go sensorDataGeneratorWithParams(dataChan, 4, 10*time.Millisecond)

	count := 0
	for range dataChan {
		count++
	}

	if count != 4 {
		t.Errorf("Ожидалось 4 значения, получено %d", count)
	}
}

func TestDataProcessor(t *testing.T) {
	dataChan := make(chan int, 20)
	processedChan := make(chan float64)

	go dataProcessor(dataChan, processedChan)

	batches := [][]int{
		{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
		{5, 15, 25, 35, 45, 55, 65, 75, 85, 95},
	}

	for _, batch := range batches {
		for _, v := range batch {
			dataChan <- v
		}
	}
	close(dataChan)

	for _, batch := range batches {
		sum := 0
		for _, v := range batch {
			sum += v
		}
		expectedAvg := float64(sum) / 10.0

		actualAvg, open := <-processedChan
		if !open {
			t.Errorf("Канал processedChan закрылся раньше времени")
			return
		}
		if actualAvg != expectedAvg {
			t.Errorf("Ожидалось %.2f, получено %.2f", expectedAvg, actualAvg)
		}
	}

	_, open := <-processedChan
	if open {
		t.Errorf("Канал processedChan должен быть закрыт")
	}
}
