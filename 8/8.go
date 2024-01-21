package main

import (
	"fmt"
	"math/rand"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var i, isSet int

	// Ввод позиции бита, с которым будет работать программа.
	fmt.Print("Position: ")
	if _, err := fmt.Scan(&i); err != nil {
		fmt.Println(err)
		return
	}

	// Проверка корректности введенных данных.
	if i > 63 {
		fmt.Println("Overflow discharge!")
		return
	} else if i < 0 {
		fmt.Println("Negative!")
		return
	}

	// Ввод действия.
	fmt.Print("Reset (0) or set (1): ")
	if _, err := fmt.Scan(&isSet); err != nil {
		fmt.Println(err)
		return
	}

	// Проверка корректности введенных данных.
	if isSet != 0 && isSet != 1 {
		fmt.Println("Incorrect!")
		return
	}

	// Генерация числа.
	numb := rand.Int63()
	fmt.Println("Numb: ", numb)

	// Определение действия.
	if isSet == 1 {
		numb = SetBit(numb, i)
	} else {
		numb = ResetBit(numb, i)
	}

	// Вывод результата.
	fmt.Println("Result: ", numb)
}

func SetBit(numb int64, pos int) int64 {
	// Применяет побитовую операцию (ИЛИ) с маской и возвращает результат.
	return numb | (1 << pos)
}

func ResetBit(numb int64, pos int) int64 {
	// Применяет побитовую операцию (И-НЕ) с маской и возвращает результат.
	return numb &^ (1 << pos)
}
