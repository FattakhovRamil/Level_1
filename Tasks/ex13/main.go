package main

import "fmt"

func main() {
	a := 10
	b := 20

	a, b = b, a

	fmt.Println(a, b)

}



