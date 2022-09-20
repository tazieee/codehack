package main

import "fmt"

func main() {
	fmt.Println("slice operations")
	input1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("input1:")
	fmt.Println(input1)
	input2 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Println("input2:")
	fmt.Println(input2)

	// operations
	a := input1[3:6]
	fmt.Println(a)
	a[0] = 123
	fmt.Println(a)
	fmt.Println(input1)
}
