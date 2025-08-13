package syntax

import "fmt"

type person struct{
	name   string
	age    int
	job    string
	salary int
}


func Info(){
	people := []person{
		{name: "Alice", age: 30, job: "Engineer", salary: 5000},
		{name: "Bob", age: 25, job: "Designer", salary: 4500},
	}
	
	// Modern for loop using range
	for _, p := range people {
		fmt.Println(p.name, p.age, p.job, p.salary)
	}
}