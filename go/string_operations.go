package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fmt.Println("#######input start#######")
	inputFile1, _ := ioutil.ReadFile("./input/input1.txt")
	input1 := string(inputFile1)
	fmt.Println("input1:")
	fmt.Println(string(input1))

	inputFile2, _ := ioutil.ReadFile("./input/input2.txt")
	input2 := string(inputFile2)
	fmt.Println("input2:")
	fmt.Println(string(input2))
	fmt.Println("#######input end#######")
	fmt.Println()

	fmt.Println("========upper========")
	upper(input1)
	fmt.Println()

	fmt.Println("========lower========")
	lower(input2)
	fmt.Println()

	fmt.Println("========reverse========")
	reverse(input1)
	fmt.Println()
}

func upper(str string) {
	fmt.Println(strings.ToUpper(str))
}

func lower(str string) {
	fmt.Println(strings.ToLower(str))
}

func reverse(str string) {
	slice := strings.Split(str, "")
	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	fmt.Println(strings.Join(slice, ""))
}
