package main

import (
	"fmt"
	"reflect"
)

/*
14. Разработать программу, которая в рантайме способна
определить тип переменной: int, string, bool, channel
из переменной типа interface{}.
*/

// используем fmt
func printTypeFmt(value interface{}) {
	fmt.Printf("%T\n", value)
}

// используем reflect
func printTypeReflect(value interface{}) {
	fmt.Println(reflect.TypeOf(value))
}

// используем type assertion с switch-case
func printTypeAssertion(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	}
}

func main() {
	var num int
	printTypeFmt(num)

	var str string
	printTypeReflect(str)

	var mybool bool
	printTypeAssertion(mybool)

	var c chan int
	printTypeAssertion(c)
}
