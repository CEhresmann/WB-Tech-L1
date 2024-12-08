package main

import (
	"fmt"
	"sync"
)

//Форматирование условия задачи на сайте некорректно, заместо 2^2 отображается 22

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	results := make(chan int, len(numbers))

	// Запускаем горутины для вычисления квадратов
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- n * n
		}(num)
	}

	// Ждем завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Суммируем результаты
	sum := 0
	for result := range results {
		sum += result
	}

	fmt.Println(sum)
}
