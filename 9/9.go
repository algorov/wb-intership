package main

import (
	"fmt"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	// Инициализируем массив из 10 элементов.
	ns := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Два буферизированных канала.
	chResult := make(chan int, len(ns))
	chOut := make(chan int, len(ns))
	defer func() {
		close(chResult)
		close(chOut)
	}()

	// Итерируется по массиву, где в процессе каждый элемент записывается в канал.
	// Далее запускается горутина, которая читает значение из этого канала, делает вычисления и отправляет в другой канал,
	for _, numb := range ns {
		chOut <- numb

		go func(chIn <-chan int, chOut chan<- int) {
			value := <-chIn
			chOut <- value * value
		}(chOut, chResult)
	}

	// Читаем из другого канала.
	for i := 0; i < len(ns); i++ {
		fmt.Println(<-chResult)
	}
}
