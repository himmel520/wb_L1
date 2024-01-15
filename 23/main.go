package main

import "fmt"

// 23. Удалить i-ый элемент из слайса.

// Используем дженерик, чтобы принимать срез любого типа
func remove[T any](slice []T, i int) []T {
	// возвращаем измененный срез типа T
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	fmt.Println(remove([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(remove([]string{"q", "w", "e", "r", "t", "y"}, 5))
}
