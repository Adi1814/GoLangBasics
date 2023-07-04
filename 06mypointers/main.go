package main

import "fmt"

func main() {
	fmt.Println("Pointers in Golang!")

	var ptrr *int
	fmt.Println("value of ptrr", ptrr)
	fmt.Println(".........")
	num := 23
	var ptr = &num
	fmt.Println("value of ptr", ptr)
	fmt.Println("value of pointer", *ptr)
	fmt.Println(".......")

	*ptr = *ptr + 2
	fmt.Println("value of pointer", *ptr)

}
