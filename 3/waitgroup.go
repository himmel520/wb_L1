package main

import (
	"fmt"
	"sync"
)

type SumSquare struct {
	mu  sync.RWMutex
	sum int
}

func NewSumSquare() *SumSquare {
	return &SumSquare{}
}

// безопасно прибавляем квадрат числа к s.sum
func (s *SumSquare) Update(num int) {
	// блокируем мьютекс для записи
	s.mu.Lock()
	// разблокируем мьютекс
	defer s.mu.Unlock()
	s.sum += num * num
}

// безопасно получаем значение s.sum
func (s *SumSquare) Get() int {
	// разрешаем множественное чтение
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sum
}

func squareSumWaitGroup() {
	// создаем WaitGroup
	var wg sync.WaitGroup
	sum := NewSumSquare()
	nums := [5]int{2, 4, 6, 8, 10}

	for _, num := range nums {
		// добавляем задачу
		wg.Add(1)
		go func(num int) {
			// сообщаем, что задача завершена
			defer wg.Done()
			sum.Update(num)
		}(num)
	}

	// Ожидаем завершение всех задач
	wg.Wait()
	fmt.Println(sum.Get())
}
