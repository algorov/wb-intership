package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	ns := []int{1, 3, 5, 7, 9, 11}

	//fmt.Println(CaseOne(ns, 3))
	//fmt.Println(ns)

	//fmt.Println(CaseOneCopy(ns, 3))
	//fmt.Println(ns)

	fmt.Println(CaseTwo(ns, 3))
	fmt.Println(ns)
}

// Если порядок не важен.
func CaseOne(ns []int, pos int) []int {
	// Проверка на валидность индекса.
	if pos >= 0 && pos < len(ns) {
		// Меняет местами удаляемый элемент с последним
		ns[pos], ns[len(ns)-1] = ns[len(ns)-1], ns[pos]

		// Возвращает слайс без последнего элемента.
		return ns[:len(ns)-1]
	}

	return ns
}

// Если порядок не важен и изменения не должны затрагивать исходный слайс.
func CaseOneCopy(ns []int, pos int) []int {
	// Проверка на валидность индекса.
	if pos >= 0 && pos < len(ns) {
		// Копирование в слайс с новым базовым массивом.
		copyNs := make([]int, len(ns))
		copy(copyNs, ns)

		// Меняет местами удаляемый элемент с последним
		copyNs[pos], copyNs[len(ns)-1] = copyNs[len(ns)-1], copyNs[pos]

		// Возвращает слайс без последнего элемента.
		return copyNs[:len(ns)-1]
	}

	return ns
}

// Если порядок важен.
func CaseTwo(ns []int, pos int) []int {
	// Проверка на валидность индекса.
	if pos >= 0 && pos < len(ns) {
		// Инициализация нового слайса.
		newNs := make([]int, 0, len(ns)-1)

		// Добавляет элементы, за исключением удаленного, в новый слайс.
		newNs = append(newNs, ns[:pos]...)
		newNs = append(newNs, ns[pos+1:]...)

		return newNs
	}

	return ns
}
