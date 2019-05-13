package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if data, err := ioutil.ReadFile("README.md");err!=nil{
		fmt.Println("README.md not found")
	} else {
		fmt.Print(string(data))
	}
}


