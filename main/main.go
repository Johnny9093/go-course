package main

import (
	"errors"
	"fmt"
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

	myMap := make(map[string]int) // dic<string, int>
	myMap["first"] = 3
	myMap["second"] = 10

	fmt.Println(myMap)

	delete(myMap, "first")

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	fmt.Println(myMap["first"]) // = 0, but it doesnt actualy exist
	v, ok := myMap["first"]
	fmt.Println(v, ok) // ok = false, it does not exist

	if val, ok2 := myMap["first"]; ok2 { // ok2 is the actualy condition, after the semicolon
		fmt.Println(val)
	}
	// val and ok2 are dead here
}