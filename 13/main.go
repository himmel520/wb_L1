package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	a, b := 10, 3
	parallelAssignment(a, b)
	math(a, b)

}

func parallelAssignment(a int, b int) {
	a, b = b, a

	fmt.Printf("a: %v\nb: %v", a, b)
}

func math(a int, b int) {
	a += b
	b = a - b
	a -= b

	fmt.Printf("a: %v\nb: %v", a, b)
}
