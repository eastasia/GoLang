package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	fmt.Println("\nStarted At:", start.Format("03:04:05"))

	time.Sleep(5 * time.Second)

	finish := time.Now()
	fmt.Println("Finished At:", finish.Format("03:04:05"))

	fmt.Println("\nStart First?:", start.Before(finish))
	fmt.Println("Finish First?:", finish.Before(start))

	diff := finish.Sub(start)
	fmt.Println("\nTime Elapsed:", diff.Round(time.Second))

}
