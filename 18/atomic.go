package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func countAtomic() {
	var (
		wg      sync.WaitGroup
		counter int64
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Используем атомарную операцию для инкрементации счетчика
			atomic.AddInt64(&counter, 1)
		}()
	}

	// ожидаем завершение всех горутин
	wg.Wait()

	fmt.Println(counter)
}
