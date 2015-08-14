package excelkata

import ()

func ToColNotation(num int) string {
	convert := num
	converted := make([]int, 0)

	for convert != 0 {
		converted = append(converted, convert%26)
		convert = convert / 26
	}

	converted = reverseInts(converted)
	col_notation := ""

	for _, ci := range converted {
		col_notation += string(ci + 'A' - 1)
	}

	return col_notation
}

func FromColNotation(str string) int {
	return 0
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
