// bufio

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// bufio buildin buffer
	// os input and output

	// bufio.NewReader reads the os input
	// it reads the cin till \n
	// how to read a full sentence then?
	user_input_reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter X:")

	// .readstring will read till the \n , \n ends the reading i it will read the \n too
	// so x = "1\n"
	// user_input stores the input and the underscore ignores the error
	// we can use user_input ,errorInInput to store both input and error
	// user_input, _ := user_input_reader.ReadString('\n')

	// try catch
	// if error !user_input ,if user_input no err (err = nil)
	User_inputX, error_in_user_input := user_input_reader.ReadString('\n')

	if error_in_user_input != nil {
		panic(error_in_user_input)
	}

	// // prints error
	// fmt.Println("ERROR ", err)

	// what type is the input?
	// fmt.Printf("Type %T \n", User_input)

	value_numX := typeConversion(User_inputX)

	// 2nd number same process
	// gets user_input -> typeConversion()
	fmt.Println("Enter Y: ")
	var User_inputY string = ""
	User_inputY, error_in_user_input = user_input_reader.ReadString('\n')

	if error_in_user_input != nil {
		panic(error_in_user_input)
	}

	value_numY := typeConversion(User_inputY)
	//finish typeConversion

	// the sum part

	fmt.Println(sum(value_numX, value_numY))
	// sumfun from another file

}
