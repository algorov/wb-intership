package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	// Инициализирует последовательность строк.
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	// Инифиализирует множество строк.
	set := make(map[string]struct{})

	// В процессе итерации добавляет строку в множество, если её нет там.
	for _, str := range strs {
		if _, ok := set[str]; !ok {
			set[str] = struct{}{}
		}
	}

	// Вывод результата.
	fmt.Println(set)
}
