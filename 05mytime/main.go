package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Handling time in golang!")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(".....")
	var formatt string = "01-02-2006 15:04:05 Monday"
	fmt.Println(presentTime.Format(formatt))
	fmt.Println(".....")
	createdDate := time.Date(2007, time.March, 4, 17, 45, 59, 23, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(".....")
	fmt.Println(createdDate.Format(formatt))
	fmt.Println(".....")
	const spookyForm = "2006 ooga booga January 02"
	t, _ := time.Parse(spookyForm, "2020 ooga booga October 31")
	fmt.Println(t.Format(spookyForm)) //reverse cool cool thingggg

	/*n Go layout string is:
	Jan 2 15:04:05 2006 MST
	1   2  3  4  5    6  -7
	cray cray for go developers to think of this!!!!!!
	*/
}
