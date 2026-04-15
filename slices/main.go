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

	// fmt.Println(students_arry)

	// it sorts the arry in increasing orders
	sort.Ints(students_arry)

	// fmt.Println(students_arry)

	//optimize way of using if we know the slice size?
	// s_arry := make([]int, 4)
	s_arry := []interface{}{234, 24, 34, "pwad", "pwd"}

	// interface{} -> anydata type I D C
	fmt.Println(s_arry)

	// slicing the slice // this is very imp //
	// s_arry = append(s_arry[1:])
	// fmt.Println(s_arry)
	// // [20 30 40] 1 to all

	// s_arry = append(s_arry[:3])
	// fmt.Println(s_arry)
	// [20 30 40]

	// removing value from the slice form a slice VVI
	index := 2

	// simple trick to remove a value from slice
	s_arry = append(s_arry[:index], s_arry[index+1:]...)
	// ... =>Take all the individual elements out of this slice and pass them as separate arguments.
	// s_arry[:1] is 0 to 1 , 1 is not inclucive so only 0 == 234
	// s_arry[1:] is from 1 to end == 24 pwad pwds

	fmt.Println(s_arry)
}
