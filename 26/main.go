package main

import (
	"fmt"
	"strings"
)

/*
26. Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true abCdefAaf — false aabcd — false
*/

func checkUnique(str string) bool {
	// создаем мапу
	symbol := map[rune]int{}

	// итерируемся по строке
	for _, v := range strings.ToLower(str) {
		// проверяем наличие ключа в мапе
		if _, ok := symbol[v]; ok {
			return false
		}
		// добавляем ключ
		symbol[v] = 0
	}

	return true

}

func main() {
	fmt.Println(checkUnique("abcd"))
	fmt.Println(checkUnique("abCdefAaf"))
	fmt.Println(checkUnique("aabcd"))
}
