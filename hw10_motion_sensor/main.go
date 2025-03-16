package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sensorDataGenerator(dataChan chan<- int) {
	timeout := time.After(time.Minute)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	localRand := rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404

	for {
		select {
		case <-timeout:
			close(dataChan)
			return
		case <-ticker.C:
			value := localRand.Intn(100)
			select {
			case dataChan <- value:
			default:
			}
		}
	}
}

func dataProcessor(dataChan <-chan int, processedChan chan<- float64) {
	batch := make([]int, 0, 10)
	for value := range dataChan {
		batch = append(batch, value)
		if len(batch) == cap(batch) {
			sum := 0
			for _, v := range batch {
				sum += v
			}
			processedChan <- float64(sum) / float64(cap(batch))
			batch = batch[:0]
		}
	}
	close(processedChan)
}

func main() {
	sensorChan := make(chan int, 10)
	processedChan := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sensorDataGenerator(sensorChan)
	}()

	go func() {
		defer wg.Done()
		dataProcessor(sensorChan, processedChan)
	}()

	go func() {
		for avg := range processedChan {
			fmt.Printf("Среднее значение: %.2f\n", avg)
		}
	}()

	wg.Wait()
}
