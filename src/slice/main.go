package main

import (
	"fmt"
)

func main() {
	arr := [10]int{1, 2}
	fmt.Println(arr[0])
	fmt.Println(arr[9])
	fmt.Println("asd")
	var a, b []int
	a = arr[0:4]
	b = arr[:3]
	fmt.Println(a)
	fmt.Println(b)
	arr[0] = 99
	fmt.Println(arr)
	fmt.Println(a)
	fmt.Println(b)
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	for k, v := range numbers {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}

}
