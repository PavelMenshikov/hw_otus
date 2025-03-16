package main

import (
	"math/rand"
	"testing"
	"time"
)

func sensorDataGeneratorWithParams(dataChan chan<- int, duration time.Duration, interval time.Duration) {
	localRand := rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404

	timeout := time.After(duration)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			close(dataChan)
			return
		case <-ticker.C:
			select {
			case dataChan <- localRand.Intn(100):
			default:
			}
		}
	}
}

func TestSensorDataGenerator(t *testing.T) {
	t.Run("Generates data for exact duration", func(t *testing.T) {
		dataChan := make(chan int)
		testDuration := 100 * time.Millisecond
		interval := 20 * time.Millisecond

		start := time.Now()
		go sensorDataGeneratorWithParams(dataChan, testDuration, interval)

		for range dataChan {
			_ = struct{}{}
		}

		elapsed := time.Since(start)
		if elapsed < testDuration {
			t.Errorf("Генератор завершился раньше времени: %v < %v", elapsed, testDuration)
		}
	})

	t.Run("Non-blocking behavior check", func(t *testing.T) {
		dataChan := make(chan int, 1)
		testDuration := 50 * time.Millisecond
		interval := 10 * time.Millisecond

		go sensorDataGeneratorWithParams(dataChan, testDuration, interval)
		time.Sleep(testDuration + 20*time.Millisecond)

		select {
		case <-dataChan:
		case <-time.After(10 * time.Millisecond):
		}
	})
}

func TestDataProcessor(t *testing.T) {
	t.Run("Full batches processing", func(t *testing.T) {
		dataChan := make(chan int, 20)
		processedChan := make(chan float64)

		go func() {
			defer close(dataChan)
			for i := 0; i < 25; i++ {
				dataChan <- i
			}
		}()

		go dataProcessor(dataChan, processedChan)

		received := 0
		expected := 2
		for range processedChan {
			received++
			if received > expected {
				break
			}
		}

		if received != expected {
			t.Errorf("Ожидалось %d батчей, получено %d", expected, received)
		}
	})
}
