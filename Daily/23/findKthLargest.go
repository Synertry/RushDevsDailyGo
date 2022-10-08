package Daily23

func FindKthLargest(arr []int, k int) int {
	// allocs
	arrLen := len(arr)

	// exceptions
	if arrLen == 1 { // returning here would be O(1)
		return arr[0]
	}

	// smallest or largest then just iterate once
	if k == 1 || k == arrLen { // returning here would be true O(n)
		return findEdge(arr, k-1, arrLen)
	}

	// else divide and conquer method with partitioning borrowing from quicksort,
	// but we won't sort the whole slice, just the part we need, so a quickselect
	// Time complexity: O(n) on average
	low, high := 0, arrLen-1
	for { // O(n) + O(n/2) + O(n/4) + ... + O(1) = O(n)
		pivot := partition(arr, low, high) // here we half the search space
		if pivot == k-1 {
			return arr[pivot]
		} else if pivot > k-1 {
			high = pivot - 1
		} else {
			low = pivot + 1
		}
	}
}

// partition moves all elements smaller than pivot to the left and all elements
// greater than pivot to the right of the pivot
// returns the index of the pivot
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for ptr := low; ptr < high; ptr++ {
		if arr[ptr] >= pivot {
			i++
			arr[i], arr[ptr] = arr[ptr], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// findEdge finds the smallest or largest element in the slice
// it is used when k is 1 or len(arr) to skip further iterations and achieve
// Time complexity: O(n)
func findEdge(arr []int, edge int, arrLen int) int {
	// find depending edge int in slice
	if edge == 0 {
		for i := 0; i < arrLen; i++ {
			if edge == 0 {
				if arr[edge] > arr[i] { // find smallest
					arr[edge] = arr[i]
				}
			} else {
				if arr[edge] < arr[i] { // find largest
					arr[edge] = arr[i]
				}
			}
		}
	}
	return arr[edge]
}
