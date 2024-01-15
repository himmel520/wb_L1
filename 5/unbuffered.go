package main

import (
	"log"
	"time"
)

func unbufferedChan() {
	// создаем небуферизированный канал
	ch := make(chan time.Time)
	// обьявляем таймер
	timeout := time.After(1 * time.Second)

	// горутина завершится, когда закончится main
	go func() {
		for {
			ch <- time.Now()
		}
	}()

	for {
		select {
		// выход, по истечении таймера
		case <-timeout:
			log.Println("time is up")
			return
		// читаем из канала
		case mytime := <-ch:
			log.Println(mytime)
		}
	}
}
