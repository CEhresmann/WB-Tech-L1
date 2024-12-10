package main

import (
	"fmt"
	"log"
)

func determineType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Тип переменной: int, значение: %d\n", v)
	case string:
		fmt.Printf("Тип переменной: string, значение: %s\n", v)
	case bool:
		fmt.Printf("Тип переменной: bool, значение: %t\n", v)
	case chan struct{}:
		fmt.Printf("Тип переменной: channel\n")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	var a interface{}

	fmt.Print("Введите целое число: ")
	var intValue int
	_, err := fmt.Scan(&intValue)
	if err == nil {
		a = intValue
		determineType(a)
		return
	}

	fmt.Print("Введите булево значение (true/false): ")
	var boolValue bool
	_, err = fmt.Scan(&boolValue)
	if err == nil {
		a = boolValue
		determineType(a)
		return
	}

	fmt.Print("Введите строку: ")
	var stringValue string
	_, err = fmt.Scan(&stringValue)
	if err == nil {
		a = stringValue
		determineType(a)
		return
	}

	log.Println("Не удалось считать значение.")
}
