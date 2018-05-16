package main

import (
	"fmt"
)

func getSquence() func() int {
	i := 1
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSquence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSquence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
