/*
	Реализовать бинарный поиск встроенными методами языка.
*/

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var N int = 15
	array := make([]int, N)

	for i := 0; i < N; i++ {
		array[i] = rand.Intn(33) - 16 // от -16 до 16
	}

	fmt.Println(array)

	quickSort(array)
	fmt.Println(array)
	elem := 5
	index, err := binSearch(array, 5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Элемент %d найден по индексу: %d\n", elem, index)
	}

}

func quickSort(array []int) {
	quickSortMain(array, 0, len(array)-1)
}

func quickSortMain(array []int, left int, right int) {
	if left > right {
		return
	}
	aver := array[(left+right)/2]

	i, j := left, right

	for i <= j {
		for ; aver > array[i]; i++ {
		}
		for ; aver < array[j]; j-- {
		}
		if i <= j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
	}

	quickSortMain(array, left, j)
	quickSortMain(array, i, right)
}

func binSearch(array []int, elem int) (int, error) {
	return binSearchMain(array, elem, 0, len(array)-1)
}

func binSearchMain(array []int, elem int, low int, up int) (int, error) {
	if up < low {
		return -1, errors.New("Элемент не найден в массиве")
	}
	aver := (low + up) / 2

	if array[aver] == elem {
		return aver, nil
	} else if array[aver] > elem {
		return binSearchMain(array, elem, low, aver-1)
	} else {
		return binSearchMain(array, elem, aver+1, up)
	}
	
}
