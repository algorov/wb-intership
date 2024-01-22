package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func main() {
	// Инициализирует сканер, который считает строку.
	scan := bufio.NewScanner(os.Stdin)

	for {
		scan.Scan()
		text := scan.Text()

		// Переводит строку в представление рун.
		runes := []rune(text)
		Swipe(runes)

		// Сборка строки.
		fmt.Println(string(runes))
	}
}

func Swipe(runes []int32) {
	// Индекс последнего элемента слайса.
	last := len(runes) - 1

	// Итерируется до середины слайса, путем хитростей значения меняются.
	for i := 0; i <= len(runes)/2; i++ {
		runes[i], runes[last-i] = runes[last-i], runes[i]
	}
}
