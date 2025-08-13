package syntax

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}	

func L(){

	// Use the counter closure
	next := counter()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3
	
	// Define and immediately call an anonymous function
	func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}("Sophea")

	greet := func(name string) string {
		return "Hi, " + name
	}
	fmt.Println(greet("Sophea"))

	numbers := []int{1, 2, 3, 4, 5}

	for _, n := range numbers {
		double := func(x int) int {
			return x * 2
		}(n)

		fmt.Println(double)
	}

}
