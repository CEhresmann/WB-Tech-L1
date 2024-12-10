package main

import (
	"fmt"
)

// Интерфейс для целочисленного калькулятора
type IntCalculator interface {
	AddInt(a, b int) int
	SubtractInt(a, b int) int
}

// Реализация целочисленного калькулятора
type SimpleIntCalculator struct{}

func (c *SimpleIntCalculator) AddInt(a, b int) int {
	return a + b
}

func (c *SimpleIntCalculator) SubtractInt(a, b int) int {
	return a - b
}

// Адаптер для целочисленного калькулятора, чтобы он соответствовал интерфейсу для работы с вещественными числами
// Позволяет использовать целочисленный калькулятор в контексте, где ожидаются вещественные числа.
type IntToFloatAdapter struct {
	intCalculator *SimpleIntCalculator
}

func (a *IntToFloatAdapter) Add(aFloat, bFloat float64) float64 {
	return float64(a.intCalculator.AddInt(int(aFloat), int(bFloat)))
}

func (a *IntToFloatAdapter) Subtract(aFloat, bFloat float64) float64 {
	return float64(a.intCalculator.SubtractInt(int(aFloat), int(bFloat)))
}

// Реализация калькулятора с плавающей запятой
type SimpleFloatCalculator struct{}

func (c *SimpleFloatCalculator) Add(a, b float64) float64 {
	return a + b
}

func (c *SimpleFloatCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func main() {
	intCalculator := &SimpleIntCalculator{}
	intAdapter := &IntToFloatAdapter{intCalculator: intCalculator}

	var a int
	fmt.Print("Введите целое число: ")
	fmt.Scanf("%d", &a)

	fmt.Println("Используем адаптер для работы калькулятора, принимающего только вещественные числа с целыми числами")
	fmt.Println("Плюс: ", intAdapter.Add(float64(a), 345))
	fmt.Println("Минус: ", intAdapter.Subtract(121, 3))

	// Создаем калькулятор с плавающей запятой
	floatCalculator := &SimpleFloatCalculator{}
	fmt.Println("Также можно использовать структуру SimpleFloatCalculator")
	fmt.Println("Плюс: ", floatCalculator.Add(512.0, 345.0))
	fmt.Println("Минус: ", floatCalculator.Subtract(float64(a), 35.1))
}
