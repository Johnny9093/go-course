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

	slc := []int{1,2,3,4,5} // slice because size isnt stated explicitly
	slc = append(slc, 6)
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))
}