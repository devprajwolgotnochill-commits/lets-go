//filehandiling

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// fmt.Println("Hello wol")

	msg := "MY NAME IS SILA SILA KI JAWANI \n"

	file, er := os.Create("./test.txt")

	// // type os.File ?
	// fmt.Printf("Type: %T\n", file)

	if er != nil {
		panic(er)
	}

	// close the file at last
	defer file.Close()

	// wow
	length, er1 := io.WriteString(file, msg)

	if er1 != nil {
		panic(er1)
	}

	fmt.Println("Wrote file of length :", length)

	readfile("./test.txt")

}

func readfile(file_path string) {

	content, er2 := os.ReadFile(file_path)

	if er2 != nil {
		panic(er2)
	}

	fmt.Println("Crazy :", content)
	fmt.Println(string(content))
}
