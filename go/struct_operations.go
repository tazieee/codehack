package main

import "fmt"

func main() {
	fmt.Println("=====struct operations=====")

	a := new(A)
	a.a = "aaabbbb"
	a.b = true

	fmt.Println(a)

	b := &A{
		a: "cccdddd",
		b: false,
	}

	fmt.Println(b)

	c := A{
		a: "eeeffff",
		b: true,
	}

	fmt.Println(c)
}

type A struct {
	a string
	b bool
}