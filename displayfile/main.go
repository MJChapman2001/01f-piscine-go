package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else {
		file, err := os.Open(args[0])

		if err != nil {
			fmt.Println("File not found")
		} else {
			var content []byte

			file.Read(content)

			fmt.Println(string(content))
		}
	}
}
