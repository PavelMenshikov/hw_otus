package main

import "testing"

func TestBinarySearch_Found(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7
	expectedIndex := 3 // Индексы начинаются с 0
	result := binarySearch(arr, target)
	if result != expectedIndex {
		t.Errorf("Ожидалось индекс %d, получено %d", expectedIndex, result)
	}
}

func TestBinarySearch_NotFound(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 8
	result := binarySearch(arr, target)
	if result != -1 {
		t.Errorf("Ожидалось -1 для несуществующего элемента, получено %d", result)
	}
}

func TestBinarySearch_EmptySlice(t *testing.T) {
	arr := []int{}
	target := 5
	result := binarySearch(arr, target)
	if result != -1 {
		t.Errorf("Ожидалось -1 для пустого слайса, получено %d", result)
	}
}
