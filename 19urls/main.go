package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=golang&paymentid=jdhfihsh"

func main() {
	fmt.Println("Welcome to handling urls in go!")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the tye of this query param %T\n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("param is :", val)
	}

	//reverse you have all the informations in chunk and parts and eants to contruct url out of it
	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=aditi",
	}
	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)
}
