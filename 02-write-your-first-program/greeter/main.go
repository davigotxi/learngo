package main

import (
	"fmt"
	"os"
)

func main() {

	//var name string

	//name = os.Args[1]
	name := os.Args[1]

	fmt.Println("Hello great", name, "!")

	name, age := "Gandalf", 2019

	//fmt.Printf("%#v\n", os.Args)
	fmt.Println("Hello great")
}
