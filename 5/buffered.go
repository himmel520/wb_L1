package main

import (
	"log"
	"time"
)

func bufferedChan() {
	// создаем буферизированный канал
	ch := make(chan time.Time, 1)
	// обьявляем таймер
	timeout := time.After(1 * time.Second)
	for {
		select {
		// выход, по истечении таймера
		case <-timeout:
			log.Println("time is up")
			return
		// читаем из канала
		case mytime := <-ch:
			log.Println(mytime)
		// пишем в канал
		case ch <- time.Now():
		}
	}
}
