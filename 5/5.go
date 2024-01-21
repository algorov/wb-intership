package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	// Переменная, которая хранит количество секунд.
	var n int

	// Получение значения из stdin.
	fmt.Print("Enter the second count: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println(err)

		return
	}

	// Инициализируем механизм синхронизации.
	wg := sync.WaitGroup{}

	// Создаем небуферизированный канал.
	ch := make(chan int)
	defer close(ch)

	// Создаем контекст с таймаутом, по истечении которого контекст будет закрыт.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	// Запускаем обработчика для чтения из канала.
	wg.Add(1)
	go SomeSide(ctx, ch, &wg)

	// Запускаем горутину, которая записывает в канал. Горутина завершается, если контекст отменен.
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Генерирует числа и записывает в канал.
				ch <- rand.Intn(100)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Ждёт завершения всех дочерних горутин.
	wg.Wait()
}

func SomeSide(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// В бесконечном цикле проверяет: если контекст отменен, то горутина подходит к завершению.
	// Если в канале что-то есть, происходит чтение.
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		case data := <-ch:
			fmt.Println(data)
		}
	}
}
