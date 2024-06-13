package main

import (
	"fmt"
	"reflect"
)

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2(v int) {
	t.a = v
}

// 如果 Go 方法要把对 receiver 参数代表的类型实例的修改，
// 反映到原类型实例上，那么我们应该选择 *T 作为 receiver 参数的类型。

// 如果 receiver 参数类型的 size 较大，以值拷贝形式传入
// 就会导致较大的性能开销，这时我们选择 *T 作为 receiver 类型可能更好些。

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set: \n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

type T_ struct{}

func (T_) M1() {}
func (T_) M2() {}

func (*T_) M3() {}
func (*T_) M4() {}

func main() {
	var t T
	println(t.a)

	t.M1()
	println(t.a)

	t.M2(12)
	println(t.a)

	p := &t
	p.M2(11)
	println(t.a)

	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t_ T_
	dumpMethodSet(t_)
	dumpMethodSet(&t_)
}

// 这个原则的选择依据就是 T 类型是否需要实现某个接口，也就是是否存在将 T 类型的变量
// 赋值给某接口类型变量的情况。如果 T 类型需要实现某个接口，那我们就要使用 T 作为
// receiver 参数的类型，来满足接口类型方法集合中的所有方法。如果 T 不需要实现某一接口，
// 但 *T 需要实现该接口，那么根据方法集合概念，*T 的方法集合是包含 T 的方法集合的，
// 这样我们在确定 Go 方法的 receiver 的类型时，参考原则一和原则二就可以了。
