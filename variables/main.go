package main

import "fmt"

//capital 1st letter = public
const PI = 3.15

//this is private
// cant be used in other package should always be used in the same package
const logincredits = "cful14307"

// Capital → "Can be Captured" outside the package
// Lowercase → "Locked inside" the package

// +----------------------+-------------------------+
// | Identifier Name      | Visibility / Access     |
// +----------------------+-------------------------+
// | Capital First Letter | Public / Exported      |
// |                      | Accessible from        |
// |                      | other packages         |
// +----------------------+-------------------------+
// | Lowercase First      | Private / Unexported   |
// | Letter               | Only accessible within |
// |                      | the same package       |
// +----------------------+-------------------------+

func main() {
	var Name string = "prajwal"

	fmt.Println("Your name is " + Name)
	fmt.Printf("This is type: %T \n", Name) //you need \n to make %t work

	// capital T represent data types and small t represent boolean wtf

	var isOnline bool = true
	fmt.Printf("Is online? %t\n", isOnline) // %t is for boolean prints true

	fmt.Printf("Is online? %T\n", isOnline) //prints boolean

	var sum int = 1 + 3

	fmt.Println(sum)
	fmt.Printf("Sum is a %T dataType \n", sum)

	// fmt.Print does not support format verbs like %t or %T.
	// fmt.Printf does support formatting.
	// fmt.Println it prints a new line

	UserName := "email1234@gg"
	// new way of declaring var
	// work only in the local
	//cant be global var
	//declaring it outside of function or globally is not allowed

	fmt.Println("Your username is " + UserName)

}
