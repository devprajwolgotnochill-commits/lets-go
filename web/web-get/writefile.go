package main

import (
	"fmt"
	"io"
	"os"
)

func WriteData(write_path string, write_content string) {

	file, er1 := os.Create(write_path)

	if er1 != nil {
		panic(er1)
	}

	defer file.Close()

	length, er2 := io.WriteString(file, write_content)

	if er2 != nil {
		panic(er2)
	}

	fmt.Println("Wrote file of length :", length)

}
