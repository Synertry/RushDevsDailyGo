package Daily28

// Create_palindrome checks if a palindrome can be created when 1 letter gets removed
func Create_palindrome(s string) bool {
	for low, high := 0, len(s)-1; low < high; low, high = low+1, high-1 {
		if s[low] != s[high] { // outer check -> removing 1 letter from either side
			return isPalindrome(s, low+1, high) || isPalindrome(s, low, high-1) // now stepping into inner check
		}
	}

	// func must at least remove a letter even if it is a palindrome already
	// we pass isPalindrome a string with middle letter removed
	return isPalindrome(s[:len(s)/2]+s[len(s)-len(s)/2+(len(s)+1)%2:], 0, len(s)-2)
}

// isPalindrome checks if passed string is a palindrome
func isPalindrome(s string, low, high int) bool {
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}
