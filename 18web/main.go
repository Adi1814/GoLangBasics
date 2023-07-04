package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("LCO web request in go!")
	response, err := http.Get(url)
	CheckNilError(err)
	fmt.Printf("response if of type %T\n", response)
	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	CheckNilError(err)
	fmt.Println("read response is in string", string(databytes))
}
func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
