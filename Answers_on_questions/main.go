// Go program to illustrate how to concatenate strings
// Using strings.Builder with WriteString() function

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Creating and initializing strings
	// Using strings.Builder with
	// WriteString() function
	var str strings.Builder

	str.WriteString("Welcome")

	fmt.Println("String: ", str.String())

	str.WriteString(" to")
	str.WriteString(" GeeksforGeeks!")

	fmt.Println("String: ", str.String())

}
