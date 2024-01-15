package main

import (
	"fmt"
	"strings"
)

/*
20. Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func reverseWords(str string) string {
	// получаем срез слов
	words := strings.Split(str, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		// меняем значения местами, 0-8 1-7 2-6 и тд..
		words[i], words[j] = words[j], words[i]
	}

	// обьединяем срез в строку с разделителем
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("snow dog sun y t r e w q"))
}
