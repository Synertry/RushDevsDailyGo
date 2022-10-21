package main

import "fmt"

func main() {
	sliceOfInts := []int{1, 1, 3, 5, 7}
	num := 1
	fmt.Printf("In sorted list %v the searched number %d has indices:\n\t-> %v\n\n", sliceOfInts, num, find_num(sliceOfInts, num))
}

func find_num(list []int, target int) []int {
	length := len(list)
	indices := []int{-1, -1}
	low, high := 0, length-1

	for low <= high {
		mid := (low + high) / 2

		if list[mid] == target { // lookaround the index for duplicates
			indices[0] = mid

			if length == 1 { // simplify further lookaround and list length is 1 check
				break
			} else if mid == 0 { // lower bound, index can't be negative
				if list[mid+1] == target {
					indices[1] = mid + 1
				}
			} else if mid == length-1 { // upper bound, index can't be out of bounds
				// never happens because we climb up from the middle
				// but neeed for index out of range protection
				// if list[mid-1] == target {
				// 	indices[0], indices[1] = mid-1, mid
				// }
			} else {
				if list[mid+1] == target { // lookahead
					indices[1] = mid + 1
				} else if list[mid-1] == target { // lookbehind
					indices[0], indices[1] = mid-1, mid
				}
			}
			break
		} else if list[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return indices
}
