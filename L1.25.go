package main

import (
	"fmt"
	"log"
	"syscall"
)

func sleep(seconds int) {
	var ts syscall.Timespec
	ts.Sec = int64(seconds)
	ts.Nsec = 0

	err := syscall.Nanosleep(&ts, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Начинаем ожидание...")
	sleep(3)
	fmt.Println("Ожидание завершено!")
}
