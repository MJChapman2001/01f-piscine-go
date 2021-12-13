package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 3 {
		runnable := true
		res := ""

		for i := range args {
			_, err := strconv.Atoi(args[i])
			if err != nil {
				runnable = false
			} else {
				res += args[i]
			}
		}

		if runnable {
			if args[1] == "/" && args[2] == "0" {
				fmt.Println("No division by 0")
			} else if args[1] == "%" && args[2] == "0" {
				fmt.Println("No modulo by 0")
			} else {
				fmt.Println(strconv.Atoi(res))
			}
		}
	}
}
