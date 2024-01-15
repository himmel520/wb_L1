package main

import (
	"fmt"
	"sync"
)

func squareWaitGroup() {
	// создаем WaitGroup
	var wg sync.WaitGroup
	nums := [5]int{2, 4, 6, 8, 10}

	for _, num := range nums {
		// добавляем задачу
		wg.Add(1)
		go func(num int) {
			// сообщаем, что задача завершена
			defer wg.Done()
			fmt.Printf("%v: %v\n", num, num*num)
		}(num)
	}

	// ожидаем завершение всех задач
	wg.Wait()
}
