package main

import "fmt"

func main() {
	fmt.Println("map operatons")

	m := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	fmt.Println(m)

	fmt.Println("====add====")
	m["d"] = 40
	fmt.Println(m)

	fmt.Println("====delete====")
	delete(m, "b")
	fmt.Println(m)

	fmt.Println("====for====")
	for k, v := range m {
		fmt.Printf("key: %s value: %d\n", k, v)
	}

}
