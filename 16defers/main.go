package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("hello again")
	defer fmt.Println("hello two")
	defer fmt.Println("hello three") //remember defer follows lifo order!! crazy concept!!
	fmt.Println("Defers in go!")
	mydefer()

}
func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //run the output and see the magic!!! and remember lifo!!
	}
}
