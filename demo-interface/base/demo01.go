package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
	Id   int
	Name string
}

func (NokiaPhone NokiaPhone) call() {
	fmt.Printf("I'm NokiaPhone, I can call you!\n")
}

type IPhone struct {
	Id   int
	Name string
}

func (IPhone IPhone) call() {
	fmt.Printf("I'm IPhone, I can call you!\n")
}

func main() {
	var phone Phone

	/*
		NokiaPhone.Id = 1
		NokiaPhone.Name = "NokiaPhone01"

		IPhone.Id = 101
		IPhone.Name = "IPhone01"
	*/

	//NokiaPhone1 := NokiaPhone{1, "test"}
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
