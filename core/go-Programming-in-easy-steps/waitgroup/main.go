package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go report(i, &wg)
	}

	wg.Wait()
}

func report(i int, wg *sync.WaitGroup) {
	fmt.Printf("\nGoroutine %v Started", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("\n\t\t\tGoroutine %v Ended", i)
	wg.Done()
}
