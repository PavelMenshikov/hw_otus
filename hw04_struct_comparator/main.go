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

type Comparator struct {
	comparisonField ComparisonField
}

func NewComparator(comparisonField ComparisonField) *Comparator {
	return &Comparator{comparisonField: comparisonField}
}

func (c *Comparator) Compare(book1, book2 *Book) bool {
	switch c.comparisonField {
	case ByYear:
		return book1.GetYear() > book2.GetYear()
	case BySize:
		return book1.GetSize() > book2.GetSize()
	case ByRate:
		return book1.GetRate() > book2.GetRate()
	default:
		return false
	}
}

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
		comparator = NewComparator(ByYear)
	case 2:
		comparator = NewComparator(BySize)
	case 3:
		comparator = NewComparator(ByRate)
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
