package main

import "fmt"

func main() {

	// null pointer
	// * used to define
	// pointer are 8 bytes
	// // how to define a pointer
	// var pointer *int

	// var
	i := 42

	// & pointer adress * pointer declaration
	// addr of i in memory is passed to p
	// p is a pointer
	// referance

	// What is p?

	// p is just a normal variable
	// —but instead of storing a number like 42, it stores a memory address.

	//  p holds the address of i
	//  *p accesses the value stored at that address
	p := &i

	//1st line prints the address of the var i in memory
	// 2nd line prints the actual value of the pointer
	// *p accese the value stored at the pointer
	fmt.Println(p, *p)

	// what is this
	// change the value AT the address stored in p
	*p = 2007

	// its 2007
	fmt.Println(*p)

	// its also 2007
	fmt.Println(i)

	// swap with the fuck
	// i dont need pointer ???
	x := 12
	y := 24

	fmt.Println(x, y)

	// needs
	// provides the addr
	ptr1 := &x
	ptr2 := &y

	swap(ptr1, ptr2)

	fmt.Println(x, y)

}
