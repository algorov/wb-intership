package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

// Структура-счетчик.
type Counter struct {
	count int
	mx    *sync.Mutex
}

// Метод, который инкрементирует счетчик в конкурентной среде.
func (c *Counter) Increment() {
	c.mx.Lock()
	c.count++
	c.mx.Unlock()
}

func main() {
	ch := make(chan struct{}, 10)
	defer close(ch)

	// Количество горутин.
	n := 1000

	// Инициализация счетчика.
	counter := Counter{
		count: 0,
		mx:    &sync.Mutex{},
	}

	for i := 0; i < n; i++ {
		go func(ctr *Counter, ch chan<- struct{}) {
			ctr.Increment()
			ch <- struct{}{}
		}(&counter, ch)
	}

	// Ожидание завершения всех горутин.
	for i := 0; i < n; i++ {
		<-ch
	}

	// Вывод результата.
	fmt.Println(counter.count)
}
