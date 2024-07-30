package main

import "fmt"

// Структура, представляющая книгу
type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// Тип для указания параметра сравнения
type ComparisonField int

const (
	ByYear ComparisonField = iota
	BySize
	ByRate
)

type compareFunc func(*Book, *Book) bool

type Comparator struct {
	compareFunc compareFunc
}

func NewComparator(compareFunc compareFunc) *Comparator {
	return &Comparator{compareFunc: compareFunc}
}

func (c *Comparator) Compare(book1, book2 *Book) bool {
	return c.compareFunc(book1, book2)
}

func main() {
	// Создаем несколько книг с использованием композитных литералов
	books := []*Book{
		{1, "Сияние", "Стивен Кинг", 1977, 447, 4.8},
		{2, "Держи марку!", "Терри Пратчетт", 2004, 353, 4.7},
		{3, "Бойцовский клуб", "Чак Паланик", 1996, 208, 4.6},
	}

	// Выводим информацию о книгах
	for _, book := range books {
		fmt.Printf("ID книги: %d\n", book.id)
		fmt.Printf("Название книги: %s\n", book.title)
		fmt.Printf("Автор книги: %s\n", book.author)
		fmt.Printf("Год издания книги: %d\n", book.year)
		fmt.Printf("Размер книги: %d\n", book.size)
		fmt.Printf("Рейтинг книги: %.1f\n", book.rate)
		fmt.Println()
	}

	// Сравнение с использованием гибкого компаратора
	compareBySize := func(b1, b2 *Book) bool { return b1.size > b2.size }
	sizeComparator := NewComparator(compareBySize)
	fmt.Println("Сравнение по размеру (book1 vs book3):", sizeComparator.Compare(books[0], books[2]))

	compareByYear := func(b1, b2 *Book) bool { return b1.year > b2.year }
	yearComparator := NewComparator(compareByYear)
	fmt.Println("Сравнение по году (book1 vs book2):", yearComparator.Compare(books[0], books[1]))

	compareByRate := func(b1, b2 *Book) bool { return b1.rate > b2.rate }
	rateComparator := NewComparator(compareByRate)
	fmt.Println("Сравнение по рейтингу (book2 vs book3):", rateComparator.Compare(books[1], books[2]))

}
