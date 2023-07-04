package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in go!")
	content := "This needs to in a file, first file creation cray crayyyyyyyy"
	file, err := os.Create("./mylcogofile.txt")
	CheckNilError(err)
	len, err := io.WriteString(file, content)
	CheckNilError(err)
	fmt.Println("Lenght is", len)
	defer file.Close()
	readFile("./mylcogofile.txt") //output is cray cray

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	CheckNilError(err)
	fmt.Println("Text data inside the file is: ", databyte)
	fmt.Println("Text data inside the file is: ", string(databyte))

}
func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
} //woohoooooo
