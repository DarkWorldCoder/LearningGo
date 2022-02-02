package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file in golang")
	content := " This is to be added in file"
	file, err := os.Create("./file.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length is:", length)
	file.Close()
}
