package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

// 7. Реализовать конкурентную запись данных в map.

type Task struct {
	// делаем мапу безопасной для конкурентности
	mu    sync.RWMutex
	tasks map[string]bool
}

func NewTask() *Task {
	return &Task{
		tasks: make(map[string]bool),
	}
}

// добавление задачи в tasks
func (t *Task) Add(name string) {
	// блокируем мьютекс
	t.mu.Lock()
	// разблокировка
	defer t.mu.Unlock()
	// безопасно добавляем k, v
	t.tasks[name] = false
}

// меняем статус задачи на true
func (t *Task) Done(name string) error {
	// блокируем мьютекс
	t.mu.Lock()
	// разблокировка
	defer t.mu.Unlock()

	// проверка на наличие ключа
	_, ok := t.tasks[name]
	if !ok {
		return errors.New("task not found")
	}

	// безопасно меняем значение
	t.tasks[name] = true
	return nil
}

// вывод статуса задачи
func (t *Task) PrintStatus(name string) error {
	// блокируем на чтение
	t.mu.RLock()
	// разблокировка
	defer t.mu.RUnlock()

	// читаем из мапы
	v, ok := t.tasks[name]
	if !ok {
		return errors.New("task not found")
	}

	fmt.Println(name, v)
	return nil
}

func main() {
	var wg sync.WaitGroup
	task := NewTask()

	// синхронно добавляем задачи
	for i := 0; i < 10; i++ {
		task.Add(fmt.Sprintf("Task %v", i))
	}

	// рандомный вызов 20 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if err := task.Done(fmt.Sprintf("Task %v", num)); err != nil {
				log.Println(err)
			}
		}(i)

		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if err := task.PrintStatus(fmt.Sprintf("Task %v", num)); err != nil {
				log.Println(err)
				return
			}
		}(i)
	}

	// ожидаем завершение всех горутин
	wg.Wait()
}
