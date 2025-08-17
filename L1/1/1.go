package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Introduce() {
	fmt.Printf("My name is %s, I am %d years old.\n", h.Name, h.Age)
}

type Action struct {
	Human // Встраивание
}

func (a *Action) Run() {
	fmt.Printf("%s is running.\n", a.Name)
}

func main() {
	action := Action{Human{Name: "Danya", Age: 18}}

	action.Introduce() // Метод из Human
	action.Run()       // Метод из Action
}
