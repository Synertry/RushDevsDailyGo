package main

import "fmt"

func main() {
	mat := [][]int{{1, 2, 3, 4}, {2, 4, 6, 8}, {3, 4, 5}}
	fmt.Printf("Intersections of %#v are:\n\t-> %v\n\n", mat, intersection(mat))

}

// intersection returns the intersection of a list of []int
// Arrays are sorted in ascending order
func intersection(lists [][]int) (iX []int) {
	var m0 map[int]bool
	maps := getIntMaps(lists)

	for _, m := range maps { // find smallest map
		if m0 == nil || len(m) < len(m0) {
			m0 = m
		}
	}

	for n := range m0 { // range over the smallest map
		if checkKey(maps, n) { // check if key exists in all maps
			iX = append(iX, n)
		}
	}

	return
}

// check if key exists in all maps
func checkKey(maps []map[int]bool, key int) bool {
	for _, m := range maps {
		if !m[key] {
			return false
		}
	}
	return true
}

// getIntMaps returns a list of maps from [][]int
func getIntMaps(lists [][]int) []map[int]bool {
	intMaps := make([]map[int]bool, len(lists))
	for i := range lists {
		intMaps[i] = getIntMap(lists[i])
	}

	return intMaps
}

// getIntMap returns a map from a []int
func getIntMap(list []int) map[int]bool {
	intMap := make(map[int]bool, len(list))
	for i := range list {
		intMap[list[i]] = true
	}

	return intMap
}
