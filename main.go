package main

import (
	"fmt"
)

type Person struct{}

type Greeter interface{}

func sayHello(){
	fmt.Println("Data Send to Server.")
}

func add(a int, b int) int {
	return a + b;
}

func main() {
	fmt.Println("Getting Start Learn Go.")
}
