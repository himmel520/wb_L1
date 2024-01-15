package main

import "fmt"

// 11. Реализовать пересечение двух неупорядоченных множеств.

// создание множества
func set(nums []int) map[int]bool {
	set := make(map[int]bool)
	// добавление уникальных значений в мапу
	for _, v := range nums {
		set[v] = true
	}

	return set
}

// нахождение пересечения множеств
func intersection(set1 map[int]bool, set2 map[int]bool) []int {
	intersection := []int{}
	// итерируемся по множеству
	for k := range set1 {
		// если ключ есть во втором множестве, то добавляем его в пересечение
		if _, ok := set2[k]; ok {
			intersection = append(intersection, k)
		}
	}

	return intersection
}

func main() {
	// создаем множества
	slice1 := set([]int{1, 2, 3, 4, 5})
	slice2 := set([]int{6, 7, 8, 4, 5})

	// выводим пересечение
	fmt.Println(intersection(slice1, slice2))

}
