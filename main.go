package main

// how to alias when import
import (
	"fmt"
	v "start/variable"
)

func main() {

	// Input
	var name string
	fmt.Print("Enter your name : ")
	fmt.Scan(&name)
	fmt.Printf("My name is %s\n", name)

    fmt.Println(v.PI)
    fmt.Println(v.Greeting)
}
