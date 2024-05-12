package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	result := intersection(slice)
	fmt.Println(result)

}

func intersection(slice []string) []string {
	result := make([]string, 0)
	set := make(map[string]struct{})

	for _, num := range slice {
		_, exists := set[num]
		if !exists {
			set[num] = struct{}{}
		}
	}

	for num := range set {
		result = append(result, num)
	}
	return result

}
