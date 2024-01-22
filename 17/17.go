package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	// Инициализация последовательности
	ns := []int{1, 4, 43, 65, 132, 999, 4353}

	// Вызывает алгоритм.
	fmt.Println(BinarySearch(ns, 43))
}

func BinarySearch(ns []int, find int) int {
	// Границы поиска.
	left, right := 0, len(ns)-1

	for left <= right {
		middle := (left + right) >> 1
		// Если предполагаемое число равно нужному.
		if ns[middle] == find {
			return middle
		}

		// Если предполагаемое число больше нужного, то сдвигает границу.
		if ns[middle] > find {
			right = middle - 1
			continue
		}

		// Если предполагаемое число меньге нужного, то сдвигает границу.
		if ns[middle] < find {
			left = middle + 1
			continue
		}
	}

	// Если такого числа нет в последовательности.
	return -1
}
