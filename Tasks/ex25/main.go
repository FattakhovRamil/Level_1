package main

/*
	Реализовать собственную функцию sleep.
*/

import (
	"fmt"
	"time"
)

func Sleep(sec int) {
	duration := time.Duration(sec) // длительность времени
	<-time.After(duration * time.Second) // длительность умножаем на секунды = 3 сек задерки основной горутины
}

func main() {

	fmt.Println("Начало работы программы")
	Sleep(3) // 3 секунды задержка
	fmt.Println("конец работы программы")

}
