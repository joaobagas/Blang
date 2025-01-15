package main

import (
	"Blang/src"
	"fmt"
)

func main() {
	fmt.Println("Write the path to the file you want to compile:")
	var path string
	_, err := fmt.Scanln(&path)

	if err != nil {
		panic(err)
	}

	src.Compile(path)
}
