package main

import 
   "fmt"

	 type Book struct {id int title string author string year int size int rate float64}

	 /*/func (b *Book) SetId (id int) {b.id = id}
	 func (b *Book) SetTitle(title string) {b.title = title}
	 func (b *Book) SetAuthor(author string) {b.author = author}
	 func (b *Book) SetYear(year int) {b.year = year}
	 func (b *Book) SetSize(size int) {b.size = size}
	 func (b *Book) SetRate(rate float64) {b.rate = rate}
	 /*/
	 
	 func (b *Book) GetId() int{ return b.id}
	 func (b *Book) GetTitle() string{ return b.title}
	 func (b *Book) GetAuthor() string{ return b.author}
	 func (b *Book) GetYear() int{ return b.year}
	 func (b *Book) GetSize() int{ return b.size}
	 func (b *Book) GetRate() float64{ return b.rate}

	 func (b *Book) SetAllFields(id int, title, author string, year, size int, rate float64) { b.id = id b.title = title b.author = author b.year = year b.size = size b.rate = rate }
   
	 book1 := &Book{} book1.SetAllFields(1, "Сияние", "Стивен Кинг", 1977, 456, 4.5) 
	 book2 := &Book{} book1.SetAllFields(2, "Держи Марку!", "Терри Пратчет", 2001, 480, 4.2) 
	 book3 := &Book{} book1.SetAllFields(3, "Бойцовский Клуб", "Чак Паланик", 2002, 430, 5.0) 

	books := []*Book(book1, book2, book3)
	or _, book := range books { fmt.Printf("ID книги: %d\n", book.GetID()) 
	fmt.Printf("Название книги: %s\n", book.GetTitle()) 
	fmt.Printf("Автор книги: %s\n", book.GetAuthor()) 
	fmt.Printf("Год издания книги: %d\n", book.GetYear()) 
	fmt.Printf("Размер книги: %d\n", book.GetSize()) 
	fmt.Printf("Рейтинг книги: %.1f\n", book.GetRate()) fmt.Println() } }

func main() {
	books := []*Book(book1, book2, book3)
	or _, book := range books { fmt.Printf("ID книги: %d\n", book.GetID()) 
	fmt.Printf("Название книги: %s\n", book.GetTitle()) 
	fmt.Printf("Автор книги: %s\n", book.GetAuthor()) 
	fmt.Printf("Год издания книги: %d\n", book.GetYear()) 
	fmt.Printf("Размер книги: %d\n", book.GetSize()) 
	fmt.Printf("Рейтинг книги: %.1f\n", book.GetRate()) fmt.Println() } }

