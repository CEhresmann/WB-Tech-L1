package main

import (
	"fmt"
	"math/big"
)

type BigIntCalcus interface {
	Add(a, b *big.Int) *big.Int
	Subtract(a, b *big.Int) *big.Int
	Multiply(a, b *big.Int) *big.Int
	Divide(a, b *big.Int) (*big.Int, error)
}

type BIC struct{}

// сложение
func (c *BIC) Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// вычитание
func (c *BIC) Subtract(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// умножение
func (c *BIC) Multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// деление
func (c *BIC) Divide(a, b *big.Int) (*big.Int, error) {
	if b.Sign() == 0 {
		return nil, fmt.Errorf("деление на ноль")
	}
	return new(big.Int).Div(a, b), nil
}

func main() {
	var a, b big.Int
	var calculator BigIntCalcus = &BIC{}

	fmt.Print("Введите значение a (больше 2^20): ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Print("Введите значение b (больше 2^20): ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Проверка, что значения больше 2^20
	twentyPower := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)
	if a.Cmp(twentyPower) <= 0 || b.Cmp(twentyPower) <= 0 {
		fmt.Println("Оба значения должны быть больше 2^20 (1048576).")
		return
	}

	sum := calculator.Add(&a, &b)
	diff := calculator.Subtract(&a, &b)
	mul := calculator.Multiply(&a, &b)
	div, err := calculator.Divide(&a, &b)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("a + b = %s\n", sum.String())
	fmt.Printf("a - b = %s\n", diff.String())
	fmt.Printf("a * b = %s\n", mul.String())
	fmt.Printf("a / b = %s\n", div.String())
}
