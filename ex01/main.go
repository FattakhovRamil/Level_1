package main

import (
	"fmt"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int16
}

func (h *Human) Speak() {

	fmt.Printf("My name is %s %s. I am %d years old\n", h.FirstName, h.LastName, h.Age)

}

type Action struct {
	Human
	Discription string
}

func (a *Action) Say() {
	fmt.Printf("My name is %s %s. I am %d years old and I'm described as %s\n", a.FirstName, a.LastName, a.Age, a.Discription)
}

func main() {
	men := Human {
		FirstName: "FName",
		LastName: "LName",
		Age: 28,
	}

	men.Speak()

	menAct := Action{
		Human: men,
		Discription: "Some text",
	}
	menAct.Say()
}
