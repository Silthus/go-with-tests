package romannumerals

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	assertRomanConversion(t, 1, "I")
	assertRomanConversion(t, 2, "II")
	assertRomanConversion(t, 3, "III")
	assertRomanConversion(t, 4, "IV")
	assertRomanConversion(t, 5, "V")
	assertRomanConversion(t, 6, "VI")
	assertRomanConversion(t, 7, "VII")
	assertRomanConversion(t, 8, "VIII")
	assertRomanConversion(t, 9, "IX")
	assertRomanConversion(t, 10, "X")
	assertRomanConversion(t, 14, "XIV")
	assertRomanConversion(t, 18, "XVIII")
	assertRomanConversion(t, 20, "XX")
	assertRomanConversion(t, 39, "XXXIX")
	assertRomanConversion(t, 40, "XL")
	assertRomanConversion(t, 47, "XLVII")
	assertRomanConversion(t, 49, "XLIX")
	assertRomanConversion(t, 50, "L")
}

func assertRomanConversion(t *testing.T, number int, expected string) {
	t.Run(fmt.Sprint(number), func(t *testing.T) {
		assert.Equal(t, expected, ConvertToRoman(number))
	})
}
