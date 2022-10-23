package main

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/Synertry/GoSysUtils/IO"
	"github.com/Synertry/GoSysUtils/Math"
	"github.com/Synertry/GoSysUtils/Math/Int"
	"github.com/google/go-cmp/cmp"
)

type benchmark struct {
	name string
	len  int
}

const maxExpArrLen = 9

var (
	resultSliceOfInts []int
	random            = Math.GetRand()
	benchmarks        = make([]benchmark, maxExpArrLen) // do not use maps! Order will be randomized
)

func init() {
	// generate benchmark data
	for i := 1; i <= maxExpArrLen; i++ {
		arrLen := Int.Pow(2, i)
		benchmarks[i-1] = benchmark{name: "ArrLen2^" + strconv.Itoa(i), len: arrLen}
	}
}

func TestMainFunc(t *testing.T) {
	want := fmt.Sprintf("Intersections of %#v are:\n\t-> %v", mat, intersection(mat))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestIntersection(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  []int
	}{
		"intro":  {input: [][]int{{1, 2, 3, 4}, {2, 4, 6, 8}, {3, 4, 5}}, want: []int{4}},
		"intro2": {input: [][]int{{1, 2, 3, 4}, {1, 2, 4, 6, 8}, {1, 3, 4, 5}}, want: []int{1, 4}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := intersection(tc.input)
			sort.Ints(got)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
				t.Log(diff)
				t.Logf("input: %#v\n", tc.input)
			}
		})
		if t.Failed() {
			break
		}
	}
}

// This will not generate intersecting arrays for this daily
// It is taken from Daily\18\searchMatrix_test.go
// TODO: make it generate intersecting arrays
func BenchmarkIntersection(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			random.Seed(time.Now().UnixNano())
			var (
				input  [][]int
				result []int
			)
			for i := 0; i < bm.len; i++ { // generate sorted matrix
				row := make([]int, bm.len/Int.Pow(2, Int.CountDigits(bm.len)))
				for j := 0; j < len(row); j++ { // fill row with ascending numbers from i
					row[j] = i
				}
				input = append(input, row)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = intersection(input)
			}
			resultSliceOfInts = result
		})
	}
}
