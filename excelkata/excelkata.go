package excelkata

import (
	"math"
	"strings"
)

func ToColNotation(num int) string {
	convert := num
	converted := make([]int, 0)

	for convert != 0 {
		var convert_mod int

		if convert%26 == 0 {
			convert_mod = 26
		} else {
			convert_mod = convert % 26
		}

		converted = append(converted, convert_mod)
		convert = convert / 26

		if convert_mod == 26 {
			convert--
		}
	}

	converted = reverseInts(converted)
	col_notation := ""

	for _, ci := range converted {
		col_notation += string(ci + 'A' - 1)
	}

	return col_notation
}

func FromColNotation(str string) int {
	convert := strings.ToUpper(str)
	convertSlice := stringToIntSlice(convert)
	convertSlice = reverseInts(convertSlice)
	converted := 0

	for i := 0; i < len(convertSlice); i++ {
		converted += int(float64(convertSlice[i]) * math.Pow(26, float64(i)))
	}

	return converted
}

func stringToIntSlice(str string) []int {
	intSlice := make([]int, 0)

	for i := 0; i < len(str); i++ {
		intSlice = append(intSlice, int(str[i]-'A'+1))
	}

	return intSlice
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
