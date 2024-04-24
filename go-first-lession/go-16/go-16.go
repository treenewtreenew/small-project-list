// The map value type in go is no constraints,
// but the key type is constrained within type with equal and not equal operation
package main

import "fmt"

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

func main() {
	initMap()
}
