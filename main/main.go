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

	root, err := root(9)
	if err != nil {
		panic(err)
	}
	fmt.Println(root)
}