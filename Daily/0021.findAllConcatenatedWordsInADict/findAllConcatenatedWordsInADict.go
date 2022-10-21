package main

import "fmt"

func main() {
	sliceOfStrings := []string{"tech", "lead", "techlead", "cat", "cats", "dog", "catsdog"}
	fmt.Printf("The concatenated words for %v:\n\t-> %v\n\n", sliceOfStrings, findAllConcatenatedWordsInADict(sliceOfStrings))
}

func findAllConcatenatedWordsInADict(dict []string) []string {
	// hashmap the lengths and words
	mapLen := make(map[int]bool, len(dict))
	mapWord := make(map[string]bool, len(dict))

	for _, word := range dict { // fill the hashmaps
		mapLen[len(word)] = true
		mapWord[word] = true
	}

	concatWords := make([]string, 0, len(dict))

	for _, word := range dict { // first loop ->  n
		if checkConcat(word, mapLen, mapWord, false) { // call with isSplit = false to not match word itself
			concatWords = append(concatWords, word)
		}
	}

	return concatWords
}

func checkConcat(word string, mapLen map[int]bool, mapWord map[string]bool, isSplit bool) bool {
	for wLen := 1; wLen < len(word); wLen++ { // second loop -> n^2, for every recurse add *n
		if mapLen[wLen] && // check if a word with specific length exists
			mapWord[word[:wLen]] && // true for first part of concatenated word exists
			checkConcat(word[wLen:], mapLen, mapWord, true) { // recurse to search for second part of concatenated word
			return true
		}
	}

	return isSplit && // if nothing has been found isSplit will still be false as called and the word itself will not be added
		mapLen[len(word)] && mapWord[word] // will be true for second part of concatenated word if found
}
