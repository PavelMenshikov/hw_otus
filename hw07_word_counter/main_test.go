package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	text := "Hello, world! Hello, Go. Go is great, and Go is fun."
	expected := map[string]int{
		"hello": 2,
		"world": 1,
		"go":    3,
		"is":    2,
		"great": 1,
		"and":   1,
		"fun":   1,
	}
	result := countWords(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestCountWords_EmptyString(t *testing.T) {
	text := ""
	expected := map[string]int{}
	result := countWords(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty map, got: %v", result)
	}
}

func TestCountWords_Punctuation(t *testing.T) {
	text := "Go, Go, Go! #awesome."
	expected := map[string]int{
		"go":      3,
		"awesome": 1,
	}
	result := countWords(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
