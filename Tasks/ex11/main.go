package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}

	result := intersection(slice1, slice2)
	fmt.Println(result)

}

func intersection(slice1 []int, slice2 []int) []int {
	result := make([]int, 0)
	set := make(map[int]struct{})

	for _, num := range slice2 {
		set[num] = struct{}{}
	}

	for _, num := range slice1 {

		_, exists := set[num]
		if exists {
			result = append(result, num)
		}
	}
	return result

}
