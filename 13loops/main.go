package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops in go!")
	var animals = []string{"cat", "cat1", "cat2", "cat3", "cat4"}
	fmt.Println(animals)
	// for d := 0; d < len(animals); d++ {
	// 	fmt.Println(animals[d])
	// }

	// for i := range animals {
	// 	fmt.Println(animals[i])
	// }

	// for index, animal := range animals {
	// 	fmt.Printf("for index %v, animal is %v\n", index, animal)
	// }

	x := 1
	for x < 10 {
		if x == 6 {
			goto lco
			//continue
		} else if x == 4 {
			fmt.Println("meow")
		}
		fmt.Println("value of x is: ", x)
		x++
	}

	//goto
lco:
	fmt.Println("jumping at blah blah ")
}
