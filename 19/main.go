package main

import "fmt"

/*
Разработать программу, которая переворачивает
подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func reverseWithStr(str string) string {
	result := ""
	// итерируемся по строке
	for _, v := range str {
		// добавляем символ в начало строки
		result = string(v) + result
	}
	return result
}

func reverseWithSlice(str string) string {
	runes := []rune{}
	// итерируемся по строке
	for _, v := range str {
		// добавляем символ в начало среза
		runes = append([]rune{v}, runes...)
	}
	return string(runes)
}

func main() {
	fmt.Println(reverseWithStr("главрыба"))
	fmt.Println(reverseWithSlice("главрыба"))
}
