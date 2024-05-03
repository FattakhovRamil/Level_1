package main

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.
import "fmt"

func main() {
	a := 10
	b := 20

	a, b = b, a

	fmt.Println(a, b)

}
