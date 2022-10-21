package main

import "fmt"

var sliceOfInts = []int{1, 1, 2, 3, 4, 4, 4, 4, 4, 5, 5, 6, 7, 9}

func main() {
	fmt.Printf("Length without dupes of %v is:\n\t-> %d", sliceOfInts, remove_dups(sliceOfInts))
	// fmt.Printf("\t-> %d\n\n", remove_dups(sliceOfInts))
}

func remove_dups(nums []int) (length int) {
	for ptr := 1; ptr < len(nums); ptr++ {
		if nums[length] != nums[ptr] {
			length++
			nums[length], nums[ptr] = nums[ptr], nums[length]
		}
	}
	return length + 1
}