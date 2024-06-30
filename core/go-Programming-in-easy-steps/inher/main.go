package main

import "fmt"

type member struct {
	firstName string
	lastName  string
}

func (m member) fullName() string {
	return m.firstName + " " + m.lastName
}

type article struct {
	title string
	body  string
	member
}

func (a article) content() {
	fmt.Println("Title:", a.title)
	fmt.Println("Content:", a.body)
	fmt.Printf("Author: %v \n\n", a.fullName())
}

func main() {

	member1 := member{
		"Mike",
		"McGrath",
	}

	article1 := article{
		"Object Oriented Programming",
		"In Go, Composition emulates Inheritance",
		member1,
	}
	article1.content()

	article2 := article{
		"Object Oriented Programming",
		"Coming next... Polymorphism",
		member1,
	}
	article2.content()
}
