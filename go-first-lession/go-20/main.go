package main

func readyExtBySwitch(ext string) {
	switch ext {
	case "json":
		println("read json file")
	case "jpg", "jpeg", "png", "gif":
		println("read image file")
	case "txt", "md":
		println("read text file")
	case "yml", "yaml":
		println("read yaml file")
	default:
		println("unsupported file extension:", ext)
	}
}

func typeSwitch() {
	var x interface{} = 123
	// x.(type) is special for switch, and x should be a interface
	switch v := x.(type) {
	case nil:
		println("x is nil")
	case int:
		println("the type of x is int, v = ", v)
	case string:
		println("the type of x is string, v = ", v)
	case bool:
		println("the type of x is boolean, v = ", v)
	default:
		println("do not support the type, v = ", v)
	}
}

func switchAndForTip_with_issue() {
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1

	// Go 语言规范中明确规定，不带 label 的 break 语句中断执行并跳出的，
	// 是同一函数内 break 语句所在的最内层的 for、switch 或 select。
	for i := 0; i < len(sl); i++ {
		switch sl[i] % 2 {
		case 0:
			firstEven = sl[i]
			break
		case 1:
			// do nothing
		}
	}
	println(firstEven)
}

func fixed_switchAndFor() {
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1

loop:
	for i := 0; i < len(sl); i++ {
		switch sl[i] % 2 {
		case 0:
			firstEven = sl[i]
			break loop
		case 1:
			// do nothing
		}
	}
	println(firstEven)
}

func main() {
	readyExtBySwitch("jpg")
	typeSwitch()
	switchAndForTip_with_issue()
	fixed_switchAndFor()
}
