package main

import (
	"fmt"
)

// main is the entry point for the program

func main() {

	// anonymous  funcutions
	// normal function are not allowed inside of a function ,nested func
	// anonumous func are allowed!
	func() {
		fmt.Println("Welcome! to anonymous  funcutions")
	}()

	hello_world()

	numebers_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// adder(1, 2, 3, 4, 5)

	// converts slice → multiple arguments
	// ...
	fmt.Println(adder(numebers_slice...))

	// nameage("wow", 21)

}

func hello_world() {
	fmt.Println("Hello ,World")

}

// ariadic parameter
// any number of int arguments
// values becomes a slice ([]int)
// ...

func adder(values ...int) int {

	total := 0

	// _ stands for the index of the slice
	// i stands for the value at index, stores the value
	// slice[i]

	for _, i := range values {
		total += i
		// fmt.Println(total)
	}
	return total

}

// to return two value ?

// func nameage(name string, age int) (string, int) {

// 	// fmt.Sprintf returns a string (instead of printing)
// 	// return fmt.Sprintf("Your name is %s and age is %d", name, age)

// 	return name, age
// }

// lamda functions in go
// A function without a name, often used inline or assigned to a variable.
// // 	result := func(a int, b int) int {
//      	   return a * b
//     	}(4, 5)
// //
