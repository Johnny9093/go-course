package pointers

import "fmt"

// pointers are safe; address-of, content-of operators
// objects are deep copied when passed to functions
// intended way is to pass by ref

// type vs struct?
type Person struct {
	FirstName string
	LastName string
}

func Shorten(person Person) {
	person.FirstName = person.FirstName[0:1]
	person.LastName = person.LastName[0:1]
}

// no change is necessary! person.FirstName works just as person->FirstName would work in c/c++
func ShortenP(person *Person) {
	person.FirstName = person.FirstName[0:1]
	person.LastName = person.LastName[0:1]
}

func inc(value int) int {
	value += 1
	return value
}

func incP(addr *int) {
	*addr += 1
}

func main() {
	a := 42
	pa := &a // address-of. pa type is *int. **int doesnt exist. no pointers of pointers

	// pointers are safe so no pointer arithmetic.
	// unsafe code is possible but not allowed by default.

	fmt.Printf("%d\n", a)
	fmt.Printf("%+v\n", pa)

	fmt.Printf("pa (content): %d\n", *pa) // content-of, number formatter
	fmt.Printf("a (address): %p\n", &a) // address-of, pointer formatter

	*pa = 13
	fmt.Printf("a: %d\n", a) // a is 13

	inc(a)
	fmt.Printf("a: %d\n", a)

	incP(&a)
	fmt.Printf("a: %d\n", a)

	p := Person{FirstName:"Yoni", LastName:"Lahav"}
	Shorten(p)
	fmt.Printf("p: %+v\n", p) // nothing changes; Shorten deep-copies p


	// best practice: always work by ref! better performance and memory consumption
	person := &Person{FirstName:"first",LastName:"last"}
	ShortenP(person)
	fmt.Printf("p: %+v\n", person) // changes because pointer is passed
}
