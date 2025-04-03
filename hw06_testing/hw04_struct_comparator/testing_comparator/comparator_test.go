package comparator_test

import (
	"testing"

	book "github.com/PavelMenshikov/hw_otus/hw06_testing/hw04_struct_comparator/testing_book"
	comparator "github.com/PavelMenshikov/hw_otus/hw06_testing/hw04_struct_comparator/testing_comparator"
	"github.com/stretchr/testify/assert"
)

func TestComparator(t *testing.T) {
	bookOne := book.NewBook(1, "Book One", "Author One", 2020, 300, 4.0)
	bookTwo := book.NewBook(2, "Book Two", "Author Two", 2021, 350, 4.5)

	yearComparator := comparator.NewComparator(comparator.PoYear)
	assert.True(t, yearComparator.Compare(*bookTwo, *bookOne), "BookTwo должна быть больше BookOne по году")

	sizeComparator := comparator.NewComparator(comparator.PoSize)
	assert.True(t, sizeComparator.Compare(*bookTwo, *bookOne), "BookTwo должна быть больше BookOne по размеру")

	rateComparator := comparator.NewComparator(comparator.PoRate)
	assert.True(t, rateComparator.Compare(*bookTwo, *bookOne), "BookTwo должна быть больше BookOne по рейтингу")

	invalidComparator := comparator.NewComparator(100)
	assert.False(t, invalidComparator.Compare(*bookOne, *bookTwo),
		"Сравнение с некорректным значением должно вернуть false")
}
