package main

import "fmt"

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var a int = 100
	var b int = 200
	var ret int

	c, d := swap("John", "Lisa")

	ret = max(a, b)

	fmt.Printf("最大值是：%d\n", ret)
	fmt.Println(c, d)
}
