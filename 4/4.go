package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	// Переменная, которая хранит количество воркеров.
	var n int

	// Получение значения из stdin.
	fmt.Print("Enter the worker count: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println(err)

		return
	}

	// Создаем буферизированный канал, куда Sender будет записывать числа.
	ch := make(chan int, 10)
	defer close(ch)

	// Не забываем про механизм синхронизации горутин.
	wg := sync.WaitGroup{}

	// Создаем контекст, который автоматически закроется, если мы принудительно завершим программу (CTRL+C).
	sigCtx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	// Запускаем воркеров.
	for i := 0; i < n; i++ {
		go Reciever(sigCtx, ch, &wg)
		wg.Add(1)
	}

	// Запускаем Sender.
	go Sender(sigCtx, ch, &wg)
	wg.Add(1)

	// Главная горутина ждёт завершения всех дочерних горутин.
	wg.Wait()
}

func Sender(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// В бесконечном цикле через select проверяем, закрылся ли контекст, если нет, то записываем в канал.
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- rand.Int()
		}
	}
}

func Reciever(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// В бесконечном цикле проверяем о готовности каналов.
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-ctx.Done():
			fmt.Println("Exit")
			//wg.Done()

			return
		case <-time.After(100 * time.Millisecond):
			fmt.Println("TimeOut")
		}
	}
}
