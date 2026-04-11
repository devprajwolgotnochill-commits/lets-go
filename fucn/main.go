package main

import "fmt"

// main is the entry point for the program

func main() {

	hello_world()

	fmt.Println("GoLang")

	// anonymous  funcutions
	// normal function are not allowed inside of a function ,nested func
	// anonumous func are allowed!
	func() {
		fmt.Println("Welcome! to anonymous  funcutions")
	}()

}

func hello_world() {
	fmt.Println("Hello ,World")

}
