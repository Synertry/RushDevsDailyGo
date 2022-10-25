package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/Math/Int"
	"github.com/Synertry/GoSysUtils/Str"
	"github.com/Synertry/RushDevsDailyGo/internal/Helper"
)

type benchmark struct {
	name string
	len  int
}

const maxLenExp = 5

var (
	resultString string
	benchmarks   = make([]benchmark, maxLenExp+1)
	tests        = map[string]struct {
		input string
		k     int
		want  string
	}{
		"intro":            {input: "aabcdefff", k: 3, want: "defff"},
		"intro2":           {input: "aabcdefff", k: 1, want: "fff"},
		"single":           {input: "a", k: 1, want: "a"},
		"doublesingle":     {input: "bb", k: 1, want: "bb"},
		"empty":            {input: "", k: 1, want: ""},
		"higherk":          {input: "aabcdefff", k: 10, want: "aabcdefff"},
		"maxWindowAtStart": {input: "aaaabcdefff", k: 2, want: "aaaab"},
	}
)

func init() {
	for i := 0; i <= maxLenExp; i++ {
		benchmarks[i] = benchmark{name: "StrLen10^" + strconv.Itoa(i), len: Int.Pow(10, i)}
	}
}

func TestMainFunc(t *testing.T) {
	if !Helper.CompareFuncStdout(main, fmt.Sprintf("Longest substring in %q with max %d distinct characters is:\n\t-> %s", str, limit, longest_substring_with_k_distinct_characters(str, limit))) {
		t.FailNow()
	}
}

func coreTestLongest_substring_with_k_distinct_characters(t *testing.T, fn func(string, int) string) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := fn(tc.input, tc.k)
			if got != tc.want {
				t.Errorf("input: %q, k: %d, expected: %q, got: %q", tc.input, tc.k, tc.want, got)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func TestLongest_substring_with_k_distinct_characters(t *testing.T) {
	coreTestLongest_substring_with_k_distinct_characters(t, longest_substring_with_k_distinct_characters)
}

func TestLongest_substring_with_k_distinct_characters_map(t *testing.T) {
	coreTestLongest_substring_with_k_distinct_characters(t, longest_substring_with_k_distinct_characters_map)
}

func coreBenchmarkLongest_substring_with_k_distinct_characters(b *testing.B, fn func(string, int) string) {
	for k, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			input, result := Str.Random(bm.len), ""
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				fn(input, k)
			}
			resultString = result
		})
	}
}

func BenchmarkLongest_substring_with_k_distinct_characters(b *testing.B) {
	coreBenchmarkLongest_substring_with_k_distinct_characters(b, longest_substring_with_k_distinct_characters)
}

func BenchmarkLongest_substring_with_k_distinct_characters_map(b *testing.B) {
	coreBenchmarkLongest_substring_with_k_distinct_characters(b, longest_substring_with_k_distinct_characters_map)
}
