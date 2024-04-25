// The map value type in go is no constraints,
// but the key type is constrained within type with equal and not equal operation
package main

import (
	"fmt"
	"time"
)

func compareMap() {
	/*
		s1 := make([]int, 1)
		s2 := make([]int, 2)
		f1 := func() {}
		f2 := func() {}
		m1 := make(map[int]string)
		m2 := make(map[int]string)
		println(s1 == s2) // invalid operation: s1 == s2 (slice can only be compared to nil)
		println(f1 == f2) // invalid operation: f1 == f2 (func can only be compared to nil)
		println(m1 == m2) // invalid operation: m1 == m2 (map can only be compared to nil)
	*/
}

func initMap() {
	var m1 map[string]int
	// m1["key"] = 1  // map unable to zero value available
	fmt.Println(m1)

	// literal init
	m2 := map[int]string{}
	m2[1] = "can be set"
	fmt.Println(m2)

	// complicated literal init
	m3 := map[int][]string{
		1: []string{"val1_1", "val1_2"},
		3: []string{"val3_1", "val3_2", "val3_3"},
		7: []string{"val7_1"},
	}
	fmt.Println(m3)

	type Position struct {
		x float64
		y float64
	}
	m4 := map[Position]string{
		Position{29.935523, 52.568915}:  "school",
		Position{25.352594, 113.304361}: "shopping-mall",
		Position{73.224455, 111.804306}: "hospital",
	}
	fmt.Println(m4)

	// syntactic sugar
	m5 := map[int][]string{
		1: {"val1_1", "val1_2"},
		3: {"val3_1", "val3_2", "val3_3"},
		7: {"val7_1"},
	}
	fmt.Println(m5)

	m6 := map[Position]string{
		{29.935523, 52.568915}:  "school",
		{25.352594, 113.304361}: "shopping-mall",
		{73.224455, 111.804306}: "hospital",
	}
	fmt.Println(m6)

	// init using make

}

func basicInsert() {
	m := make(map[int]string)
	m[1] = "v1"
	m[2] = "v2"
	m[3] = "v3"
	m[3] = "v3_"
	fmt.Println(m)
}

func getOp() {
	m := map[string]int{
		"v1": 1,
		"v2": 2,
	}
	fmt.Println(len(m))
	fmt.Println(m["v1"])
	fmt.Println(m["v3"])
	fmt.Println(len(m))
	m["k3"] = 3
	fmt.Println(len(m))
}

func searchAndRead() {
	// not recommend
	m := make(map[string]int)
	v := m["key1"]
	fmt.Println(v)

	// comma ok
	v1, ok := m["key1"]
	if !ok {
		fmt.Println("key1 does not exist")
		fmt.Println(v1)
	}
}

func delete1() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m["key1"])
	delete(m, "key1")
	fmt.Println(len(m))
}

func iterateAllEle() {
	m := map[string]int{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
		"k5": 5,
	}
	fmt.Printf("{ ")
	for k, v := range m {
		fmt.Printf("[%s, %d] ", k, v)
	}
	fmt.Printf("}\n")

	// only key
	for k := range m {
		fmt.Printf("[%s] ", k)
	}
	// only value
	for _, v := range m {
		fmt.Printf("[%d] ", v)
	}
}

func sideEffect(m map[string]int) {
	m["key1"] = 2
	m["key2"] = 3
}

func mapInternal() {
	// when map key or value is bigger than certain threshold, the go
	// runtime do not store their value in bucket and store their pointer
	// $GOROOT/src/runtime/map.go
	const (
		maxKeySize  = 128
		maxElemSize = 128
	)
}

func autoScale() {
	/*
			// $GOROOT/src/runtime/map.go
		const (
		  ... ...

		  loadFactorNum = 13
		  loadFactorDen = 2
		  ... ...
		)

		func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
		  ... ...
		  if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		    hashGrow(t, h)
		    goto again // Growing the table invalidates everything, so try again
		  }
		  ... ...
		}
	*/
}

func doIteration(m map[int]int) {
	for k, v := range m {
		_ = fmt.Sprintf("[%d. %d] ", k, v)
	}
}

func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}

func main() {
	// initMap()
	// basicInsert()
	// getOp()
	// searchAndRead()
	// delete1()
	// iterateAllEle()
	/*m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	sideEffect(m)
	fmt.Println(m["key1"])
	fmt.Println(m["key2"])
	*/

	m := map[int]int{1: 11, 2: 12, 3: 13}

	go func() {
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()
	time.Sleep(5 * time.Second)
	// fatal error: concurrent map iteration and map write

	// invalid operation: cannot take address of m[1] (map index expression of type int)
	// p := &m[1]
}
