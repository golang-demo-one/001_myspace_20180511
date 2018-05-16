package main

import (
	"fmt"
)

func main() {
	var p *int //p=nil
	//*p = 8
	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal 0xc0000005 code=0x1 addr=0x0 pc=0x401068]

		goroutine 1 [running]:
		main.main()
		C:/001_myspace_20180511/demo-base/demo-pointer/demo01.go:9 +0x28
	*/

	var i int
	p = &i //令p的值等于i的内存的值
	*p = 8 //相当于修改了i的值

	//fmt.Printf(p, i, &i, *p)
	/*
		.\demo01.go:14: cannot use p (type *int) as type string in argument to fmt.Printf
	*/
	fmt.Printf("%d, %d, %d, %d", p, i, &i, *p)
}
