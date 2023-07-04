package main

import "fmt"

func main() {
	fmt.Println("arrays in golang") //small bummer not used much in this language

	var fruitList = [4]string{}
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "tomato"

	fmt.Println("The fruitlist:", fruitList)

	var movieList = [3]string{"Spiderman-into the spiderverse", "Spiderman-across the spiderverse", "Spiderman-Beyond the spiderverse"}

	fmt.Println(".........")
	fmt.Println("The movie List:", movieList)
}
