package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	// Задает исходную строку.
	text := "snow dog sun"

	// Разбивает на токены.
	tokens := strings.Split(text, " ")
	Swipe(tokens)

	// Сборка и вывод строки.
	fmt.Println(strings.Join(tokens, " "))
}

func Swipe(tokens []string) {
	// Индекс последнего элемента слайса.
	last := len(tokens) - 1

	// Итерируется до середины слайса, путем хитростей значения меняются.
	for i := 0; i <= len(tokens)/2; i++ {
		tokens[i], tokens[last-i] = tokens[last-i], tokens[i]
	}
}
