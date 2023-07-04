package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //just reading as the name suggests
	fmt.Println("Enter the rating for our pizza!")

	//whatever reader reads i want to store that in a varibale thats where comma ok or err ok comes into role
	// so in golang there is no concept of try and catch cuz the pardigm expects you to treat problems or errors like variables

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)
	fmt.Printf("type of this rating is: %T", input)

}
