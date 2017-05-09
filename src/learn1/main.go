package main

import (
	"fmt"
	"time"
)

func pstring(a string, b string) (string, string) {
	return a, b
}
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	defer fmt.Println("world")
	defer fmt.Println("worlrldd")
	fmt.Println("hello")
	fmt.Println("asdas")

	a := 4
	fmt.Println(a)
	if a := 10; a < 10 {
		fmt.Println(a)
	} else if a >= 10 {
		fmt.Println(a)
	}
	fmt.Println(a)
	fmt.Println("The time is ", time.Now())

	a1, _ := pstring("aaaa", "bbbb")
	fmt.Println(a1)
}
