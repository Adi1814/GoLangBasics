package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ifelse in go!")

	if num := 8; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}
}
