package main

import "fmt"

func squareChan() {
	// создаем небуферизированный канал
	square := make(chan int)
	nums := [5]int{2, 4, 6, 8, 10}

	// итерируемся по массиву
	for _, v := range nums {
		go func(num int) {
			// блокируемся, пока данные не будут получены
			square <- num * num
		}(v)
	}

	for range nums {
		// блокируемся, пока не прочитаем из канала
		fmt.Println(<-square)
	}
}
