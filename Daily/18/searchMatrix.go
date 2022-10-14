package Daily18

// SearchMatrix finds a target value in a sorted 2D matrix via binary search
// approach split the 2D into 1D and apply binary search
func SearchMatrix(mat [][]int, target int) bool {
	rows := len(mat)
	cols := len(mat[0])

	if rows == 0 {
		return false
	}

	low, high := 0, rows*cols-1 // split of 2D into 1D

	for low <= high {
		node := (low + high) / 2
		row := node / cols
		col := node % cols
		nodeVal := mat[row][col]
		if nodeVal == target {
			return true
		} else if nodeVal < target {
			low = node + 1
		} else {
			high = node - 1
		}
	}
	return false
}