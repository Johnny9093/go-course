package main

import (
	"errors"
	"math"
)

func divide(a, b int) (result, residue int) { // private func, returns 2 ints
	result = a / b
	residue = a % b
	return
}

func root(num float64) (root float64, error error) {
	if num < 0 {
		error = errors.New("negative input")
		return
	}

	root = math.Sqrt(num)
	return
}

// clojures
func getFib() func() int {

	first := 0
	second := 1

	return func() int {
		res := first
		first, second = second, first + second;
		return res;
	}
}

func panicAttack(a, b int) float64 {
	if b == 0 {
		panic("division by zero")
	}

	return float64(a) / float64(b)
}

func main(){
	// var name string = "Yoni"
	// var name string
	// var name = "Yoni"
	// name := "Yoni"
	// fmt.Printf("Hi there, %s", name)

	// myStruct := internal.MyStruct{FirstName: "Yoni"}
	//myStruct := internal.MyStruct{FirstName: "Yoni"}
	//fmt.Printf("%+v\n", myStruct)
	//internal.Something()

	//a:= 9
	//b := 3
	//_, r2 := divide(a, b) // not using first return value
	//println(r2)

	//root, err := root(9)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(root)

	// arr := [3] int {1,2,3}
	//arr := [...] int {1,2,3,4,5}
	//slc := arr[:]
	//slc[2] = 15 // changes arr since it points to the same array
	//
	//fmt.Println(arr)
	//fmt.Println(len(arr))
	//fmt.Println(cap(arr))
	//
	//fmt.Println(slc)
	//fmt.Println(len(slc))
	//fmt.Println(cap(slc))

	//slc := []int{1,2,3,4,5} // slice because size isnt stated explicitly
	//slc = append(slc, 6)
	//fmt.Println(slc)
	//fmt.Println(len(slc))
	//fmt.Println(cap(slc))

	//slc := make([]int, 3, 5) // make slice of ints, length 3 and cap 5
	//slc = append(slc, 1)
	//slc = append(slc, 2)
	//slc = append(slc, 3) // capacity doubles
	//
	//fmt.Println(slc)
	//fmt.Println(len(slc))
	//fmt.Println(cap(slc))
	//
	//slc2 := []int {4,5}
	//slc = append(slc, slc2...) // spread operator - opens up all of the initialized values of slc2
	//
	//for i := range slc { // i = indices
	//	fmt.Println(i)
	//}
	//
	//for k, v := range slc {
	//	fmt.Println(k, v) // key = index / key
	//}
	//
	//for _, v := range slc {
	//	fmt.Println(v) // v = values
	//}

	//myMap := make(map[string]int) // dic<string, int>
	//myMap["first"] = 3
	//myMap["second"] = 10
	//
	//fmt.Println(myMap)
	//
	//delete(myMap, "first")
	//
	//for k, v := range myMap {
	//	fmt.Println(k, v)
	//}
	//
	//fmt.Println(myMap["first"]) // = 0, but it doesnt actualy exist
	//v, ok := myMap["first"]
	//fmt.Println(v, ok) // ok = false, it does not exist
	//
	//if val, ok2 := myMap["first"]; ok2 { // ok2 is the actualy condition, after the semicolon
	//	fmt.Println(val)
	//}
	//// val and ok2 are dead here

	//var fib = getFib()
	//fmt.Println(fib())
	//fmt.Println(fib())
	//fmt.Println(fib())
	//fmt.Println(fib())
	//fmt.Println(fib())
	//fmt.Println(fib())
	//fmt.Println(fib())
	//fmt.Println(fib())

	//// shortened if
	//if i := 4; i < 15 {
	//
	//} // i dies here
	//
	//if i := 4; i < 15 {
	//} else { // works
	//}
	////else { // compilation error
	////}
	//
	//switch i := 5; i {
	//case 5:
	//	println(i)
	//	fallthrough // break is default not necessary, in stead use explicit fallthrough
	//default:
	//	println("Unknown value")
	//}
	//
	//// anti-switch, easier way to write ifs
	//i := 3
	//switch {
	//case i < 4:
	//case i == 4:
	//case i > 4:
	//}

	// exception handling
	// panics are only for severe situations (db connection lost etc)
	//a, b := 8, 2
	//println(a, b)

	//// defer is using a stack and runs at the end of the function
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("recovered:", r)
	//	}
	//}()
	//
	//panic("oh no")
	// ask about print orders!




	// lesson 3
	// http stuff


}