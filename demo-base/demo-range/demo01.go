package main

import (
	"fmt"
)

func main() {
	l := []string{"lili", "xiaoming", "xiaoma"}
	for k, v := range l {
		fmt.Printf("pos: %d, value: %s\n", k, v)
	}
}
