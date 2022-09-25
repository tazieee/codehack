package main

import "fmt"

func main() {
	fmt.Println("struct operations 2")

	person1 := Human{
		age:  10,
		name: "mike",
	}

	fmt.Println("person1:")
	fmt.Println(person1)
	person1.sayHello()

	fmt.Println("======birth day======")
	person1.birthDay()
	fmt.Println(person1)
}

type Human struct {
	age  int
	name string
}

func (h *Human) birthDay() {
	h.age += 1
}

func (h Human) sayHello() {
	fmt.Printf("Hello, I'm %s.\n", h.name)
}
