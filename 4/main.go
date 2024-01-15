package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

/*
4. Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные
данные из канала и выводят в stdout.

Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func worker(data <-chan int, id int) {
	// читаем из канала
	for {
		num, ok := <-data
		// проверяем закрыт ли канал
		if !ok {
			log.Printf("worker %v exit", id)
			return
		}
		log.Printf("worker %v; value: %v\n", id, num)
	}

}

func main() {
	// получаем аргументы командной строки
	args := os.Args

	if len(args) == 1 || len(args) > 2 {
		log.Fatal("invalid count of args")
	}

	// получаем количество воркеров
	n, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("invalid type of arg")
	}

	data := make(chan int)
	exit := make(chan os.Signal, 1)
	// отслеживание ctrl+c и отправка в канал exit
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < n; i++ {
		// запускаем воркер и передаем его номер
		go worker(data, i)
	}

	for {
		select {
		// при ctrl+c читаем сигнал и выходим из main
		case <-exit:
			log.Println("exit..")
			// закрываем канал
			close(data)
			return
		// отправляем в data рандомные числа от 0 до 1000
		case data <- rand.Intn(1000):
		}
	}

}
