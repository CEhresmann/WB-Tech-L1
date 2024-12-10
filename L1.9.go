package main

import (
	"fmt"
	"sync"
)

func write(numbers []int, inputChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range numbers {
		inputChannel <- num
	}
	close(inputChannel)
}

func double(inputChannel <-chan int, outputChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range inputChannel {
		outputChannel <- num * 2
	}
	close(outputChannel)
}

func Res(outputChannel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range outputChannel {
		fmt.Println(result)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go write(numbers, input, &wg)

	wg.Add(1)
	go double(input, output, &wg)

	wg.Add(1)
	go Res(output, &wg)

	wg.Wait()
}
