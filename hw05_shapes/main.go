package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height

}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func calculateArea(s Shape) (float64, error) {
	if s == nil {
		return 0, fmt.Errorf(
			"переданный объект не является фигурой.")
	}
	return s.Area(), nil
}

func main() {
	var choice int
	fmt.Println("Выберите фигуру для вычисления площади:")
	fmt.Println("1. Круг")
	fmt.Println("2. Прямоугольник")
	fmt.Println("3. Треугольник")
	fmt.Print("Ваш выбор: ")
	fmt.Scanf("%d", &choice)

	var area float64
	var err error
	switch choice {
	case 1:
		var radius float64
		fmt.Print("Введите радиус круга: ")
		fmt.Scanf("%f", &radius)
		circle := Circle{Radius: radius}
		area, err = calculateArea(circle)
	case 2:
		var width, height float64
		fmt.Print("Введите ширину прямоугольника: ")
		fmt.Scanf("%f", &width)
		fmt.Print("Введите высоту прямоугольника: ")
		fmt.Scanf("%f", &height)
		rectangle := Rectangle{Width: width, Height: height}
		area, err = calculateArea(rectangle)
	case 3:
		var base, height float64
		fmt.Print("Введите основание треугольника: ")
		fmt.Scanf("%f", &base)
		fmt.Print("Введите высоту треугольника: ")
		fmt.Scanf("%f", &height)
		triangle := Triangle{Base: base, Height: height}
		area, err = calculateArea(triangle)
	default:
		fmt.Println("Некорректный выбор.")
		return
	}

	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Площадь:", area)
	}
}
