package main

import "fmt"

// 21. Реализовать паттерн «адаптер» на любом примере.

// Интерфейс для персонажа
type Character interface {
	Talk()
}

// структура для en персонажа
type EnCharacter struct{}

// конструктор
func NewEnCharacter() *EnCharacter {
	return &EnCharacter{}
}

// метод EnCharacter
func (ch *EnCharacter) EnTalk() {
	fmt.Println("qwerty..")
}

// адаптер для EnCharacter
type EnAdapter struct {
	// встраиваем структуру с методами
	en *EnCharacter
}

// конструктор адаптера, реализующий интерфейс Character
func NewEnAdapter(en *EnCharacter) Character {
	return &EnAdapter{en}
}

// дополнили метод EnTalk(), не изменяя его
func (a *EnAdapter) Talk() {
	fmt.Println("en character")
	a.en.EnTalk()
}

func main() {
	// создание адаптера
	adapter := NewEnAdapter(NewEnCharacter())
	// вызов Talk
	adapter.Talk()
}
