package main

import (
	"fmt"
	"sync"
)

/*
18. Реализовать структуру-счетчик, которая будет
инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	// обеспечиваем безопасный доступ к count
	mu    sync.RWMutex
	count int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	// блокируем мьютекс
	c.mu.Lock()
	defer c.mu.Unlock()
	// разблокируем
	c.count++
}

func (c *Counter) Get() int {
	// блокируем на чтение
	c.mu.RLock()
	defer c.mu.RUnlock()
	// разблокируем
	return c.count
}

func countRWMutex() {
	var wg sync.WaitGroup
	counter := NewCounter()

	for i := 0; i < 1000; i++ {
		// добавляем горутину в wg
		wg.Add(1)
		go func() {
			// сообщаем, что горутина завершена
			defer wg.Done()
			counter.Increment()
		}()
	}

	// ожидаем завершение всех горутин
	wg.Wait()

	fmt.Println(counter.Get())
}
