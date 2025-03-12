package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sensorDataGeneratorWithParams(dataChan chan<- int, iterations int, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for i := 0; i < iterations; i++ {
		dataChan <- rand.Intn(100)
		<-ticker.C
	}
	close(dataChan)
}

func sensorDataGenerator(dataChan chan<- int) {
	sensorDataGeneratorWithParams(dataChan, 120, 500*time.Millisecond)
}

func dataProcessor(dataChan <-chan int, processedChan chan<- float64) {
	var batch []int
	for value := range dataChan {
		batch = append(batch, value)
		if len(batch) == 10 {
			sum := 0
			for _, v := range batch {
				sum += v
			}
			processedChan <- float64(sum) / 10.0
			batch = nil
		}
	}
	close(processedChan)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	sensorChan := make(chan int)
	processedChan := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		sensorDataGenerator(sensorChan)
		wg.Done()
	}()

	go func() {
		dataProcessor(sensorChan, processedChan)
		wg.Done()
	}()

	go func() {
		for avg := range processedChan {
			fmt.Println("Среднее значение:", avg)
		}
	}()

	wg.Wait()
}
