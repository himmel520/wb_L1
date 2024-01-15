package main

import (
	"fmt"
	"math/big"
)

/*
22. Разработать программу, которая
перемножает, делит, складывает, вычитает
две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	a := new(big.Int).Lsh(big.NewInt(2), 25)
	b := new(big.Int).Lsh(big.NewInt(2), 20)
	fmt.Printf("a: %v b:%v\n", a, b)

	mul := new(big.Int).Mul(a, b)
	fmt.Printf("mul: %v\n", mul)

	div := new(big.Int).Div(a, b)
	fmt.Printf("div: %v\n", div)

	add := new(big.Int).Add(a, b)
	fmt.Printf("add: %v\n", add)

	sub := new(big.Int).Sub(a, b)
	fmt.Printf("sub: %v\n", sub)
}
