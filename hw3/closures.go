//Julia Rieger
//HW3 Part 2
//Oct 19 2023
//this program contains functions and closures to compute values of numeric series

package main

import (
	"fmt"
)

//computes a squared geometric series
func geometricComputer(stateInt int) int {
	return stateInt*stateInt
}

//computes an arithmetic series that counts by 5s
func arithmeticComputer(stateInt int) int {
	return stateInt * 5
}

//computes an arithmetic series that consists of all positive odd numbers
func oddComputer(stateInt int) int {
	return stateInt * 2 - 1
}

//returns a function to compute values of arbitrary numeric series whos state is given by startIndex
func maker(computer func(int) int, startIndex int) func() int{
	index := startIndex
	return func() int {
		index++
		return computer(index-1)
	}
}

func main() {	
	fmt.Printf("test run: standard geometric series\n")
	nextInSeriesG := maker(geometricComputer, 1)
	fmt.Printf("%v\n",nextInSeriesG())
	fmt.Printf("%v\n",nextInSeriesG())
	fmt.Printf("%v\n",nextInSeriesG())
	fmt.Printf("%v\n",nextInSeriesG())
	fmt.Printf("%v\n",nextInSeriesG())
	fmt.Printf("%v\n",nextInSeriesG())


	fmt.Printf("\ntest run: arithmetic series by 5s\n")
	nextInSeriesA := maker(arithmeticComputer, 1)
	fmt.Printf("%v\n",nextInSeriesA())
	fmt.Printf("%v\n",nextInSeriesA())
	fmt.Printf("%v\n",nextInSeriesA())
	fmt.Printf("%v\n",nextInSeriesA())
	fmt.Printf("%v\n",nextInSeriesA())
	fmt.Printf("%v\n",nextInSeriesA())

	fmt.Printf("\nadditional tests:\n")
	fmt.Printf("Odds starting at 1\n")
	fmt.Printf("Geometric starting at index 5\n")
	fmt.Printf("Arithmetic starting at index 3\n")
	nextInSeriesOTest := maker(oddComputer, 1)
	nextInSeriesGTest := maker(geometricComputer, 5)
	nextInSeriesATest := maker(arithmeticComputer, 3)
	fmt.Printf("next O: %v\n",nextInSeriesOTest())
	fmt.Printf("next G: %v\n",nextInSeriesGTest())
	fmt.Printf("next A: %v\n",nextInSeriesATest())
	fmt.Printf("next O: %v\n",nextInSeriesOTest())
	fmt.Printf("next G: %v\n",nextInSeriesGTest())
	fmt.Printf("next A: %v\n",nextInSeriesATest())
	fmt.Printf("next O: %v\n",nextInSeriesOTest())
	fmt.Printf("next G: %v\n",nextInSeriesGTest())
	fmt.Printf("next A: %v\n",nextInSeriesATest())
	fmt.Printf("next O: %v\n",nextInSeriesOTest())
	fmt.Printf("next G: %v\n",nextInSeriesGTest())
	fmt.Printf("next O: %v\n",nextInSeriesOTest())
	fmt.Printf("next O: %v\n",nextInSeriesOTest())
	fmt.Printf("next G from original G func: %v\n",nextInSeriesG())
	fmt.Printf("next A from original A func: %v\n",nextInSeriesA())
}