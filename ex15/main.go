package main

import "fmt"

/*К каким негативным последствиям может привести
данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // 1024
	justString = v[:100]
	fmt.Println(justString)
}

func main() {
	someFunc()
}

func createHugeString(leng int) string {
	byteArray := make([]byte, leng)
	for i := range byteArray {
		byteArray[i] = 'A' // Например, заполним символами 'A'
	}
	return string(byteArray)

}
