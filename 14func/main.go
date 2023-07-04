package main

import (
	"fmt"
)

func main() {
	fmt.Println("functions in go!")
	greater()
	result := adder(3, 5)
	fmt.Println("result is :", result)

	proresult := proadder(2, 5, 6, 7, 7, 8)
	fmt.Println("result is :", proresult)
}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}
func proadder(values ...int) int {
	total := 0
	for _, tem := range values {
		total += tem
	}
	// for tem := range values {
	// 	total += tem
	// } major different without the _ why tho?? (figure out in office tomorrow)
	return total
}
func greater() {
	fmt.Println("called the greeter func!")
}
