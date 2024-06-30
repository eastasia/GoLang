package main

import (
	"fmt"
	"time"
)

func main() {

	dt := time.Now()

	fmt.Printf("DateTime: %v \n", dt)
	fmt.Printf("DateTime Type: %T \n", dt)

	fmt.Printf("Today is: %v \n", dt.Weekday())

	y, m, d := dt.Date()
	fmt.Printf("Date: %v %v, %v \n", m, d, y)

	yr, wk := dt.ISOWeek()
	fmt.Printf("Week No.: %v in %v\n", wk, yr)

	dy := dt.YearDay()
	fmt.Printf("Day No.: %v \n", dy)

}
