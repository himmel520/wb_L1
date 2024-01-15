package main

import (
	"context"
	"log"
	"runtime"
	"sync"
	"time"
)

// 6. Реализовать все возможные способы остановки выполнения горутины.

// завершение main
func main() {
	go func() {}()
}

// использование сигналов
func endWithSignal() {
	exit := make(chan int, 1)
	go func() {
		for {
			select {
			// читаем из exit и завершаем горутину
			case <-exit:
				log.Println("exit..")
				return
			default:
				log.Println("default")
			}
		}

	}()
	time.Sleep(1 * time.Millisecond)
	// пишем в exit
	exit <- 0
}

// использование runtime
func endWithGoexit() {
	go func() {
		log.Println("exit..")
		// завершение горутины(на остальные не влияет)
		runtime.Goexit()
	}()
	time.Sleep(1 * time.Millisecond)
}

// использование context
func endWithContext() {
	var wg sync.WaitGroup
	// создаем контекст с ограничением по времени
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// по истечении времени, завершаем горутину
			case <-ctx.Done():
				log.Println("exit..")
				return
			default:
				log.Println("default")
				time.Sleep(2 * time.Millisecond)
			}
		}
	}()

	// ожидаем выполнение горутины
	wg.Wait()
}
