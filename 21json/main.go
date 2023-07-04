package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json in go!")
	Encodejson()
	Decodejson()
}

func Encodejson() {
	lcocourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 399, "LearnCodeOnline.in", "bch123", nil},
	}
	//package data as json data
	finalJson, err := json.MarshalIndent(lcocourses, "", "\t") //a way to implementation of json
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func Decodejson() {
	jsonDataFromWeb := []byte(`	
	{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "LearnCodeOnline.in",
			"Tags": ["web-dev","js"]
	}
`)
	var lcocourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		//fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON isnt valid :(")
	}
	//cases where uou gotta add data to key-value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("for key %v value is %v and type is %T\n", k, v, v)
	}
}
