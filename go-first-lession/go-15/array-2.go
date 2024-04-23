package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arrayLengthAndSize()
	arrayInitWith0()
	arrayInitWithLiteral()
	sliceTest()
	crateSL()
	autoExtend1()
}

func arrayLengthAndSize() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("array length is : ", len(arr))
	fmt.Println("array size is : ", unsafe.Sizeof(arr))
}

func arrayInitWith0() {
	var arr [5]int
	v := 0
	i := 0
	for i, v = range arr {
		fmt.Println("array element ", i, v)
	}
}

func arrayInitWithLiteral() {
	var arr = [6]int{11, 12, 13, 14, 15, 16}
	var arr2 = [...]int{1, 2, 3, 4}
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", arr2)

	var arr3 = [10]int{5: 5}
	i, v := 0, 0
	for i, v = range arr3 {
		fmt.Println("iterate arr3", i, v)
	}
}

func sliceTest() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(nums))
	nums = append(nums, 7)
	fmt.Println(len(nums))
}

func crateSL() {
	sl := make([]byte, 5, 10)
	fmt.Println(len(sl))
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl2 := arr[3:7:9]
	fmt.Println(len(sl2))
	for i, v := range sl2 {
		fmt.Println(i, v)
	}
	sl2[0] += 10
	fmt.Println(arr[3])
}

func autoExtend1() {
	var s []int
	s = append(s, 11)
	fmt.Println(len(s), cap(s)) //1 1
	s = append(s, 12)
	fmt.Println(len(s), cap(s)) //2 2
	s = append(s, 13)
	fmt.Println(len(s), cap(s)) //3 4
	s = append(s, 14)
	fmt.Println(len(s), cap(s)) //4 4
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) //5 8
}
