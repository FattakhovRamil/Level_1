package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	tepm := []float64{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
	}

	maxI := maxItem(tepm)
	minI := minItem(tepm)
	fmt.Println(minI)
	myMap := make(map[int][]float64)

	for i := minI; i <= maxI; i += 10 {
		taat := make([]float64, 0)
		for _, val := range tepm {
			if val > 0 {
				if val > float64(i) && val < float64(i+10) {
					taat = append(taat, val)
				}
			} else {
				if val < float64(i) && val > float64(i-10) {
					taat = append(taat, val)
				}
			}

		}
		if len(taat) > 0 {
			myMap[i] = taat
		}

	}
	fmt.Println(myMap)
}

func maxItem(tepm []float64) int {
	item := tepm[0]
	for i := 0; i < len(tepm); i++ {
		if item < tepm[i] {
			item = tepm[i]
		}
	}
	return (int(item/10) + 1) * 10
}

func minItem(tepm []float64) int {
	item := tepm[0]
	for i := 0; i < len(tepm); i++ {
		if item > tepm[i] {
			item = tepm[i]
		}
	}
	return (int(item/10) - 1) * 10
}
