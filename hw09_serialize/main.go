package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"

	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v2"
)

type Book struct {
	ID     int     `json:"id" xml:"id" yaml:"id"`
	Title  string  `json:"title" xml:"title" yaml:"title"`
	Author string  `json:"author" xml:"author" yaml:"author"`
	Year   int     `json:"year" xml:"year" yaml:"year"`
	Size   int     `json:"size" xml:"size" yaml:"size"`
	Rate   float64 `json:"rate" xml:"rate" yaml:"rate"`
	Sample []byte  `json:"sample" xml:"sample" yaml:"sample"`
}

func (b Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	return json.Marshal(Alias(b))
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*b = Book(aux)
	return nil
}

type BooksWrapper struct {
	XMLName xml.Name `xml:"books"`
	Books   []Book   `xml:"book"`
}

func BooksToJSON(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func BooksFromJSON(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	return books, err
}

func BooksToXML(books []Book) ([]byte, error) {
	wrapper := BooksWrapper{Books: books}
	return xml.Marshal(wrapper)
}

func BooksFromXML(data []byte) ([]Book, error) {
	var wrapper BooksWrapper
	err := xml.Unmarshal(data, &wrapper)
	return wrapper.Books, err
}

func BooksToYAML(books []Book) ([]byte, error) {
	return yaml.Marshal(books)
}

func BooksFromYAML(data []byte) ([]Book, error) {
	var books []Book
	err := yaml.Unmarshal(data, &books)
	return books, err
}

func BooksToGob(books []Book) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(books); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func BooksFromGob(data []byte) ([]Book, error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var books []Book
	if err := dec.Decode(&books); err != nil {
		return nil, err
	}
	return books, nil
}

func BooksToMsgPack(books []Book) ([]byte, error) {
	return msgpack.Marshal(books)
}

func BooksFromMsgPack(data []byte) ([]Book, error) {
	var books []Book
	err := msgpack.Unmarshal(data, &books)
	return books, err
}

func BooksToProto(books []Book) ([]byte, error) {
	var pbBooks []*bookpb.Book
	for _, b := range books {
		pbBooks = append(pbBooks, &bookpb.Book{
			Id:     int32(b.ID),
			Title:  b.Title,
			Author: b.Author,
			Year:   int32(b.Year),
			Size:   int32(b.Size),
			Rate:   b.Rate,
			Sample: b.Sample,
		})
	}

	pbWrapper := &bookpb.Books{Books: pbBooks}
	return proto.Marshal(pbWrapper)
}

func BooksFromProto(data []byte) ([]Book, error) {
	var pbWrapper bookpb.Books
	err := proto.Unmarshal(data, &pbWrapper)
	if err != nil {
		return nil, err
	}
	var books []Book
	for _, pbBook := range pbWrapper.Books {
		books = append(books, Book{
			ID:     int(pbBook.Id),
			Title:  pbBook.Title,
			Author: pbBook.Author,
			Year:   int(pbBook.Year),
			Size:   int(pbBook.Size),
			Rate:   pbBook.Rate,
			Sample: pbBook.Sample,
		})
	}
	return books, nil
}

type serializeFunc func([]Book) ([]byte, error)
type deserializeFunc func([]byte) ([]Book, error)

func demonstrateFormat(name string, to serializeFunc, from deserializeFunc, books []Book) {
	data, err := to(books)
	if err != nil {
		log.Fatalf("%s: serialization error: %v", name, err)
	}
	fmt.Printf("%s data:\n%s\n\n", name, data)
	newBooks, err := from(data)
	if err != nil {
		log.Fatalf("%s: deserialization error: %v", name, err)
	}
	fmt.Printf("Deserialized books in %s:\n%+v\n\n", name, newBooks)
}

func main() {
	gob.Register(Book{})
	books := []Book{
		{
			ID:     1,
			Title:  "Грокаем алгоритмы",
			Author: "Адитья Бхаргава",
			Year:   2016,
			Size:   256,
			Rate:   4.8,
			Sample: []byte("Пример книги."),
		},
		{
			ID:     2,
			Title:  "Совершенный код",
			Author: "Стив Макконнелл",
			Year:   2004,
			Size:   512,
			Rate:   4.5,
			Sample: []byte("Другой пример."),
		},
	}
	demonstrateFormat("JSON", BooksToJSON, BooksFromJSON, books)
	demonstrateFormat("XML", BooksToXML, BooksFromXML, books)
	demonstrateFormat("YAML", BooksToYAML, BooksFromYAML, books)
	demonstrateFormat("GOB", BooksToGob, BooksFromGob, books)
	demonstrateFormat("MessagePack", BooksToMsgPack, BooksFromMsgPack, books)
	demonstrateFormat("Protobuf", BooksToProto, BooksFromProto, books)
}
