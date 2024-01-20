package main

import (
	"fmt"
	"sync"
)

/*
Задание:
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

// Глобальная переменная, с которой горутины будут взаимодействовать.
var sum int = 0

func main() {
	// Инициализируем массив с элементами.
	ns := [5]int{2, 4, 6, 8, 10}

	// Запуск вариаций реализации задачи.
	RunCaseOne(ns)
	RunCaseTwo(ns)
}

// Способ решения с sync.WaitGroup
func RunCaseOne(ns [5]int) {
	// Обнуляем глобальную переменную.
	sum = 0

	// Инициализируем механизмы синхронизации. Одна - для горутин, а другая - для критической секции.
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	// В процессе итерации по массиву, запускаем горутину и инкрементируем счетчик у WaitGroup.
	for _, numb := range ns {
		go growSumWaitGroup(numb, &wg, &mx)

		wg.Add(1)
	}

	// Главная горутина ждёт, пока дочерние не завершат свою работу.
	wg.Wait()

	// Вывод результата.
	fmt.Printf("Sum: %d\n", sum)
}

// Способ решения с chan.
func RunCaseTwo(ns [5]int) {
	// Обнуляем глобальную переменную.
	sum = 0

	// Инициализируем канал и Mutex.
	ch := make(chan struct{}, len(ns))
	mx := sync.Mutex{}

	// В процессе итерации по массиву, запускаем горутину.
	for _, numb := range ns {
		go growSumChannel(numb, ch, &mx)
	}

	// Главная горутина ждёт, пока дочерние не завершат свою работу.
	for i := 0; i < len(ns); i++ {
		<-ch
	}

	// Вывод результата.
	fmt.Printf("Sum: %d\n", sum)
}

func growSumWaitGroup(numb int, wg *sync.WaitGroup, mx *sync.Mutex) {
	// Очень интересная штука, которая гарантировано вызывает другую функцию, до возврата из функции.
	defer wg.Done()

	// Квадрат числа.
	val := numb * numb

	// Критическая секция, где синхронизируется посредством Mutex.
	mx.Lock()
	sum += val
	mx.Unlock()
}

func growSumChannel(numb int, ch chan<- struct{}, mx *sync.Mutex) {
	// Очень интересная штука, которая гарантировано вызывает другую функцию, до возврата из функции.
	defer func() {
		ch <- struct{}{}
	}()

	// Квадрат числа.
	val := numb * numb

	// Критическая секция, где синхронизируется посредством Mutex.
	mx.Lock()
	sum += val
	mx.Unlock()
}
