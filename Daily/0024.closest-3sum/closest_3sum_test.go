package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Synertry/GoSysUtils/IO"
	"github.com/Synertry/GoSysUtils/Math/Int"
	"github.com/google/go-cmp/cmp"
)

var (
	resultSliceOfInts []int
	random            = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func TestMainFunc(t *testing.T) {
	want := fmt.Sprintf("Closest 3sum to %d in %v is:\n\t-> %v", num, sliceOfInts, closest_3sum(sliceOfInts, num))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestClosest_3sum(t *testing.T) {
	tests := map[string]struct {
		input  []int
		target int
		want   []int
	}{
		"intro":   {input: []int{2, 1, -5, 4}, target: -1, want: []int{-5, 1, 2}},
		"0":       {input: []int{0, 0, 0}, target: 1, want: []int{0, 0, 0}},
		"random1": {input: []int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4}, target: 1, want: []int{-4, 2, 2}},
		"random2": {input: []int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4}, target: 0, want: []int{-4, 2, 2}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := closest_3sum(tc.input, tc.target)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
				t.Fatalf(diff)
			}
		})
	}
}

func BenchmarkClosest_3sum(b *testing.B) {
	maxExpArrLen := 4
	type benchmark struct {
		name string
		len  int
	}

	benchmarks := make([]benchmark, maxExpArrLen+1)    // do not use maps! Order will be randomized; + 1 for 2^0
	benchmarks[0] = benchmark{name: "ArrLen2", len: 2} // start case

	for i := 1; i <= maxExpArrLen; i++ {
		arrLen := Int.Pow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			random.Seed(time.Now().UnixNano())
			input, target, result := random.Perm(bm.len), random.Intn(bm.len-1)+1, make([]int, 3)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = closest_3sum(input, target)
			}
			resultSliceOfInts = result
		})
	}
}