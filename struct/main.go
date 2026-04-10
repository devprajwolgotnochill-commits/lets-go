package main

import "fmt"

// no oop in go
// capital name = public
type User_name struct {
	Name      string
	letter1st rune //rune (int32)
	Email     string
	Status    bool
	Age       int
}

// rune = char -> single ''
// rune is A single Unicode character stored as an int32
// letter1st = 'a'
//
// cant use single '' for string  ,string need to use ""

func main() {

	user_1st := User_name{"jonny", 'j', "jonny123@gmail.com", true, 23}

	fmt.Println(user_1st)

	fmt.Printf("Your name is %v and you are %v years old.", user_1st.Name, user_1st.Age)

	// fmt.Printf("Details : %+v \n", user_1st)
	// fmt.Printf("Details : %#v \n", user_1st)

}
