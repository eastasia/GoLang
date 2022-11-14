package main

import(
	"fmt"
	"sort"
)

type player struct{ 
	Name string
	Country string
}

func (p player) String() string {
	return fmt.Sprintf("toString() - %s: %s", p.Name, p.Country)
}

type byCountry []player

func (c byCountry) Len() int { return len(c) }
func (c byCountry) Swap (i, j int) { c[i], c[j] = c[j], c[i] }
func (c byCountry) Less (i, j int) bool { return c[i].Country < c[j].Country }

func main(){
	
	p := []player{
		{ "Roger Federer","Switzerland" },
		{ "Lionel Messi", "Argentina"},
		{ "Micheal Jordan", "USA"},
	}

	fmt.Println(&p[0])
	fmt.Println(p)

	fmt.Println()
	sort.Sort(byCountry(p))
	fmt.Println(p)

}
