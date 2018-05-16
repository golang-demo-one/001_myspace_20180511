//这个示例程序展示如何声明
//并使用方法

package main

import (
	"fmt"
)

//user 在程序里定义了一个用户类型
type user struct {
	name  string
	email string
}

//notify 使用值接受者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

//changeEmail 使用指针接受者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

//main 是应用程序的入口
func main() {
	//user 类型的值可以用来调用
	//使用值接受者使用的方法
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	//指向user类型值的指针也可以用来调用
	//使用值接受者使用的方法
	lisa := &user{"lisa", "lisa@email.com"}
	lisa.notify()

	//user 类型的值可以用来调用
	//使用指针接受者声明的方法
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	//指向user类型值的指针可以用来调用
	//使用指针接受者声明的方法
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
