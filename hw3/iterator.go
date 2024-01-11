//Julia Rieger
//HW3 Part 3
//Oct 19 2023
//this program writes custom iterators to move through an integer slice

package main

import (
	"fmt"
)


//@param intSlice slice of ints to iterate through
//@return a function that iterates through this slice
func iterGenerator(intSlice []int) func() (int, bool){
	index := -1

	//iterator function returns (nextElement, nextElementIsValid)
	return func() (int, bool) {
		index++
		if (index<len(intSlice)) { //if still in bounds of slice return next element
			return intSlice[index], (index<len(intSlice))
		}
		return -1, false //if out of bounds of slice return -1 and mark not valid
	} //iterator()
}

func main() {	
	fmt.Printf("1) testing iterator\n")
	a := []int{0, 2, 4, 6, 8}
	iterator := iterGenerator(a)

	for j:=0; j>-1; j++ { //infinite loop
		e, v := iterator() //bind iteratorA returns to e and v
		if !v { //if no longer valid break
			break
		}
		fmt.Printf("next element: %v valid: %v\n", e, v) //print each element
	}

	fmt.Printf("\n2) verify iterators do not interact\n")
	iteratorA := iterGenerator(a)

	b := []int{1, 3, 5, 7, 9}
	iteratorB := iterGenerator(b)

	fmt.Printf("\nprinting e1 of a\n")
	e1, v1 := iteratorA() 
	fmt.Printf("\tnext element: %v valid: %v\n", e1, v1) 
	fmt.Printf("printing e1 of b\n")
	e2, v2 := iteratorB()
	fmt.Printf("\tnext element: %v valid: %v\n", e2, v2) 
	fmt.Printf("printing e2 of a\n")
	e3, v3 := iteratorA() 
	fmt.Printf("\tnext element: %v valid: %v\n", e3, v3) 
	fmt.Printf("printing e2 of b\n")
	e4, v4 := iteratorB() 
	fmt.Printf("\tnext element: %v valid: %v\n", e4, v4) 
	fmt.Printf("printing e3 of a\n")
	e5, v5 := iteratorA() 
	fmt.Printf("\tnext element: %v valid: %v\n", e5, v5) 
	fmt.Printf("printing e3 of b\n")
	e6, v6 := iteratorB() 
	fmt.Printf("\tnext element: %v valid: %v\n", e6, v6) 

	fmt.Printf("\n3) testing changing slice contents\n")

	c := []int{1,2,3,4,5}
	iteratorC := iterGenerator(c)
	fmt.Printf("original slice: %v\n",c)
	c[0] = 6
	c[1] = 7
	c[2] = 8
	c[3] = 9
	c[4] = 10
	fmt.Printf("\nafter changing slice content: \n")
	for j:=0; j>-1; j++ { //infinite loop
		e, v := iteratorC() //bind iteratorA returns to e and v
		if !v { //if no longer valid break
			break
		}
		fmt.Printf("next element: %v valid: %v\n", e, v) //print each element
	}

	fmt.Printf("\n4) testing appending new items to slice\n")
	d := []int{1,2,3,4,5}
	iteratorD := iterGenerator(d)
	fmt.Printf("original slice: %v\n",d)
	d = append(d, 6)

	fmt.Printf("\nafter changing slice content: \n")
	for j:=0; j>-1; j++ { //infinite loop
		e, v := iteratorD() //bind iteratorA returns to e and v
		if !v { //if no longer valid break
			break
		}
		fmt.Printf("next element: %v valid: %v\n", e, v) //print each element
	}
}