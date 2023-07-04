package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang!")

	language := make(map[string]string)

	language["JS"] = "Javascript"
	language["PY"] = "Python"
	language["RB"] = "Ruby"

	fmt.Println(language)
	fmt.Println("js stands for:", language["JS"])

	delete(language, "RB")
	fmt.Println(language)

	//loopssssss
	for _, value := range language {
		fmt.Printf("for key v, value is %v\n", value)
	}

}
