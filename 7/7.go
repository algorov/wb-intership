package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	//Mutex()
	SyncMap()
}

func Mutex() {
	// Инициализирует механизм синхронихации для горутин, критической секции и map.
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	mapa := make(map[int]int)

	// В цикле запускает горутину, которая записывает в map данные.
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(mx *sync.Mutex, mapa map[int]int) {
			defer wg.Done()

			key := rand.Int()
			value := rand.Int()

			// Критическая секция, которая контролируется мьютексом.
			mx.Lock()
			mapa[key] = value
			mx.Unlock()
		}(&mx, mapa)
	}

	// Ждёт завершения дочерних горутин.
	wg.Wait()

	// Вывод результата.
	fmt.Println(mapa)
}

func SyncMap() {
	// Инициализирует механизм синхронихации для горутин, критической секции и map.
	wg := sync.WaitGroup{}

	// Инициализация конкурентной мапы.
	var mapa sync.Map

	// В цикле запускает горутину, которая записывает в map данные.
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(mapa *sync.Map) {
			defer wg.Done()

			// Запись в map.
			mapa.Store(rand.Int(), rand.Int())
		}(&mapa)
	}

	// Ждёт завершения дочерних горутин.
	wg.Wait()

	// Вывод результата.
	fmt.Println(mapa)
}
