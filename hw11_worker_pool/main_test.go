package main

import "testing"


func TestWorkerPool(t *testing.T) {
	wp := NewWorkerPool()
	nWorkers := 5
	tasksPerWorker := 200
	expected := nWorkers * tasksPerWorker

	wp.Run(nWorkers, tasksPerWorker)

	if got := wp.Counter(); got != expected {
		t.Errorf("Ожидалось %d, получено %d.", expected, got)
	}
}
