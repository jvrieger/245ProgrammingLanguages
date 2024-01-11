//Julia Rieger
//HW3 Part 1
//Oct 19 2023
//this program proves Go uses static scoping, further discussion in README

package main

import (
	"fmt"
)

var a = 1 //global var in global scope

//f() is called by g()
//returns the int a (from what scope?)
func f() int {
	return a
}

//g() declares its own local a var
//calls f()
func g() int {
	a := 2
	fmt.Printf("bypass Go's declaration without use error: %v\n", a) //compile time error handling
	return f()
}

func main() {
	fmt.Printf("value of a when g() is called: %v\n", g()) //call g()
}