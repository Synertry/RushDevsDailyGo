package Daily18

// SearchMatrixProto is my initial approach to the problem
func SearchMatrixProto(mat [][]int, target int) bool {
	// compare every y[0] against target starting from the middle
	lenY := len(mat)
	lenDY := lenY / 2
	var numD int
	// compare if target is lower, bigger or equal to middle y[0]
	numD = mat[lenDY-1][0] - target
	switch {
	case numD == 0:
		return true
	case numD > 0:
		lenDY = lenDY / 2
	case numD < 0:
		lenDY = lenDY + (lenY-lenDY)/2
	}
	return false
}

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
