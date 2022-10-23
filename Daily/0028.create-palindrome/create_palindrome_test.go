package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/IO"
	"github.com/Synertry/GoSysUtils/Math/Int"
	"github.com/Synertry/GoSysUtils/Str"
)

var resultBool bool

func TestMainFunc(t *testing.T) {
	want := fmt.Sprintf("Can create a palindrome if we remove a letter from %q:\n\t-> %t", word, create_palindrome(word))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestCreate_palindrome(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"intro":               {input: "abcdcbea", want: true},
		"intro2":              {input: "abccba", want: true},
		"intro3":              {input: "abccaa", want: false},
		"evenPalindromeShort": {input: "aa", want: true},
		// "single":              {input: "a", want: false},
		// "empty":               {input: "", want: false},
		"unevenPalindrome": {input: "abcba", want: true},
		"random":           {input: "vianefzbafwaaefruh", want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := create_palindrome(tc.input)
			if got != tc.want {
				t.Errorf("expected: %t, got: %t", tc.want, got)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func BenchmarkCreate_palindrome(b *testing.B) {
	type benchmark struct {
		name string
		len  int
	}

	maxExpStrLen := 3
	benchmarks := make([]benchmark, maxExpStrLen+1) // + 1 for single 10^0 -> 1

	for i := 0; i <= maxExpStrLen; i++ { // -1 as start, because substraction is more costly than addition
		strLen := Int.Pow(10, i)
		benchmarks[i] = benchmark{name: "StrLen10^" + strconv.Itoa(i), len: strLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			input, result := Str.GenRandom(bm.len), false
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = create_palindrome(input)
			}
			resultBool = result
		})
	}
}
