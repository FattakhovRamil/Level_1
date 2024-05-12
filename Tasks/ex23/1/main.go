package main

import "fmt"

/*
	Удалить i-ый элемент из слайса.
*/

func main() {
	slices1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slicesDel := removeElement(slices1, 1)
	fmt.Println(slices1)
	fmt.Println(slicesDel)
}

func removeElement(sli []int, index int) []int {
	if index < 0 || index >= len(sli) {
		fmt.Println("Индекс за пределами дипазона слайса")
		return sli
	}
	result := make([]int, 0, len(sli)-1)

	for ind, val := range sli {
		if ind == index {
			continue
		}
		result = append(result, val)
	}
	return result
}
