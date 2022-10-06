package Daily16

// GenerateAllSubsets finds all combinations of a given set of integers [1, 2] -> [[], [1], [2], [1, 2]]
// Time complexity: O(2^n)
func GenerateAllSubsets(ints []int) (subsets [][]int) {
	length := len(ints)
	maxSubsets := 1 << length // maxSubsets is number of bitmasks

	for mask := 0; mask < maxSubsets; mask++ { // 0 to include empty set
		var subset []int

		for bit := 0; bit < length; bit++ { // loop through length of bits. bit -> index of mask
			if mask&(1<<bit) != 0 { // apply mask to check which index of ints to include
				subset = append(subset, ints[bit])
			}
		}

		subsets = append(subsets, subset)
	}

	return
}
