package main

import (
	"errors"
	"fmt"
)

type QuackableAnimal interface {
	Quack()
}

type Duck struct{}

func (Duck) Quack() {
	println("duck quack!")
}

type Dog struct{}

func (Dog) Quack() {
	println("dog quack!")
}

type Bird struct{}

func (Bird) Quack() {
	println("bird quack!")
}

func AnimalQuackInterfaceTest(a QuackableAnimal) {
	a.Quack()
}

// question form https://go.dev/doc/faq#nil_error
// nil error 值 != nil
type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func main() {
	// 接口的静态特性体现在接口类型变量具有静态类型
	// 编译器会检查右值的类型是否实现了该接口方法集合中的所有方法。如果不满足，就会报错：
	// var err error = 1 // cannot use 1 (type int) as type error in assignment: int does not implement error (missing Error method)

	// 动态特性，接口类型变量在运行时右值的真实类型信息

	animals := []QuackableAnimal{new(Duck), new(Dog), new(Bird)}
	for _, animal := range animals {
		animal.Quack()
	}

	// nil error 值 != nil
	err := returnsError()
	if err != nil {
		fmt.Printf("error occur: %+v\n", err)
		return
	}
	fmt.Println("ok!")

}
