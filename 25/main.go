package main

import (
	"fmt"
	"time"
)

// 25. Реализовать собственную функцию sleep.

func Sleep(d time.Duration) {
	// блокируемся, пока указанное время не пройдет и пишем в канал
	<-time.After(d)
}

func main() {
	fmt.Println(time.Now())
	Sleep(2 * time.Second)
	fmt.Println(time.Now())
}
