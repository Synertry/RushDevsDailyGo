package Daily20

func FindDupesInSortedList(list []int, target int) []int {
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
				if list[mid-1] == target {
					indices[0], indices[1] = mid-1, mid
				}
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
