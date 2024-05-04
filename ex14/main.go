package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func main() {
	var intt int = 20
	var stringg string = "some text"
	var boolean bool = false
	ch := make(chan int)

	fmt.Println(getType(intt))
	fmt.Println(getType(stringg))
	fmt.Println(getType(boolean))
	fmt.Println(getType(ch))
	fmt.Println(getType(1.5))
}

func getType(v interface{}) string {
	switch v.(type) {
	case int:

		return "int"

	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "channel"
	default:
		return "unknown"
	}

}
