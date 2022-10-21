package main

import "fmt"

func main() {
	sliceOfStrings := []string{"a", "b", "c", "d", "e", "f"}
	word := "bed"
	fmt.Printf("Can spell %q from %v:\n\t-> %t\n\n", word, sliceOfStrings, canSpell(sliceOfStrings, word))
}

// canSpell returns true if the letters in the array can be used to spell the string.
func canSpell(magazine []string, note string) bool {
	// Create a map of the magazine letters.
	letters := make(map[string]int)
	for _, letter := range magazine {
		letters[letter]++
	}
	// Check if the letters in the note are in the map.
	for _, letter := range note {
		if letters[string(letter)] == 0 {
			return false
		}
		letters[string(letter)]--
	}
	return true
}
