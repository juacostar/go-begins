package main

import "fmt"

type animal interface {
	move() string
}

type dog struct {
}

type fish struct {
}

type bird struct {
}

// overridign interface's method
func (bird) move() string {
	return "bird moving"
}

func (dog) move() string {
	return "dog moving"
}

func (fish) move() string {
	return "fish moving"
}

// setting a base interface's function
func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {

	d := dog{}
	f := fish{}
	b := bird{}

	moveAnimal(d)
	moveAnimal(f)
	moveAnimal(b)

}
