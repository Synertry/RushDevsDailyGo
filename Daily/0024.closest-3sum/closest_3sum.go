package main

import (
	"fmt"
	"sort"
)

var (
	sliceOfInts = []int{2, 1, -5, 4}
	num         = -1
)

func main() {
	fmt.Printf("Closest 3sum to %d in %v is:\n\t-> %v", num, sliceOfInts, closest_3sum(sliceOfInts, num))
}

func closest_3sum(nums []int, target int) (threeNums []int) {
	sort.Ints(nums)    // first sort, then search by pattern
	tempSum := 1 << 31 // max int32

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] { // skip duplicates
			continue
		}

		low, high := i+1, len(nums)-1 // two pointers

		for low < high { // binary search with two pointers
			s := nums[i] + nums[low] + nums[high] // sum of two pointers and current index

			if s < target { // if sum is less than target, move low pointer up
				low++
				if tempSum > target-s {
					tempSum = target - s
					threeNums = []int{nums[i], nums[low], nums[high]}
				}
			} else if s > target { // if sum is greater than target, move high pointer down
				high--
				if tempSum > s-target {
					tempSum = s - target
					threeNums = []int{nums[i], nums[low], nums[high]}
				}
			} else {
				return []int{nums[i], nums[low], nums[high]}
			}

		}
	}
	return
}