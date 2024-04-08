package pkg2

import (
	"fmt"
	_ "github.com/bigwhite/prog-init-order/pkg3"
)​

func Main() {
	main()
}​

func main() {
	fmt.Println("main func for pkg2")
}