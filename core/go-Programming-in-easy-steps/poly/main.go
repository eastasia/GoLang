package main

import "fmt"

type bird interface {
	speak() string
	move() string
}

type parrot struct{}

func (parrot) speak() string {
	return "Squawk, squawk!"
}

func (parrot) move() string {
	return "A parrot flies away."
}

type chicken struct{}

func (chicken) speak() string {
	return "Cluck, cluck!"
}

func (chicken) move() string {
	return "Chickens cannot fly."
}

func nudge(b bird) {
	fmt.Printf("\n%v \n", b.speak())
	fmt.Printf("%v \n\n", b.move())
}

func main() {

	var bird1 parrot
	var bird2 chicken

	nudge(bird1)
	nudge(bird2)

}
