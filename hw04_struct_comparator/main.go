package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetRate() float64 {
	return b.rate
}

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

func compareByYear(b1, b2 *Book) bool { return b1.GetYear() > b2.GetYear() }
func compareBySize(b1, b2 *Book) bool { return b1.GetSize() > b2.GetSize() }
func compareByRate(b1, b2 *Book) bool { return b1.GetRate() > b2.GetRate() }

func main() {
	books := []*Book{
		{1, "Сияние", "Стивен Кинг", 1977, 447, 4.8},
		{2, "Держи марку!", "Терри Пратчетт", 2004, 353, 4.7},
		{3, "Бойцовский клуб", "Чак Паланик", 1996, 208, 4.6},
	}

	var firstBook, secondBook, choice int
	fmt.Println("Выберите первую книгу (0, 1 или 2):")
	fmt.Scan(&firstBook)
	fmt.Println("Выберите вторую книгу (0, 1 или 2):")
	fmt.Scan(&secondBook)
	fmt.Println("Выберите параметры сравнения:")
	fmt.Println("1. По году")
	fmt.Println("2. По размеру")
	fmt.Println("3. По рейтингу")
	fmt.Print("Введите номер вашего выбора: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	}

	var comparator *Comparator
	switch choice {
	case 1:
		comparator = NewComparator(compareByYear)
	case 2:
		comparator = NewComparator(compareBySize)
	case 3:
		comparator = NewComparator(compareByRate)
	default:
		fmt.Println("Некорректный выбор")
		return
	}

	fmt.Printf(
		"Сравнение книги '%s' с книгой '%s':\n", books[firstBook].GetTitle(), books[secondBook].GetTitle())
	result := comparator.Compare(books[firstBook], books[secondBook])
	if result {
		fmt.Println(
			"Первая книга больше второй по выбранному параметру.")
	} else {
		fmt.Println("Первая книга меньше или равна второй по выбранному параметру.")
	}
}
