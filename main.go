package main

import (
	"fmt"

	Daily11 "github.com/Synertry/RushDevsDailyGo/Daily/11"
	Daily13 "github.com/Synertry/RushDevsDailyGo/Daily/13"
	Daily14 "github.com/Synertry/RushDevsDailyGo/Daily/14"
	Daily16 "github.com/Synertry/RushDevsDailyGo/Daily/16"
	Daily18 "github.com/Synertry/RushDevsDailyGo/Daily/18"
	Daily20 "github.com/Synertry/RushDevsDailyGo/Daily/20"
	Daily21 "github.com/Synertry/RushDevsDailyGo/Daily/21"
	Daily23 "github.com/Synertry/RushDevsDailyGo/Daily/23"
	Daily24 "github.com/Synertry/RushDevsDailyGo/Daily/24"
	Daily26 "github.com/Synertry/RushDevsDailyGo/Daily/26"
	Daily27 "github.com/Synertry/RushDevsDailyGo/Daily/27"
	Daily28 "github.com/Synertry/RushDevsDailyGo/Daily/28"
	Daily31 "github.com/Synertry/RushDevsDailyGo/Daily/31"
)

var (
	num            int
	word           string
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

	// Daily 23
	sliceOfInts = []int{8, 7, 2, 3, 4, 1, 5, 9, 0}
	num = 3
	fmt.Printf("The %dth largest element in slice %v is:\n\t-> %d\n\n", num, sliceOfInts, Daily23.FindKthLargest(sliceOfInts, num))

	// Daily 24
	sliceOfInts = []int{2, 1, -5, 4}
	num = -1
	fmt.Printf("Closest 3sum to %d in %v is:\n\t-> %v\n\n", num, sliceOfInts, Daily24.Closest_3sum(sliceOfInts, num))

	// Daily 26
	sliceOfInts = []int{1, 1, 2, 3, 4, 4, 4, 4, 4, 5, 5, 6, 7, 9}
	fmt.Printf("Length without dupes of %v is:\n", sliceOfInts)
	fmt.Printf("\t-> %d\n\n", Daily26.Remove_dups(sliceOfInts))

	// Daily 27
	sliceOfStrings = []string{"a", "b", "c", "d", "e", "f"}
	word = "bed"
	fmt.Printf("Can spell %q from %v:\n\t-> %t\n\n", word, sliceOfStrings, Daily27.CanSpell(sliceOfStrings, word))

	// Daily 28
	word = "abcdcbea"
	fmt.Printf("Can create a palindrome if we remove a letter from %q:\n\t-> %t\n\n", word, Daily28.Create_palindrome(word))

	// Daily 31
	num = 28
	fmt.Printf("Number %d is a perfect number:\n\t-> %t\n\n", num, Daily31.CheckPerfectNumber(num))
}
