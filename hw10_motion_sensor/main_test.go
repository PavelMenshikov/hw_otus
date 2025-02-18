package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// Тест генерации данных
func TestGenerateSensorData(t *testing.T) {
	dataCh := make(chan float64, 10)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		generateSensorData(dataCh)
	}()

	time.Sleep(time.Second * 2) // Ждём некоторое время, чтобы собрать данные
	close(dataCh)
	wg.Wait()

	if len(dataCh) == 0 {
		t.Errorf("Данные с сенсора не были сгенерированы")
	}
}

// Тест вычисления среднего
func TestProcessSensorData(t *testing.T) {
	dataCh := make(chan float64, 10)
	resultCh := make(chan float64, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		processSensorData(dataCh, resultCh)
	}()

	for i := 0; i < 10; i++ {
		dataCh <- float64(rand.Intn(100))
	}
	close(dataCh)
	wg.Wait()

	select {
	case avg := <-resultCh:
		if avg <= 0 {
			t.Errorf("Среднее арифметическое должно быть больше 0, получено: %f", avg)
		}
	case <-time.After(time.Second):
		t.Error("Не получено среднее значение из канала")
	}
}

// Тест работы всей системы
func TestFullPipeline(t *testing.T) {
	dataCh := make(chan float64, 10)
	resultCh := make(chan float64, 10)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		generateSensorData(dataCh)
	}()

	go func() {
		defer wg.Done()
		processSensorData(dataCh, resultCh)
	}()

	time.Sleep(time.Second * 2) // Ждём некоторое время
	close(dataCh)
	wg.Wait()

	select {
	case avg := <-resultCh:
		if avg <= 0 {
			t.Errorf("Среднее арифметическое должно быть больше 0, получено: %f", avg)
		}
	case <-time.After(time.Second):
		t.Error("Не получено среднее значение из канала")
	}
}
