package main

import (
	"fmt"
	"time"
)

func main() {

	dt := time.Date(2025, time.January, 1, 12, 0, 0, 0, time.Local)
	fmt.Printf("\nDateTime: %v \n\n", dt)

	dt = dt.AddDate(2, 6, 3)
	fmt.Printf("New DateTime: %v \n\n", dt)

	layout := "2006-Jan-02 03:04PM "
	str := "2030-Dec-25 12:30AM"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Parsed DateTime: %v \n", t)
	}

}
