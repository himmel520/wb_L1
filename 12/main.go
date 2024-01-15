package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

func set(strings []string) []string {
	mappa := make(map[string]bool)
	// добавление уникальных значений в мапу
	for _, v := range strings {
		mappa[v] = true
	}

	var set []string
	// добавление уникальных ключей в срез
	for k := range mappa {
		set = append(set, k)
	}

	// возврат множества строк
	return set
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	mySet := set(strings)
	fmt.Println(mySet)
}
