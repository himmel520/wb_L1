package main

import (
	"errors"
	"fmt"
	"log"
)

/*
15. К каким негативным последствиям может
привести данный фрагмент кода, и как это исправить?

Приведите корректный пример реализации.

// использование глобальных переменных может усложнить код
var justString string

func someFunc() {
	// создается большая строка, может привести к overflows
	v := createHugeString(1 << 10)
	// В этом моменте можно поймать панику, если строка меньше, чем срез
	justString = v[:100] // меняется глобальная переменная
}
func main() {
	someFunc()
}
*/

// получаем строку указанного размера
func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() (string, error) {
	// получаем строку
	v := createHugeString(1 << 10)

	// проверка на размер
	if len(v) < 100 {
		return v, errors.New("index out of range")
	}

	// возврат среза
	return v[:100], nil
}
func main() {
	justString, err := someFunc()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(justString)
}
