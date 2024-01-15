package main

import (
	"fmt"
	"math/rand"
)

func quickSort(slice []int) []int {
	// базовый случай
	if len(slice) <= 1 {
		return slice
	}

	// выбираем случайный опорный элемент
	pivot := rand.Intn(len(slice))

	var less, equal, greater []int
	// разделяем элементы по срезам
	for _, v := range slice {
		switch {
		case v < slice[pivot]:
			// append может быть неэффективным при больших данных
			less = append(less, v)
		case v == slice[pivot]:
			equal = append(equal, v)
		case v > slice[pivot]:
			greater = append(greater, v)
		}
	}

	// Рекурсивно применяем quickSort к less/greather и объединяем с equel.
	return append(append(quickSort(less), equal...), quickSort(greater)...)
}

func main() {
	slice := []int{2, 6, 1, 3, 0, 5, 2}
	fmt.Println(quickSort(slice))

	// example
	// slice 		   pivot less  equal greater
	// [2 6 1 3 0 5 2] 2     [1 0] [2 2] [6 3 5]
	// [1 0]           0     []    [0]   [1]
	// [6 3 5]         5     [3]   [5]   [6]

	// [0 1 2 2 3 5 6]
}
