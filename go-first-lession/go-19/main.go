package main

import (
	"fmt"
	"time"
)

func firstFor() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)
}

func noDeclareFor() {
	i := 0
	var sum int
	for ; i < 10; i++ {
		sum += i
	}
	println(sum)
}

func noFrontEndFor() {
	i := 0
	var sum int
	for i < 10 {
		sum += i
		i++
	}
	println(sum)
}

func infiniteLoop() {
	for {

	}

	// for ;;{
	// }

	// for true {
	// }
}

func forRange() {
	var sl = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sl); i++ {
		println("for range ", sl[i])
	}

	for i, v := range sl {
		fmt.Printf("for range index %d is %d \n", i, v)
	}

	for i := range sl {
		println("index ele is ", sl[i])
	}

	for _, v := range sl {
		println("ele is ", v)
	}

	var index int
	for range sl {
		println("sl index is ", index)
		index++
	}
}

func forRangeString() {
	var s = "中国人"
	for i, v := range s {
		fmt.Printf("rune index is %d, rune value is %c, rune ox value is 0x%x \n", i, v, v)
	}
}

// the only way to traversing a map
func forMap() {
	var m = map[string]int{
		"Rob":  1,
		"Rose": 2,
		"John": 3,
	}
	for k, v := range m {
		fmt.Printf("The key is %s, value is %d\n", k, v)
	}
}

func labelContinue() {
	var sl = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

outer:
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			fmt.Printf("The sl[%d][%d] value is %d\n", i, j, sl[i][j])
			if sl[i][j] == 13 {
				fmt.Printf("Found 13 at [%d, %d], and will end this loop\n", i, j)
				continue outer
			}
		}
	}
}

func variableReuse_1() {
	var m = []int{1, 2, 3, 4, 5}

	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 10)
}

func variableReuse_2() {
	var m = []int{1, 2, 3, 4, 5}
	{
		i, v := 0, 0
		for i, v = range m {
			go func() {
				time.Sleep(time.Second * 3)
				fmt.Println(i, v)
			}()
		}
	}
	time.Sleep(time.Second * 10)
}

func variableReuse_3() {
	var m = []int{1, 2, 3, 4, 5}
	{
		i, v := 0, 0
		for i, v = range m {
			go func(i, v int) {
				time.Sleep(time.Second * 3)
				fmt.Println(i, v)
			}(i, v)
		}
	}
	time.Sleep(time.Second * 10)
}

func main() {
	firstFor()
	noDeclareFor()
	noFrontEndFor()
	forRange()
	forRangeString()
	forMap()
	labelContinue()
	variableReuse_1()
	variableReuse_2()
	variableReuse_3()
}
