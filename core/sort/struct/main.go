package main

import (
	"fmt"
	"sort"
)

type person []string

func (p person) Len() int { return len(p) }
func (p person) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p person) Less(i, j int) bool { return p[i] < p[j] }

func main(){

	s := person{"Susan", "Tyler", "Ava", "Nick"}
	sort.Sort(s)
	fmt.Println(s)

	fmt.Println(sort.StringSlice(s))

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	//fmt.Println(sort.Reverse(sort.StringSlice(s)))
	
	fmt.Println(s)
}