package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("Please rate our pizza!")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	// numstring, err := strconv.ParseFloat(input, 64), this wont work cuz input is reading \n too so when you do any arithemetic operation in that it will give you error

	numstring, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Added 1 to your rating", numstring+1)
	}
}
