package entity

import (
	"fmt"
)

type Human struct {
	name      string
	age       uint8
	languages []string
}

func NewHuman(name string, age uint8) *Human {
	return &Human{
		name: name,
		age:  age,
	}
}

// добавление элемента в срез languages
func (h *Human) AddLanguage(language string) {
	h.languages = append(h.languages, language)
}

// удаление первого найденного элемента из среза languages
func (h *Human) RemoveLanguage(language string) error {
	for i, v := range h.languages {
		if v == language {
			h.languages = append(h.languages[:i], h.languages[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("not found value: %v", language)
}

func (h *Human) GetInfo() {
	fmt.Printf("name: %v\nage: %v\nlanguages: %q", h.name, h.age, h.languages)
}
