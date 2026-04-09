package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello\n")

	sum := Add(2, 3) // directly call Add
	subtract := subtract(2, 3)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", subtract)

	playground()

}

//  go run .
//  go build .
//  build all the .go files in the dir
//  need go.mod (wtf)
