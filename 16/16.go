package main

import "fmt"

/*
	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	// Инициализация неупорядоченной последовательности
	ns := []int{1, 4665, 12, 442, 654, 99}

	// Запускает алгоритм.
	QuickSort(ns)

	// Выводит результат.
	fmt.Println(ns)
}

func QuickSort(ns []int) {
	// Условие остановки рекурсии.
	if len(ns) > 1 {
		// Сортирует и вычисляет индекс опорного элемента.
		partPosition := partition(ns)

		// Рекурсивно вычисляет обе части от опорного элемента.
		QuickSort(ns[:partPosition])
		QuickSort(ns[partPosition+1:])
	}
}

func partition(ns []int) int {
	// Граница, где с левой стороны меньшие элементы от опорноного, а с правой - большие.
	side := 0
	right := len(ns) - 1

	// Опорный элемент.
	pivot := ns[right]

	// Итерируется по всей последовательности.
	for pointer := 0; pointer < len(ns); pointer++ {
		// Если рассматриваемый элемент меньше опорного, то нам нужно поместить в левую сторону от side.
		// Это делается посредством инкрементации переменной side.
		// Также возможно, что между рассматриваемыми элементами есть элементы, которые больше опорного.
		// В таком случае меняет значения местами.
		if ns[pointer] < pivot {
			if pointer != side {
				ns[pointer], ns[side] = ns[side], ns[pointer]
			}

			side++
			continue
		}

		// Когда алгоритм рассмотрел все элементы, опорный элемент меняется местом с правостоящим от side элементом.
		if ns[pointer] == pivot {
			ns[side], ns[pointer] = ns[pointer], ns[side]
		}
	}

	// Возвращает индекс side.
	return side
}