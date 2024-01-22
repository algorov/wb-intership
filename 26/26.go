package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	// Последовательность примеров.
	examples := []string{"abcd", "abCdefAaf", "aabcd"}

	// Проверка на работоспособность алгоритма.
	for _, example := range examples {
		fmt.Println(IsUnique(example))
	}
}

func IsUnique(word string) bool {
	// Устанавливает все руны в нижний регистр.
	word = strings.ToLower(word)

	// Перевод представления на слайс рун.
	runes := []rune(word)

	// Мапа, чтобы хранить уникальные элементы.
	set := make(map[int32]struct{})

	// Итерируется по последовательности рун.
	// Если в мапе под таким ключом отстутствует, то добавляем, иначе последовательность не уникальна.
	for _, symb := range runes {
		if _, isExist := set[symb]; !isExist {
			set[symb] = struct{}{}
			continue
		}

		return false
	}

	return true
}
