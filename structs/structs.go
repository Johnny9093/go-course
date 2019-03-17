package structs

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type (
	Person struct {
		FirstName string `json:"first_name"` // check these lines again
		LastName string `json:"last_name"`
		age int
	}

	SuperHero struct {
		*Person
		SuperPower string
	}

	Rectangle struct {
		Width int
		Height int
	}

	Triangle struct {
		Width int
		Height int
	}

	Shape interface {
		GetArea() int
		GetName() string
	}

	MyInt int // typedef?
)

// Extend int, add custom functionality
func (this MyInt)Sig() int {
	if this < 0 {
		return -1
	}
	if this == 0 {
		return 0
	}
	if this > 0 {
		return 1
	}
}

// no need for * since the interface is already a pointer
// interface contains two pointers?.. to the concrete implementation..? check this out
// interface contains one pointer to the actual struct, and one pointer to the type of the struct
func PrintShape(shape Shape) {
	fmt.Printf("%s: My area is: %d\n", shape.GetName(), shape.GetArea())
}

func (this *Rectangle)GetArea() int {
	return this.Height * this.Width
}

func (this *Rectangle)GetName() string {
	return "Rectangle"
}

func (this *Triangle)GetArea() int {
	return this.Height * this.Width / 2
}

func (this *Triangle)GetName() string {
	return "Triangle"
}

// "ctor" (name doesn't matter here)
func NewPerson() *Person {
	return &Person{"Yoni","Lahav",23}
}

// function to method!
func (this *Person) Shorten(length int) {
	this.FirstName = this.FirstName[0:length]
	this.LastName = this.LastName[0:length]
}

func main() {
	a := NewPerson()
	a.Shorten(1)

	b := &SuperHero{SuperPower:"flight", Person:a}

	// b contains FirstName because Person is <Embedded> in SuperHero. sort of inheritance?
	println(b.FirstName)

	// b.Person.Shorten(3) = b.Shorten(3) !!

	// if we had func(this *SuperHero) Shorten then it would mask the first shorten
	// to access the base class (super): b.Person.Shorten()

	// built-in json (and other formats [xml, yaml..]) serialization
	aSer, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", a)
	fmt.Println(aSer) // or string(aSer)?

	var c *Person = &Person{} // have to initialize? check what happens without initialization
	err = json.Unmarshal(aSer, c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", c)

	rect := &Rectangle{4,3}
	PrintShape(rect)

	triangle := &Triangle{4, 3}
	PrintShape(triangle)

	var shape Shape
	PrintShape(shape) // what does this do? (nil?)

	shape = rect
	// now shape POINTS to rect (i think xD)

	//var anotherRectangle

	var myInt MyInt = 15
	fmt.Println(myInt) // 15. actual int

	rect1 := &Rectangle{15, 3}
	var shape1 Shape = rect1

	concShape, ok := shape1.(*Rectangle)
	println(ok)
	println(concShape) // address of rect1?

	concShape2, ok2 := shape1.(*Triangle)
	println(2)
	println(concShape2)
}