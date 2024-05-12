package main

import (
	"fmt"
	"log"
)

/*К каким негативным последствиям может привести
данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

*/

// var justString string // Использование глобальной переменной justString — плохая практика

// func someFunc() { // В случае возникновение ошики 
// 	v := createHugeString(1 << 10) // Не корректный способ передавать длину строки через побитовый сдвиг
// 	justString = v[:100]           // Может быть вызвана панику из-за выхода за пределы массива
// }

func main() {
	err := bestSomeFunc()
	if err != nil {
		log.Println(err)
	}
}

func bestSomeFunc() error { // Обработка ошибок позволяет программе избежать нежелательного поведения или аварийного завершения
	v := createHugeString(1 << 10)
	if len(v) >= 100 { // Проверка на потенциальную панику из-за выхода за пределы массива при присваивании justString = v[:100]
		return nil
	}
	return fmt.Errorf("Substring length less than 100") 
}

func createHugeString(size int) string {
	if size < 0 { // Необходимо проверить, что размер строки в createHugeString неотрицательный.
		log.Fatalln("Invalid size!")
	}

	res := make([]byte, size)
	for i := range res {
		res[i] = byte(i + 97) // наполнение строки символами a, b, c, ...
	}
	return string(res)
}
