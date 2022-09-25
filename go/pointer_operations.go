package main

import "fmt"

func main() {
	fmt.Println("pointer operations")

	value1 := 4
	p1 := &value1
	fmt.Printf("value1 value: %d address: %p\n", value1, p1)
	fmt.Println("double")
	double(&value1)
	fmt.Printf("value1 value: %d address: %p\n", value1, p1)

	value2 := 7
	p2 := &value2
	fmt.Printf("value2 value: %d address: %p\n", value2, p2)

	fmt.Printf("p1 : %d p2 : %d\n", *p1, *p2)
	fmt.Println("1 to 2 change")
	temp := p1
	p1 = p2
	p2 = temp
	fmt.Printf("p1 : %d p2 : %d\n", *p1, *p2)

	fmt.Println("===========")
	fmt.Printf("value1 value: %d address: %p\n", value1, p1)
	fmt.Printf("value2 value: %d address: %p\n", value2, p2)

}

func double(v *int) {
	*v = *v * 2
}
