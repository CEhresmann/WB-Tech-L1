package main

import (
	"fmt"
	"sort"
	"sync"
)

func square(n int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	result <- n * n
}

func main() {
	var wg sync.WaitGroup

	var arr = []int{2, 4, 6, 8, 10}
	var out = make([]int, 0, len(arr))
	res := make(chan int, len(arr))

	for _, n := range arr {
		wg.Add(1)
		go square(n, &wg, res)
	}

	wg.Wait()
	close(res)
	for n := range res {
		out = append(out, n)
	}
	sort.Ints(out) //Сортируем для вывода в нужном порядке, просто вывод результатов из канала дал бы: 100\n 4\n 16\n 36\n 64
	for _, n := range out {
		fmt.Println(n) //В данном случае все преимущества конкурентного выполнения теряются из-за I/O burst
		//Но если бы вместо функции возведения в квадрат стояло что-то более тяжелое, это имело бы смысл
	}
}
