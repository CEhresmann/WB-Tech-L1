package main

import (
	"fmt"
	"log"
)

func setBit(num int64, i int) int64 {
	return num | (1 << i)
}

func clearBit(num int64, i int) int64 {
	return num &^ (1 << i)
}

func isSet(num int64, i int) bool {
	return (num & (1 << i)) != 0
}

func main() {
	var num int64
	var i int
	var action string

	fmt.Println("Введите число (int64):")
	_, err := fmt.Scan(&num)
	if err != nil {
		log.Fatal("тип вводимых данных указан в скобках, неверный ввод ", err)
	}

	fmt.Println("Введите индекс бита (i):")
	_, err1 := fmt.Scan(&i)
	if err1 != nil {
		log.Fatal("тип вводимых данных указан в скобках, неверный ввод ", err)
	}
	fmt.Println("Введите действие (set для установки в 1, clear для установки в 0):")
	_, err2 := fmt.Scan(&action)
	if err2 != nil {
		log.Fatal("тип вводимых данных указан в скобках, неверный ввод ", err)
	}
	switch action {
	case "set":
		num = setBit(num, i)
		fmt.Printf("Число после установки %d-го бита в 1: %d\n", i, num)
	case "clear":
		num = clearBit(num, i)
		fmt.Printf("Число после установки %d-го бита в 0: %d\n", i, num)
	default:
		fmt.Println("Неверное действие. Используйте 'set' или 'clear'.")
		return
	}

	if isSet(num, i) {
		fmt.Printf("%d-й бит установлен в 1\n", i)
	} else {
		fmt.Printf("%d-й бит установлен в 0\n", i)
	}
}
