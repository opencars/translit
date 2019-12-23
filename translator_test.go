package translit

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneToUA(t *testing.T) {
	assert.Equal(t, 'А', RuneToUA('A'))
	assert.Equal(t, 'В', RuneToUA('B'))
	assert.Equal(t, 'С', RuneToUA('C'))
	assert.Equal(t, 'Z', RuneToUA('Z'))
	assert.Equal(t, '1', RuneToUA('1'))
}

func TestRuneToLatin(t *testing.T) {
	assert.Equal(t, 'A', RuneToLatin('А'))
	assert.Equal(t, 'B', RuneToLatin('В'))
	assert.Equal(t, 'C', RuneToLatin('С'))
	assert.Equal(t, 'Z', RuneToLatin('Z'))
	assert.Equal(t, '1', RuneToLatin('1'))
}

func TestToUA(t *testing.T) {
	t.Run("simple strings", func(t *testing.T) {
		assert.Equal(t, "АВС", ToUA("ABC"))
		assert.Equal(t, "АА0000АА", ToUA("AA0000AA"))
	})

	t.Run("without latin", func(t *testing.T) {
		assert.Equal(t, "АБВ", "АБВ")
		assert.Equal(t, "123456789", "123456789")
		assert.Equal(t, "АХ1234ВА", "АХ1234ВА")
	})

	t.Run("latin to cyrillic for each region", func(t *testing.T) {
		fixtures := []string{
			"AK", "AB", "AC", "AE", "AH", "AM", "AO", "AP", "AT",
			"AA", "AI", "BA", "BB", "BC", "BE", "BH", "BI", "BK",
			"CH", "BM", "BO", "AX", "BT", "BX", "CA", "CB", "CE",
		}

		expected := []string{
			"АК", "АВ", "АС", "АЕ", "АН", "АМ", "АО", "АР", "АТ",
			"АА", "АІ", "ВА", "ВВ", "ВС", "ВЕ", "ВН", "ВІ", "ВК",
			"СН", "ВМ", "ВО", "АХ", "ВТ", "ВХ", "СА", "СВ", "СЕ",
		}

		for i := range fixtures {
			assert.Equal(t, expected[i], ToUA(fixtures[i]))
		}
	})
}

func TestToLatin(t *testing.T) {
	t.Run("simple strings", func(t *testing.T) {
		assert.Equal(t, "ABC", ToLatin("АВС"))
		assert.Equal(t, "AA0000AA", ToLatin("АА0000АА"))
	})

	t.Run("without cyrillic", func(t *testing.T) {
		assert.Equal(t, "ABC", ToLatin("ABC"))
		assert.Equal(t, "123456789", ToLatin("123456789"))
		assert.Equal(t, "AX1234BA", ToLatin("AX1234BA"))
	})

	t.Run("cyrillic to latin for each region", func(t *testing.T) {
		actual := []string{
			"АК", "АВ", "АС", "АЕ", "АН", "АМ", "АО", "АР", "АТ",
			"АА", "АІ", "ВА", "ВВ", "ВС", "ВЕ", "ВН", "ВІ", "ВК",
			"СН", "ВМ", "ВО", "АХ", "ВТ", "ВХ", "СА", "СВ", "СЕ",
		}

		expected := []string{
			"AK", "AB", "AC", "AE", "AH", "AM", "AO", "AP", "AT",
			"AA", "AI", "BA", "BB", "BC", "BE", "BH", "BI", "BK",
			"CH", "BM", "BO", "AX", "BT", "BX", "CA", "CB", "CE",
		}

		for i := range actual {
			assert.Equal(t, expected[i], ToLatin(actual[i]))
		}
	})
}

func BenchmarkToUA(b *testing.B) {
	fixtures := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		fixtures[i] = strconv.Itoa(1000 + rand.Int()%8999)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ToUA("AA" + fixtures[i] + "BP")
	}
}

func BenchmarkRuneToUA(b *testing.B) {
	fixtures := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		fixtures[i] = rune('A' + rand.Int()%'Z')
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RuneToUA(fixtures[i])
	}
}

func BenchmarkToLatin(b *testing.B) {
	fixtures := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		fixtures[i] = strconv.Itoa(1000 + rand.Int()%8999)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ToLatin("АА" + fixtures[i] + "ВР")
	}
}

func BenchmarkRuneToLatin(b *testing.B) {
	fixtures := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		fixtures[i] = rune('А' + rand.Int()%'Я')
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RuneToLatin(fixtures[i])
	}
}
