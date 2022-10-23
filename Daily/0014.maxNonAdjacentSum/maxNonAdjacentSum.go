package main

import "fmt"

var sliceOfInts = []int{2, 1, 2, 7, 3}

func main() {
	fmt.Printf("Maximum non-adjacent sum of slice %v is:\n\t-> %d", sliceOfInts, maxNonAdjacentSum(sliceOfInts))
}

func maxNonAdjacentSum(elems []int) int {
	eLen := len(elems)

	// preliminary return checks // avoid them if possible
	// if elemLen == 0 {
	// 	return 0
	// } else if elemLen == 1 {
	// 	return elems[0]
	// }

	// allocate slice
	intList := make([]int, eLen+1)

	// starting pair // not needed anymore as I implemented index check
	// intList[0] = elems[0] // first element is already initialized as 0
	// intList[1] = max(elems[0], elems[1])

	// check at each index
	for i := 1; i <= eLen; i++ {
		befAdj := 0 // Go does not have ternary if-functions
		if i > 1 {  // first two elements are starting base
			befAdj = intList[i-2] // assign before adjacent int
		}
		// at each index we check if the already max combination is higher than the current element plus the before adjacent element
		intList[i] = max(intList[i-1], befAdj+elems[i-1])
	}

	return intList[eLen] // return last highest determined sum
}

// max is helper function to return the higher number of two integers
// simple and no special cases
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
