package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел,
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	// Создаем массив из пяти элементов.
	ms := [5]int{2, 4, 6, 8, 10}

	// Тут мы запускаем разные реализации задачи.
	RunCaseOne(ms)
	RunCaseTwo(ms)
}

// Способ решения с sync.WaitGroup
func RunCaseOne(ms [5]int) {
	// Создаем механизм синхронизации.
	wg := sync.WaitGroup{}

	// Итерируемся по массиву. В процессе итерации число передаем функции, которая отработает конкурентно.
	// Счетчик для механизма увеличиваем.
	for _, numb := range ms {
		go procWaitGroup(numb, &wg)
		wg.Add(1)
	}

	// Главная горутина ждёт, пока завершатся дочерние горутины.
	wg.Wait()
}

// Способ решения с chan.
func RunCaseTwo(ms [5]int) {
	// Инициализируем буферизированный канал.
	chIn := make(chan struct{}, len(ms))

	// Итерируемся по массиву. В процессе итерации число передаем функции, которая отработает конкурентно.
	// Также передаем сам канал.
	for _, numb := range ms {
		go procChannel(numb, chIn)
	}

	// Сама суть синхронизации. Тут мы читаем ровно столько раз, сколько элементов в массиве.
	// Каждый раз главная горутина будет ждать данные из канала, в который записывают дочерние горутины.
	for i := 0; i < len(ms); i++ {
		// Само чтение из канала.
		<-chIn
	}
}

func procWaitGroup(numb int, wg *sync.WaitGroup) {
	// Вывод в stdout.
	fmt.Printf("Numb: %d; ^2: %d\n", numb, numb*numb)

	// Вызываем метод, тогда счетчик декрементируется,
	wg.Done()
}

func procChannel(numb int, ch chan<- struct{}) {
	// Вывод в stdout.
	fmt.Printf("Numb: %d; ^2: %d\n", numb, numb*numb)

	// Записываем в канал пустую структуру.
	// Да, это самая дешевая операция относительно памяти, поскольку пустая структура весит 0 байт.
	ch <- struct{}{}
}
