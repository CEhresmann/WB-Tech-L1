package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10
	//математически
	a = a + b
	b = a - b
	a = a - b
	//с использованием побитовых операций
	a = a ^ b
	b = a ^ b
	a = a ^ b
	//с использованием множественного присваивания
	a, b = b, a
	fmt.Println("После обмена:", a, b)
}
