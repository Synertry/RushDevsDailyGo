package main

import "fmt"

var (
	mat = [][]int{
		{1, 3, 5, 8},
		{10, 11, 15, 16},
		{24, 27, 30, 31},
	}
	num = 120
)

func main() {
	fmt.Printf("In matrix %v the searched value exists:\n\t-> %d %v", mat, num, searchMatrix(mat, num))
}

// searchMatrix finds a target value in a sorted 2D matrix via binary search
// approach split the 2D into 1D and apply binary search
func searchMatrix(mat [][]int, target int) bool {
	rows := len(mat)
	if rows == 0 {
		return false
	}
	cols := len(mat[0])

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