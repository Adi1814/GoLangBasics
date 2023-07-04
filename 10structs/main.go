package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in golang!")
	//no inheritance in golang no parent no super woohoo!

	Aditi := User{"aditi", "aditicat@tekion.com", true, 23}
	fmt.Println(Aditi)
	fmt.Printf("aditi detials are %+v\n", Aditi)
	fmt.Printf("name is %v, email is %v\n", Aditi.Name, Aditi.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type meow string

const (
	meeow meow = "meowwwwwww"
)
