package main

import (
	"fmt"
)

var str = "aabcdefff"
var limit = 3

func main() {
	fmt.Printf("Longest substring in %q with max %d distinct characters is:\n\t-> %s", str, limit, longest_substring_with_k_distinct_characters(str, limit))
}

// longest_substring_with_k_distinct_characters returns the longest substring with at most k distinct characters
// by utilizing sliding windows as usual for substring problems.
// We use 4 pointers, two each for the current and the maximum sliding window.
// Using an array as a map allows us faster access, while also limiting space requirements.
// Time complexity: O(n)
// Space complexity: O(1) <- we use a fixed array size to store the characters
func longest_substring_with_k_distinct_characters(s string, k int) string {
	charmap := make([]uint8, 123) // ASCII table until 'z' at index 122; choose bigger uint if the sliding window could contain more than 255 character
	var l, r, wL, wR, dChars int  // left and right pointer, sliding window edges, distinct characters in the sliding window

	for r < len(s) {
		if charmap[s[r]] == 0 {
			dChars++
		}
		charmap[s[r]]++ // grow the sliding window
		r++
		for dChars > k { // shrink the window until it has at most k distinct characters
			charmap[s[l]]--
			if charmap[s[l]] == 0 {
				dChars--
			}
			l++
		}
		if r-l > wR-wL { // update the maximum sliding window
			wL, wR = l, r
		}
	}
	return s[wL:wR]
}

// longest_substring_with_k_distinct_characters_map returns the longest substring with at most k distinct characters
// This is the initial approach, using a map to store the distinct characters and their count.
// Hashing is slower than using an array, but it allows us to use any character.
// Time complexity: O(n)
func longest_substring_with_k_distinct_characters_map(s string, k int) string {
	charmap := make(map[uint8]uint8)
	var l, r, wL, wR, dChars int

	for r < len(s) {
		if charmap[s[r]] == 0 {
			dChars++
		}
		charmap[s[r]]++
		r++
		for dChars > k {
			charmap[s[l]]--
			if charmap[s[l]] == 0 {
				dChars--
			}
			l++
		}
		if r-l > wR-wL {
			wL, wR = l, r
		}
	}
	return s[wL:wR]
}
