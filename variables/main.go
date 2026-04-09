package main

import "fmt"

func main() {
	var Name string = "prajwal"

	fmt.Println("Your name is " + Name)
	fmt.Printf("This is type: %T \n", Name) //you need \n to make %t work

	// capital T represent data types and small t represent boolean wtf

	var isOnline bool = true
	fmt.Printf("Is online? %t\n", isOnline) // %t is for boolean prints true

	fmt.Printf("Is online? %T\n", isOnline) //prints boolean

}
