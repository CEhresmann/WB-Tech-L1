package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sender(chanel chan int) {
	for {
		num := rand.Intn(100)
		chanel <- num
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	fmt.Println("Введите количество секунд, через которое канал должен быть закрыт")
	var N int
	fmt.Scan(&N)
	timer := time.After(time.Duration(N) * time.Second)

	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	go sender(ch)
	for {
		select {
		case <-timer:
			fmt.Println("Время вышло! Сворачиваюсь...")
			close(ch)
			return
		case num := <-ch:
			fmt.Printf("Получено значение %d\n", num)
		}
	}
}
