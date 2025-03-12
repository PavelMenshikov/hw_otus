package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sensorDataGenerator(dataChan chan<- int) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < 120; i++ {
		dataChan <- rand.Intn(100)
		<-ticker.C
	}
	close(dataChan)
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
