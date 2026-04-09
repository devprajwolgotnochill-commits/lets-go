// how to convert the string to a number so i can perform some arthmetici operation
// what are arthimatic operations

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// i dont need to import the function user made modules // i just use it

func typeConversion(given_string string) int {

	//bringing User_input for main.go //use go run ./
	// need to build both file inorder to work
	// fmt.Println("You entered ", given_string)

	// trims the \n part or every \n
	str := strings.TrimSpace(given_string)

	str_to_num, error_in_conversion := strconv.ParseFloat(str, 64)

	if error_in_conversion != nil {

		fmt.Println(error_in_conversion)
		fmt.Println("Err enter a number!!!")

	} else {
		return int(str_to_num)
	}
	return -1
}
