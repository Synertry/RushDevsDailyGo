package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/IO"
	"github.com/Synertry/GoSysUtils/Math/Int"
	"github.com/google/go-cmp/cmp"
)

var resultSliceOfInts []int

func TestMainFunc(t *testing.T) {
	copySlice := make([]int, len(sliceOfInts))
	copy(copySlice, sliceOfInts)
	want := fmt.Sprintf("The list %#v sorted is:\n\t-> %#v", copySlice, sortNums(copySlice))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected:\n%s\n, got:\n%s", want, got) // <- easier debugging with distinct line breaks
	}
}

func TestSortNums(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"intro":         {input: []int{3, 3, 2, 1, 3, 2, 1}, want: []int{1, 1, 2, 2, 3, 3, 3}},
		"same":          {input: []int{3, 3, 3, 3, 3, 3, 3, 3, 3}, want: []int{3, 3, 3, 3, 3, 3, 3, 3, 3}},
		"single":        {input: []int{1}, want: []int{1}},
		"empty":         {input: []int{}, want: []int{}},
		"sorted":        {input: []int{1, 1, 1, 2, 2, 2, 3, 3, 3}, want: []int{1, 1, 1, 2, 2, 2, 3, 3, 3}},
		"sortedReverse": {input: []int{3, 3, 3, 2, 2, 2, 1, 1, 1}, want: []int{1, 1, 1, 2, 2, 2, 3, 3, 3}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sortNums(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("expected: %d, got: %d", tc.want, got)
				t.Log(diff)
				t.Logf("input: %#v\n", tc.input)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func BenchmarkSortNums(b *testing.B) {
	maxExpArrLen := 6
	type benchmark struct {
		name string
		len  int
	}
	benchmarks := make([]benchmark, maxExpArrLen+1) // do not use maps! Order will be randomized; + 1 for 2^0

	for i := 0; i <= maxExpArrLen; i++ {
		arrLen := Int.Pow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			input, result := make([]int, bm.len), make([]int, bm.len)
			for i := 0; i < bm.len; i++ { // fill input with alternating pattern
				input[i] = i%3 + 1
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sortNums(input)
			}
			resultSliceOfInts = result
		})
	}
}
