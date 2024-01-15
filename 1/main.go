package main

import (
	"log"

	"github.com/himmel520/wb_L1/1/entity"
)

/*
1. Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от
родительской структуры Human (аналог наследования).
*/

func main() {
	human := entity.NewHuman("vlad", 19)
	// получаем action с методами Human
	action := entity.NewAction(human)

	action.AddLanguage("en")
	action.AddLanguage("ru")

	err := action.RemoveLanguage("en")
	if err != nil {
		log.Println(err)
	}

	action.GetInfo()
}
