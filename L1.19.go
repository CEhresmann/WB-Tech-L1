package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func rev(s string) string {
	r := []rune(s)
	ln := len(r)
	for i := 0; i < ln/2; i++ {
		r[i], r[ln-i-1] = r[ln-i-1], r[i]
	}
	return string(r)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку которую хотите перевернуть, после нажмите ENTER: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Перевёрнутая строка:", rev(str))
}
