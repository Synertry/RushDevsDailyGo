package main

import "fmt"

var (
	sliceOfInts = []int{3, 3, 2, 1, 3, 2, 1} // holds unknown amount of 3 unique numbers (1,2,3)
)

func main() {
	fmt.Printf("The list %#v sorted is:\n\t-> %#v", sliceOfInts, sortNums(sliceOfInts))
}

// sortNums sorts a list of 3 unique numbers (1,2,3) in O(n) time with O(1) space
func sortNums(nums []int) []int {
	l, m, r := 0, 0, len(nums)-1 // left, middle, right pointers

	for m <= r {
		switch nums[m] {
		case 1: // if middle is 1, swap with left and move both lower pointers up
			nums[l], nums[m] = nums[m], nums[l]
			l++
			m++
		case 2: // if middle is 2, move middle pointer up
			m++
		case 3: // if middle is 3, swap with right and move right pointer down
			nums[m], nums[r] = nums[r], nums[m]
			r--
		}
	}

	return nums
}

// // identify3Nums returns the three unique numbers in a slice of ints
// func identify3Nums(nums []int) (one, two, three int) {
// 	// use a map to hold the count of each number
// 	counts := make(map[int]int)
//
// 	// iterate through nums and increment the count of each number
// 	for _, num := range nums {
// 		counts[num]++
// 	}
// 	uniqueNums := make([]int, 0, len(counts))
// 	// retrieve unique numbers
// 	for num := range counts {
// 		uniqueNums = append(uniqueNums, num)
// 	}
//
// 	// sort unique numbers with pointer swapping in ascending order
// 	// TODO: implement proper short sort w/o recursion
// 	for i := 0; i < len(uniqueNums)-1; i++ {
// 		if uniqueNums[i] > uniqueNums[i+1] {
// 			uniqueNums[i], uniqueNums[i+1] = uniqueNums[i+1], uniqueNums[i]
// 		}
// 	}
//
// 	return uniqueNums[0], uniqueNums[1], uniqueNums[2]
// }
