package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	//Timer()
	//Channel()
	//ContextTimeOut()
	//ContextCancel()
}

func Timer() {
	// Инициализируем канал, для отслеживания состояния горутины.
	ch := make(chan struct{})
	defer close(ch)

	// Количество секунд для таймера.
	n := 5

	go func() {
		defer func() {
			fmt.Println("Dead.")
			ch <- struct{}{}
		}()

		// После истичения времени возвращается канал.
		chTimer := time.After(time.Duration(n) * time.Second)

		for {
			select {
			// Если время истекло, то происходит возврат из горутины.
			case <-chTimer:
				return
			default:
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Work...")
			}
		}
	}()

	// Главная горутина ждёт данные из канала.
	<-ch
}

func Channel() {
	// Инициализируем 2 канала (для принятия / отправки сигналов)
	chIn := make(chan struct{})
	chOut := make(chan struct{})
	defer func() {
		close(chIn)
		close(chOut)
	}()

	// Запускаем горутину.
	go func() {
		defer func() {
			fmt.Println("Dead.")
			// Отправляет сигнал о завершении в канал.
			chIn <- struct{}{}
		}()

		for {
			select {
			// Если есть сигнал завершения, горутина завершается.
			case <-chOut:
				return
			default:
				fmt.Println("Work...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Ждёт для наглядности.
	time.Sleep(4 * time.Second)

	// Отправляет в канал сигнал о завершении.
	chOut <- struct{}{}
	// Ждёт сигнал о завершении горутины.
	<-chIn
}

func ContextTimeOut() {
	// Инициализирует контекст с таймаутом.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Инициализирует канал для получения сигнала завершения от другой горутины.
	ch := make(chan struct{})
	defer cancel()

	go func() {
		defer func() {
			fmt.Println("Dead.")
			ch <- struct{}{}
		}()

		for {
			select {
			// Проверяет на наличие отмены контекста.
			case <-ctx.Done():
				return
			default:
				fmt.Println("Work...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Главная горутина ждёт сигнала из канала.
	<-ch
}

func ContextCancel() {
	// Инициализирует контекст с таймаутом.
	ctx, cancel := context.WithCancel(context.Background())

	// Инициализирует канал для получения сигнала завершения от другой горутины.
	ch := make(chan struct{})

	go func() {
		defer func() {
			fmt.Println("Dead.")
			ch <- struct{}{}
		}()

		for {
			select {
			// Проверяет на наличие отмены контекста.
			case <-ctx.Done():
				return
			default:
				fmt.Println("Work...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// После истичения 5 секунд, отменяет контекст.
	time.Sleep(5 * time.Second)
	cancel()

	// Главная горутина ждёт сигнала из канала.
	<-ch
}
