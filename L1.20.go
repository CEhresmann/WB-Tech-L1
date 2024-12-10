package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func revWords(s string) string {
	words := strings.Fields(s)
	narr := make([]string, 0, len(words))

	for i := len(words) - 1; i >= 0; i-- {
		narr = append(narr, words[i])
	}

	return strings.Join(narr, " ")
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку, которую хотите перевернуть. По окончании ввода нажмите клавишу ENTER и введите q!: ")

	var input string
	for reader.Scan() {
		l := reader.Text()
		if l == "q!" {
			break
		}
		input += l + " "
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Перевёрнутая строка:", revWords(strings.TrimSpace(input)))
}
