package main

import (
	"fmt"
)

/*
	https://www.jianshu.com/p/dbd4e6b4900c
*/

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(song string) {
	fmt.Println("La la, la la la, la la la la la,la..." + song)
}

func (h Human) Guzzle(beesrtein string) {
	fmt.Printf("Guzzle Guzzle Guzzle", beesrtein)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, call me on %s\n", e.name,
		e.company, e.phone)
}

func (s Student) BorrowMoney(amount float32) {
	s.loan += amount
	fmt.Printf("I borrow money %d, now I have %d", amount, s.loan)
}

func (e Employee) SpendSalary(amount float32) {
	e.money -= amount
	fmt.Printf("Hi, I am %s, I spend salary %d, now i have money %d", e.name, amount, e.money)
}

type Men interface {
	SayHi()
	Sing(song string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendMoney(amount float32)
}

/*
Interface values

因为接口也是一种类型，你会困惑于一个接口类型的值到底是什么。

有一个好消息就是：如果你声明了一个接口变量，这个变量能够存储任何实现该接口的对象类型。

也就是说，如果我们声明了Men类型的接口变量m，那么这个变量就可以存储Student和Employee类型的对象，还有Human类型（差点忘掉）。这是因为他们都实现了Men接口声明的方法签名。

如果m能够存储不同数据类型的值，我们可以轻松实现一个Men切片，该切片包含不同的数据类型的实例。

作者：Zuozuohao
链接：https://www.jianshu.com/p/dbd4e6b4900c
*/

/*
//dev-002
你可能已经注意到，接口类型是一组抽象的方法集，他本身并不实现方法或者精确描述数据结构和方法的实现方式。接口类型只是说：“兄弟，我实现了这些方法，我能胜任”。

值得注意的是这些数据类型没有提及任何的关于接口的信息（我的理解是Student和Employee数据类型），方法签名的实现部分也没有包含给定的接口类型的信息。

同样的，一个接口类型也不会去关心到底是什么数据类型实现了他自身，看看Men接口没有涉及Student和Employee的信息就明白了。接口类型的本质就是如果一个数据类型实现了自身的方法集，那么该接口类型变量就能够引用该数据类型的值。

The case of the empty interface

空接口类型interface{}一个方法签名也不包含，所以所有的数据类型都实现了该方法。

空接口类型在描述一个对象实例的行为上力不从心，但是当我们需要存储任意数据类型的实例的时候，空接口类型的使用使得我们得心应手。

作者：Zuozuohao
链接：https://www.jianshu.com/p/dbd4e6b4900c

如果一个函数的参数包括空接口类型interface{}，实际上函数是在说“兄弟，我接受任何数据”。如果一个函数返回一个空接口类型，那么函数再说“我也不确定返回什么，你只要知道我一定返回一个值就好了”。
*/

func main() {
	mike := Student{Human{"Mike", 25, "222-222-222"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-111-111"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-444-444"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 36, "555-555-555"}, "Things Inc.", 5000}

	//a variable of the interface type Men
	var i Men

	//i can store a Student
	i = mike
	fmt.Printf("This is Mike, a Student: ")
	i.SayHi()
	i.Sing("November rain")

	// i can store a Employee too
	i = tom
	fmt.Printf("This is Tome, a Empoyee: ")
	i.SayHi()
	i.Sing("Born to be wind")

	//a slice of Men
	fmt.Println("Let's use a slice of Men and see what happens: ")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

	//-->//dev-002
	//a is a empty interface type
	var a interface{}
	var i2 int = 5
	s := "Hello World"

	//there are legal statement
	a = i2
	fmt.Println(a)

	a = s
	fmt.Println(a)
}
