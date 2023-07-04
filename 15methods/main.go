package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in golang!")
	//no inheritance in golang no parent no super woohoo!

	Aditi := User{"aditi", "aditicat@tekion.com", false, 23}
	fmt.Println(Aditi)
	fmt.Printf("aditi detials are %+v\n", Aditi)
	fmt.Printf("name is %v, email is %v\n", Aditi.Name, Aditi.Email)
	Aditi.getStatus()
	Aditi.NewEmail()
	fmt.Printf("name is %v, email is %v\n", Aditi.Name, Aditi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("is User Active?", u.Status)
}

//lets try manipulating email

func (u User) NewEmail() {
	u.Email = "newemail@tekion.com"
	fmt.Println("new email is:", u.Email)
} //hawww it didnt get manipulated cuz whn you pass an obkect it passes a copy of it and not object exactly hence pointers come into role
