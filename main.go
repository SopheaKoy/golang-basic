package main

// how to alias when import
import (
	"fmt"
	gettings "start/greet"
	v "start/variable"
)

func main() {

	msg, _ := gettings.Hello("Seav mey")
	fmt.Println(msg)

	// Input
	var name string
	fmt.Print("Enter your name : ")
	fmt.Scan(&name)
	fmt.Printf("My name is %s\n", name)

    fmt.Println(v.PI)
    fmt.Println(v.Greeting)
}
