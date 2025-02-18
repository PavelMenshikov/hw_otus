package main

import (
	"encoding/gob"
	"reflect"
	"testing"
)

var sampleBooks = []Book{
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

func init() {
	gob.Register(Book{})
}

func testRoundTrip(t *testing.T, formatName string, toFunc func([]Book) ([]byte,
	error), fromFunc func([]byte) ([]Book, error),
) {
	t.Helper()
	data, err := toFunc(sampleBooks)
	if err != nil {
		t.Fatalf("%s: serialization error: %v", formatName, err)
	}
	newBooks, err := fromFunc(data)
	if err != nil {
		t.Fatalf("%s: deserialization error: %v", formatName, err)
	}
	if !reflect.DeepEqual(sampleBooks, newBooks) {
		t.Errorf("%s: expected %+v, got %+v", formatName, sampleBooks, newBooks)
	}
}

func TestJSONSerialization(t *testing.T) {
	testRoundTrip(t, "JSON", BooksToJSON, BooksFromJSON)
}

func TestXMLSerialization(t *testing.T) {
	testRoundTrip(t, "XML", BooksToXML, BooksFromXML)
}

func TestYAMLSerialization(t *testing.T) {
	testRoundTrip(t, "YAML", BooksToYAML, BooksFromYAML)
}

func TestGOBSerialization(t *testing.T) {
	testRoundTrip(t, "GOB", BooksToGob, BooksFromGob)
}

func TestMsgPackSerialization(t *testing.T) {
	testRoundTrip(t, "MessagePack", BooksToMsgPack, BooksFromMsgPack)
}

func TestProtoSerialization(t *testing.T) {
	testRoundTrip(t, "Protobuf", BooksToProto, BooksFromProto)
}
