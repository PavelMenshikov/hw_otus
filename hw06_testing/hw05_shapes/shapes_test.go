package shapes

import (
	"testing"
)

func TestCircle_Area(t *testing.T) {
	c := Circle{Radius: 10}
	expected := 314.1592653589793
	if area := c.Area(); area != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, area)
	}
}

func TestRectangle_Area(t *testing.T) {
	r := Rectangle{Width: 5, Height: 10}
	expected := 50.0
	if area := r.Area(); area != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, area)
	}
}

func TestTriangle_Area(t *testing.T) {
	tg := Triangle{Base: 4, Height: 3}
	expected := 6.0
	if area := tg.Area(); area != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, area)
	}
}

func TestCalculateArea(t *testing.T) {
	r := Rectangle{Width: 5, Height: 10}
	area, err := CalculateArea(r)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := 50.0
	if area != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, area)
	}
}

func TestReadFloat(t *testing.T) {
	_, err := readFloat("Введите число: ")
	if err != nil {
		t.Errorf("Ошибка при чтении числа: %v", err)
	}
}

func TestReadInt(t *testing.T) {
	_, err := readInt("Введите целое число: ")
	if err != nil {
		t.Errorf("Ошибка при чтении числа: %v", err)
	}
}
