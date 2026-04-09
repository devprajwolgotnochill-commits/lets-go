package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	// slices alt for arry or better arry
	var slices_number = []int{1, 2, 3, 44}

	fmt.Println(slices_number)

	fmt.Printf("%T \n", slices_number)

	// for i := 0; i < 5; i++ {
	// 	slices_number = append(slices_number, i)
	// }

	// for using range method
	// works with // arrays// slices// strings // maps// channels

	for i := range 10 {
		// starts with a zero
		slices_number = append(slices_number, i)
	}

	fmt.Println(slices_number)

	// slicing slice
	// can use '' single
	// use "" for string in slices
	// var s = []string{"a", "b", "c", "d", "e"}

	// s = append(s[0:3])

	// fmt.Println(s)

	// its slice not arry
	// slice can grow
	students_arry := make([]int, 5)

	for i := range len(students_arry) {
		// generats random number jst 4 fun
		students_arry[i] = i + rand.Intn(677)
	}

	// we can use append to increase the memory size
	// make([]int, 5)
	students_arry = append(students_arry, 222, 333, 444, 555)

	fmt.Println(students_arry)

	// it sorts the arry in increasing orders
	sort.Ints(students_arry)

	fmt.Println(students_arry)

	fmt.Println(sort.IntsAreSorted(students_arry))
}
