package main

import (
	"errors"
	"fmt"
	"log"
)

/*
8. Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(num int64, pos uint, flag bool) (int64, error) {
	// валидируем позицию
	if pos > 63 {
		return 0, errors.New("index must be less than 64")
	}

	// создаем маску с 1 на указанной позиции
	var mask int64 = 1 << pos

	if flag {
		// ставим бит в 1 помощью побитового или
		return num | mask, nil
	}

	// ставим бит в 0 с помощью побитового и-нет
	return num &^ mask, nil

}

func main() {
	num, err := setBit(10, 1, false)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%064b", num)

	// число     0000000000000000000000000000000000000000000000000000000000001010
	// маска     0000000000000000000000000000000000000000000000000000000000000010
	// результат 0000000000000000000000000000000000000000000000000000000000001000
}
