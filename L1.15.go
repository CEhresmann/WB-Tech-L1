package main

import (
	"fmt"
	"strings"
)

//Функция создает очень большую строку (1 КБ).
//Если эта строка создается в функции someFunc, а затем присваивается переменной justString, то в памяти будет храниться не только первые 100 символов, но и вся строка, которая была создана.
//Это может привести к избыточному потреблению памяти.
//
//Если createHugeString создает строки в цикле или многократно вызывается, то может появиться накопление неиспользуемых строк в памяти,
//что в свою очередь может вызвать утечки памяти.

func createHugeString(size int) string {
	// Используем strings.Builder для эффективного создания строки
	var builder strings.Builder
	for i := 0; i < size; i++ {
		builder.WriteByte('W') // Заполняем строку символами 'W'
		builder.WriteByte('B')
		builder.WriteByte(' ')
	}
	return builder.String()
}

var justString *string

func someFunc() {
	v := createHugeString(1 << 10) // Создаем строку размером 1 КБ
	subString := v[:100]           // Сохраняем только первые 100 символов
	justString = &subString
}

func main() {
	someFunc()
	fmt.Println("Сохраненная строка:", *justString)
}
