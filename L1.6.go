package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var stop bool
var mu sync.Mutex

func CtxRoutine(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\t\t\tгорутина остановлена контекстом\n")
			return
		default:
			n := rand.Intn(1000)
			fmt.Printf("контекстная работа, контекст номер %d\n", n)
			time.Sleep(time.Duration(n) * time.Millisecond)
		}
	}
}

func FlagRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mu.Lock()
		if stop {
			fmt.Printf("\t\t\tгорутина остановлена благодаря флагу\n")
			mu.Unlock()
			return
		}
		mu.Unlock()
		n := rand.Intn(1000)
		fmt.Printf("имитация с флагом, рандомное число: %d\n", n)
		time.Sleep(time.Duration(n/2) * time.Millisecond)
	}
}

func WgRoutine(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		t := rand.Intn(10000)
		fmt.Printf("имитирую работу, но мне осталось работать %d итераций, работа %d\n", n-i, t)
		time.Sleep(time.Duration(n*100) * time.Millisecond)
	}
	fmt.Printf("\t\t\tРутина с обычной WaitGroup прекратила свое выполнение после %d итераций\n", n)
}

func ChanRoutine(ch <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			fmt.Printf("\t\t\tгорутина остановлена после получения сигнала из канала\n")
			return
		default:
			num := rand.Intn(1000)
			fmt.Printf("имитирую работу, число %d\n", num)
			time.Sleep(time.Duration(num) * time.Millisecond)
		}
	}
}

func main() {
	var N, V, X, T int
	fmt.Println("введите число секунд, которое будет работать рутина с контекстом")
	_, err := fmt.Scan(&N)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Теперь же введите, сколько будет работать рутина с отменой по каналу")
	_, err = fmt.Scan(&V)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("время работы для рутины с отменой по флагу")
	_, err = fmt.Scan(&X)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Ну и количество итераций (не секунд!) работы для рутины с WaitGroup")
	_, err = fmt.Scan(&T)
	if err != nil {
		log.Println(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go WgRoutine(T, &wg)

	wg.Add(1)
	go CtxRoutine(ctx, &wg)

	wg.Add(1)
	go ChanRoutine(ch, &wg)

	wg.Add(1)
	go FlagRoutine(&wg)

	time.Sleep(time.Duration(X) * time.Second)
	mu.Lock()
	stop = true
	mu.Unlock()

	time.Sleep(time.Duration(N) * time.Second)
	cancel() // Отменяем контекст

	time.Sleep(time.Duration(V) * time.Second)
	close(ch) // Отправляем сигнал остановки по каналу

	wg.Wait() // Ждем завершения всех горутин

	fmt.Printf("\n\n\t\t\tвсе горутины закончили работу\n\n")
}
