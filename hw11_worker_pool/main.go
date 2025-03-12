package main

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	counter int
	mu      sync.Mutex
	wg      sync.WaitGroup
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{}
}

func (wp *WorkerPool) Worker(tasks int) {
	defer wp.wg.Done()
	for i := 0; i < tasks; i++ {
		wp.mu.Lock()
		wp.counter++
		wp.mu.Unlock()
	}
}

func (wp *WorkerPool) Run(nWorkers, tasksPerWorker int) {
	wp.wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go wp.Worker(tasksPerWorker)
	}
	wp.wg.Wait()
}

func (wp *WorkerPool) Counter() int {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	return wp.counter
}

func main() {
	pool := NewWorkerPool()
	pool.Run(10, 100)
	fmt.Println("Final counter:", pool.Counter())
}
