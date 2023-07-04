package main

import "fmt"

const LoggedInToken string = "meowmeowmeow" //so in golang when you declare a variable name in capital meaning the first letter is capital that means its public(in other languages we use the public keyword) so here LoggedInToken is accessible by any other file into this folder or in this program.

func main() {
	var username string = "Aditi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallfloat float32 = 255.87468756386
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	//default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherstring string
	fmt.Println(anotherstring)
	fmt.Printf("Variable is of type: %T \n", anotherstring)

	//implicit type
	var website = "aditigolang.com"
	fmt.Println(website)

	//no var style
	cats := "beauty"
	fmt.Println(cats)

	fmt.Println(LoggedInToken)
	fmt.Printf("Variable is of type: %T \n", LoggedInToken)
}
