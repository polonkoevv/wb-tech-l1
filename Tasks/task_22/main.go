package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация больших чисел
	a := new(big.Int)
	b := new(big.Int)

	// Установка значений (например, 2^21)
	a.SetString("2097152", 10) // 2^21
	b.SetString("4194304", 10) // 2^22

	// Создание переменных для хранения результатов
	sum := new(big.Int)
	sub := new(big.Int)
	mul := new(big.Int)
	quotient := new(big.Int)
	remainder := new(big.Int)

	// Выполнение операций
	sum.Add(a, b)
	sub.Sub(a, b)
	mul.Mul(a, b)
	quotient.QuoRem(a, b, remainder)

	// Вывод результатов
	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Printf("a + b = %s\n", sum.String())
	fmt.Printf("a - b = %s\n", sub.String())
	fmt.Printf("a * b = %s\n", mul.String())
	fmt.Printf("a / b = %s\n", quotient.String())
	fmt.Printf("a %% b = %s\n", remainder.String())
}
