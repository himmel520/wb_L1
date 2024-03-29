package main

import "fmt"

/*
10. Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func group(temp []float64) map[int][]float64 {
	// инициализируем мапу
	group := map[int][]float64{}

	// проходим по temp
	for _, v := range temp {
		// получаем ключ(шаг 10)
		k := int(v/10) * 10
		// обновляем срез по ключу
		group[k] = append(group[k], v)
	}

	return group
}

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(group(temp))
}
