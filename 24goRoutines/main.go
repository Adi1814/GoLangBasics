package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	x := time.Now().UnixMilli()
	var wg sync.WaitGroup
	wg.Add(2)
	go nocat(4, &wg)        // new thread
	go nocat(5, &wg)        // new thread
	time.Sleep(time.Second) // main thread
	fmt.Println((time.Now().UnixMilli() - x) / 1000)
	wg.Wait() // wait
	fmt.Println("end")
}
func nocat(val int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(val)
	time.Sleep(time.Second * 5)
	fmt.Println("val")
}
