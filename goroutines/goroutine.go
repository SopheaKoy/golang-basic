package goroutines

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func work() {
	defer wg.Done()
	fmt.Println("Working....")
}

func Controller() {
	wg.Add(1)
	go work()
}
