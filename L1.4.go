package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(id int, ctx context.Context, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				// Если канал закрыт, выходим из горутины
				return
			}
			fmt.Printf("Воркер %d делает работу (все числа случайны): %s\n", id, job)
		case <-ctx.Done():
			// Если контекст завершен, выходим из горутины
			return
		}
	}
}

func main() {
	// Устанавливаем количество воркеров
	var numWorkers int
	fmt.Print("Введите количество воркеров ")
	_, err := fmt.Scan(&numWorkers)
	if err != nil {
		log.Println(err)
	}
	jobs := make(chan string)

	var wg sync.WaitGroup
	// Создаем контекст для завершения работы воркеров
	ctx, cancel := context.WithCancel(context.Background())
	// Запускаем воркеры
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, ctx, jobs, &wg)
	}

	// Обработка сигнала завершения
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		fmt.Println("\nReceived interrupt signal, shutting down...")
		cancel()
		close(jobs)
	}()

	go func() {
		for {
			job := fmt.Sprintf("Вот такое %d число!", rand.Intn(100))
			jobs <- job
			time.Sleep(1 * time.Second) // Задержка для имитации работы
		}
	}()
	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}
