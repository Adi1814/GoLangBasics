package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices in golang") //used alot in this lang
	var fruitList = []string{"apple", "banana", "peach"}
	fmt.Printf("type of fruitlist %T\n", fruitList)

	fruitList = append(fruitList, "mango", "kiwi")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) //before the colon starts, after ends and if not specified take default(0,to end[where this is exclusive])
	fmt.Println(fruitList)

	highscores := make([]int, 4)
	highscores[0] = 345
	highscores[1] = 234
	highscores[2] = 432
	highscores[3] = 111

	highscores = append(highscores, 555, 666)

	fmt.Println(highscores)

	sort.Ints(highscores)

	//how to remove a value from slices based n index

	var animals = []string{"cat", "cat1", "cat2", "cat3", "cat4"}
	fmt.Println(animals)
	var index int = 2
	animals = append(animals[:index], animals[index+1:]...)
	fmt.Println(animals)

}
