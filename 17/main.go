package main

import "fmt"

// 17. Реализовать бинарный поиск встроенными методами языка.

func binarySearch(slice []int, value int) (int, bool) {
	// обьявляем границы
	left, right := 0, len(slice)-1

	for left <= right {
		// получаем средний индекс
		mid := (left + right) / 2

		// выход, если элемент найден
		if slice[mid] == value {
			return mid, true
		} else if slice[mid] > value {
			// если элемент меньше, то правую границу уменьшаем
			right = mid - 1
		} else {
			// если элемент больше, то увеличиваем левую границу
			left = mid + 1
		}
	}

	// в случае, если не нашли элемент
	return -1, false
}

func main() {
	nums := []int{2, 3, 6, 7, 9, 13, 15, 16}
	ind, state := binarySearch(nums, 15)
	if !state {
		fmt.Println("not found")
	} else {
		fmt.Printf("index: %v", ind)
	}

}
