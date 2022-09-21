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

	// len and cap
	// len <= cap
	s1 := make([]int, 5)
	fmt.Printf("len: %d cap: %d value: %v\n", len(s1), cap(s1), s1)

	s2 := make([]int, 0, 5)
	fmt.Printf("len: %d cap: %d value: %v\n", len(s2), cap(s2), s2)

	s3 := append(s1, 7)
	fmt.Printf("len: %d cap: %d value: %v\n", len(s3), cap(s3), s3)

	s4 := append(s1, s2...)
	fmt.Printf("len: %d cap: %d value: %v\n", len(s4), cap(s4), s4)
}
