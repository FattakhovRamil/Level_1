package main

import (
	"fmt"
)

/*
	Реализовать паттерн «адаптер» на любом примере.
*/

type NewFunc interface { // новый код
	NewPrintInfo()
}

type LegacyFunc struct{} // старый код

func (lf *LegacyFunc) OldPrint() string {
	return "Старый код"
}

type AdapterLegacyFunc struct {
	LegacyFunc *LegacyFunc
}

func (ad *AdapterLegacyFunc) NewPrintInfo() string {
	return ad.LegacyFunc.OldPrint()
}

func main() {
	old := &LegacyFunc{}
	adapter := &AdapterLegacyFunc{
		LegacyFunc: old,
	}
	message := adapter.NewPrintInfo()
	fmt.Println(message)
}
