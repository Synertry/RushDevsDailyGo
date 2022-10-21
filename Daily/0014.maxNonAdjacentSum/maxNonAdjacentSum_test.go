package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Synertry/GoSysUtils/Math"
	"github.com/google/go-cmp/cmp"
)

type benchmark struct {
	name string
	len  int
}

const maxExpArrLen = 6

var (
	resultInt  int
	random     = rand.New(rand.NewSource(time.Now().UnixNano()))
	benchmarks = make([]benchmark, maxExpArrLen+1) // do not use maps! Order will be randomized; + 1 for 2^0
)

func init() {
	// generate benchmark data
	for i := 0; i <= maxExpArrLen; i++ {
		arrLen := Math.IntPow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}
}

func TestMaxNonAdjacentSum(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"intro":  {input: []int{3, 4, 1, 1}, want: 5},
		"intro2": {input: []int{2, 1, 2, 7, 3}, want: 9},
		"single": {input: []int{1}, want: 1},
		"empty":  {input: []int{}, want: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := maxNonAdjacentSum(tc.input)
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

func BenchmarkMaxNonAdjacentSum(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			random.Seed(time.Now().UnixNano())
			input, result := random.Perm(bm.len), 0
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = maxNonAdjacentSum(input)
			}
			resultInt = result
		})
	}
}
