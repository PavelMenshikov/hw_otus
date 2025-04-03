package comparator

import book "github.com/PavelMenshikov/hw_otus/hw06_testing/hw04_struct_comparator/testing_book"

type PoWhat int

const (
	PoYear PoWhat = iota
	PoSize
	PoRate
)

type Comparator struct {
	fieldCompare PoWhat
}

func NewComparator(fieldCompare PoWhat) *Comparator {
	return &Comparator{fieldCompare}
}

func (c *Comparator) Compare(bookOne, bookTwo book.Book) bool {
	switch c.fieldCompare {
	case PoYear:
		return bookOne.Year() > bookTwo.Year()
	case PoSize:
		return bookOne.Size() > bookTwo.Size()
	case PoRate:
		return bookOne.Rate() > bookTwo.Rate()
	default:
		return false
	}
}
