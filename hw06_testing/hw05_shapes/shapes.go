package shapes

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
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

func CalculateArea(s Shape) (float64, error) {
	return calculateArea(s)
}

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, fmt.Errorf("переданный объект не реализует интерфейс Shape")
	}
	return shape.Area(), nil
}

func readFloat(prompt string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения: %w", err)
	}
	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("неверный формат числа")
	}
	return value, nil
}

func readInt(prompt string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения: %w", err)
	}
	input = strings.TrimSpace(input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("неверный формат числа")
	}
	return value, nil
}

func main() {
	for {
		fmt.Println("\nВыберите фигуру для вычисления площади:")
		fmt.Println("1. Круг")
		fmt.Println("2. Прямоугольник")
		fmt.Println("3. Треугольник")
		fmt.Println("0. Выход")

		choice, err := readInt("Ваш выбор: ")
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		if choice == 0 {
			fmt.Println("Выход из программы.")
			break
		}

		var shape any
		switch choice {
		case 1:
			var radius float64
			radius, err = readFloat("Введите радиус круга: ")
			if err != nil {
				fmt.Println(err)
				continue
			}
			shape = Circle{Radius: radius}
		case 2:
			var width, height float64
			width, err = readFloat("Введите ширину прямоугольника: ")
			if err != nil {
				fmt.Println(err)
				continue
			}
			height, err = readFloat("Введите высоту прямоугольника: ")
			if err != nil {
				fmt.Println(err)
				continue
			}
			shape = Rectangle{Width: width, Height: height}
		case 3:
			var base, height float64
			base, err = readFloat("Введите основание треугольника: ")
			if err != nil {
				fmt.Println(err)
				continue
			}
			height, err = readFloat("Введите высоту треугольника: ")
			if err != nil {
				fmt.Println(err)
				continue
			}
			shape = Triangle{Base: base, Height: height}
		default:
			fmt.Println("Некорректный выбор.")
			continue
		}

		area, err := calculateArea(shape)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Printf("Площадь: %.2f\n", area)
		}
	}
}
