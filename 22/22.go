package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

type BigInteger struct {
	a, b *big.Int
}

// Инициализирует структуру.
func New(a, b string) *BigInteger {
	integer := BigInteger{new(big.Int), new(big.Int)}
	integer.a.SetString(a, 10)
	integer.b.SetString(b, 10)

	return &integer
}

// Суммирует первое число со вторым.
func (b *BigInteger) Add() *big.Int {
	return new(big.Int).Add(b.a, b.b)
}

// Вычитает первое второе число из первого.
func (b *BigInteger) Sub() *big.Int {
	return new(big.Int).Sub(b.a, b.b)
}

// Умножает первое число со вторым.
func (b *BigInteger) Mul() *big.Int {
	return new(big.Int).Mul(b.a, b.b)
}

// Делит первое число на второе.
func (b *BigInteger) Div() *big.Int {
	return new(big.Int).Div(b.a, b.b)
}

func main() {
	bigInt := New("623042860176616787109600504394185787518921", "18000000003534613450000034234234123423400000")

	// Демонстрация работы программы.
	fmt.Println(bigInt.Add())
	fmt.Println(bigInt.Add())
	fmt.Println(bigInt.Sub())
	fmt.Println(bigInt.Div())
	fmt.Println(bigInt.Mul())
}
