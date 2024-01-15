package main

import "fmt"

func squareSumChan() {
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

	var sum int
	for range nums {
		// блокируемся, пока не прочитаем из канала
		sum += <-square
	}

	fmt.Println(sum)
}
