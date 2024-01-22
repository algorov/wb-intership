package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	/*
		Размышления. Тут какая-то неопределенность: может вернуться как []rune, так и string.
		Можно явно конвертировать в строку - тогда создадутся новые данные и ссылки на huge string не будет. А дальше GC
		сделает своё дело.
	*/

	justString = string(v[:100])
}

func main() {
	someFunc()
}
