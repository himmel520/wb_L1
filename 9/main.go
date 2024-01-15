package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

/*
9. Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func worker(ctx context.Context, x <-chan int, double chan<- int) {
	for {
		select {
		// завершаем работу воркера, если прошел таймаут
		case <-ctx.Done():
			log.Println("time's up")
			return
		case value, ok := <-x:
			// проверяем закрыт ли канал
			if !ok {
				log.Println("closed channel")
				return
			}
			// удваиваем число и пишем в канал
			double <- value * 2
		}
	}
}

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	x := make(chan int, 1)
	double := make(chan int, 1)

	// создаем воркеров
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, x, double)
		}()
	}

	go func() {
		// ожидаем завершение воркеров
		wg.Wait()
		// закрываемканал double, тк
		close(double)
	}()

	// записываем в x числа из массива
	go func() {
		for _, num := range nums {
			x <- num
		}
		// закрываем канал, тк все числа записались
		close(x)
	}()

	// читаем из канала и выводим числа
	for num := range double {
		fmt.Println(num)
	}
}
