package main

import "fmt"

func main() {
	sliceOfInts := []int{1, 2, 3}
	fmt.Printf("All generated subsets of slice %v are:\n\t-> %v\n\n", sliceOfInts, generateAllSubsets(sliceOfInts))
}

// GenerateAllSubsets finds all combinations of a given set of integers [1, 2] -> [[], [1], [2], [1, 2]]
// Time complexity: O(2^n)
func generateAllSubsets(ints []int) (subsets [][]int) {
	length := len(ints)
	maxSubsets := 1 << length // maxSubsets is number of bitmasks

	for mask := 0; mask < maxSubsets; mask++ { // 0 to include empty set
		subset := make([]int, 0)

		for bit := 0; bit < length; bit++ { // loop through length of bits. bit -> index of mask
			if mask&(1<<bit) != 0 { // apply mask to check which index of ints to include
				subset = append(subset, ints[bit])
			}
		}

		subsets = append(subsets, subset)
	}

	return
}
