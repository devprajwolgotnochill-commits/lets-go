package main

import "fmt"

func main() {
	fmt.Println("ARRYS IN THE GO LANG")

	//
	var arry_num [12]int

	for i := 0; i < len(arry_num); i++ {

		arry_num[i] = i + 1
	}

	fmt.Println(arry_num)

	// Array with a fixed size of 3
	//doesnt calculate the values
	// even with 2 elements only len = 3
	numbers := [3]int{10, 20}

	// Array where the compiler counts the elements
	colors := [...]string{"red", "green", "blue", "yellow"}

	fmt.Println(numbers, "Lenth = ", len(numbers))

	fmt.Println(colors, "Lenth = ", len(colors))

}
