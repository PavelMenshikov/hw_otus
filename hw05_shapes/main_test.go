package main

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	const epsilon = 1e-6
	return math.Abs(a-b) < epsilon
}

func TestCalculateArea_Circle(t *testing.T) {
	c := Circle{Radius: 5}
	expected := math.Pi * 25
	area, err := calculateArea(c)
	if err != nil {
		t.Errorf("Непредвиденная ошибка: %v", err)
	}
	if !almostEqual(area, expected) {
		t.Errorf("Ожидалось: %.6f, получили: %.6f", expected, area)
	}
}

func TestCalculateArea_Rectangle(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	expected := 50.0 // 10 * 5
	area, err := calculateArea(r)
	if err != nil {
		t.Errorf("Непредвиденная ошибка: %v", err)
	}
	if !almostEqual(area, expected) {
		t.Errorf("Ожидалось: %.6f, получили: %.6f", expected, area)
	}
}

func TestCalculateArea_Triangle(t *testing.T) {
	tr := Triangle{Base: 8, Height: 6}
	expected := 0.5 * 8 * 6 // 24
	area, err := calculateArea(tr)
	if err != nil {
		t.Errorf("Непредвиденная ошибка: %v", err)
	}
	if !almostEqual(area, expected) {
		t.Errorf("Ожидалось: %.6f, получили: %.6f", expected, area)
	}
}

func TestCalculateArea_InvalidType(t *testing.T) {

	var notAShape = "это не фигура"
	_, err := calculateArea(notAShape)
	if err == nil {
		t.Errorf("ошибки не получено")
	}
}
