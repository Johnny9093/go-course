package main

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T){
	fmt.Println("Hello test")
	t.Fail()
}
