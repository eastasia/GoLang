package main

import (
	"fmt"
	"time"
)

func main() {

	dt := time.Now()

	fmt.Println("\nDefault Format:", dt)
	fmt.Println("Unix format:", dt.Format(time.UnixDate))
	fmt.Println("ANSIC format:", dt.Format(time.ANSIC))
	fmt.Println("RFC3339 format:", dt.Format(time.RFC3339))
	fmt.Println("Custom Format:", dt.Format("January 2, 2006 [Monday]"))

	fmt.Println("\nUS Date Format:", dt.Format("January 2, 2006"))
	fmt.Println("UK Date Format:", dt.Format("2 January, 2006"))

	fmt.Println("\nTime 12-hour:", dt.Format(time.Kitchen))
	fmt.Println("Time 24-Hour:", dt.Format("15:04"))

}
