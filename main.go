package main

import (
	"fmt"

	"github.com/Synertry/RushDevsDailyGo/Daily/11"
	"github.com/Synertry/RushDevsDailyGo/Daily/13"
	"github.com/Synertry/RushDevsDailyGo/Daily/14"
	"github.com/Synertry/RushDevsDailyGo/Daily/16"
	"github.com/Synertry/RushDevsDailyGo/Daily/18"
	"github.com/Synertry/RushDevsDailyGo/Daily/20"
	"github.com/Synertry/RushDevsDailyGo/Daily/21"
)

var (
	num            int
	sliceOfInts    []int
	sliceOfStrings []string
	mat            [][]int
)

func main() {
	// Daily 11
	num = 23
	fmt.Printf("Number of set bits in %d is:\n\t-> %d\n", num, Daily11.CountOneBits(num))
	fmt.Printf("Number of set bits in %d is (without bitwise operation):\n\t-> %d\n\n", num, Daily11.CountOneBitsNonBitOps(num))

	// Daily 13
	sliceOfInts = []int{3, 5, 3, 3, 2, 4, 3}
	fmt.Printf("Majority element of slice %v is:\n\t-> %d\n\n", sliceOfInts, Daily13.FindMajorElem(sliceOfInts))

	// Daily 14
	sliceOfInts = []int{2, 1, 2, 7, 3}
	fmt.Printf("Maximum non-adjacent sum of slice %v is:\n\t-> %d\n\n", sliceOfInts, Daily14.MaxNonAdjacentSum(sliceOfInts))

	// Daily 16
	sliceOfInts = []int{1, 2, 3}
	fmt.Printf("All generated subsets of slice %v are:\n\t-> %v\n\n", sliceOfInts, Daily16.GenerateAllSubsets(sliceOfInts))

	// Daily 18
	mat = [][]int{
		{1, 3, 5, 8},
		{10, 11, 15, 16},
		{24, 27, 30, 31},
	}
	num = 120
	fmt.Printf("In matrix %v the searched value exists:\n\t-> %d %v\n\n", mat, num, Daily18.SearchMatrix(mat, num))

	// Daily 20
	sliceOfInts = []int{1, 1, 3, 5, 7}
	num = 1
	fmt.Printf("In sorted list %v the searched number %d has indices:\n\t-> %v\n\n", sliceOfInts, num, Daily20.FindDupesInSortedList(sliceOfInts, num))

	// Daily 21
	sliceOfStrings = []string{"tech", "lead", "techlead", "cat", "cats", "dog", "catsdog"}
	fmt.Printf("The concatenated words for %v:\n\t-> %v\n\n", sliceOfStrings, Daily21.FindAllConcatenatedWordsInADict(sliceOfStrings))
}