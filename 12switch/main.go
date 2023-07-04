package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch cases in go!")
	rand.Seed(time.Now().Unix())
	dicenum := rand.Intn(6) + 1
	fmt.Println("value of dice number is :", dicenum)

	switch dicenum {
	case 1:
		fmt.Println("value is 1 and you can open!")
	case 2:
		fmt.Println("value is 2 and you can open!")
	case 3:
		fmt.Println("value is 3 and you can open!")
		fallthrough
	case 4:
		fmt.Println("value is 4 and you can open!")
		fallthrough
	case 5:
		fmt.Println("value is 5 and you can open!")
	case 6:
		fmt.Println("value is 6 and you can open!")
	default:
		fmt.Println("sed")
	}

}
