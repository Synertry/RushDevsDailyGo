package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/IO"
	"github.com/Synertry/GoSysUtils/Math/Int"
	"github.com/google/go-cmp/cmp"
)

type benchmark struct {
	name string
	len  int
}

const maxExpArrLen = 7

var (
	resultInt int
	// random     = rand.New(rand.NewSource(time.Now().UnixNano()))
	benchmarks = make([]benchmark, maxExpArrLen+1) // do not use maps! Order will be randomized; + 1 for 2^0
)

func init() {
	// generate benchmark data
	for i := 0; i <= maxExpArrLen; i++ {
		arrLen := Int.Pow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}
}

func TestMainFunc(t *testing.T) {
	copySlice := make([]int, len(sliceOfInts))
	copy(copySlice, sliceOfInts)
	want := fmt.Sprintf("Length without dupes of %v is:\n\t-> %d", copySlice, remove_dups(copySlice))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestRemove_dups(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"intro":  {input: []int{1, 1, 2, 3, 4, 4, 4, 4, 4, 5, 5, 6, 7, 9}, want: 8},
		"same":   {input: []int{1, 1, 1, 1, 1, 1}, want: 1},
		"single": {input: []int{0}, want: 1},
		"empty":  {input: []int{}, want: 1},
	}

	// runs too long for coverage report
	// for i := 1; i < 1000; i++ { // random tests
	// 	random.Seed(time.Now().UnixNano()) // ensure pseudo-randomness
	// 	var input []int

	// 	for k := 0; k < i; k++ {
	// 		input = append(input, k)
	// 		for j, dupeCount := 0, random.Intn(9)+1; j < dupeCount; j++ {
	// 			input = append(input, k)
	// 		}
	// 	}

	// 	tests["randomMinLen"+strconv.Itoa(i)] = struct {
	// 		input []int
	// 		want  int
	// 	}{input: input, want: i}
	// }

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := remove_dups(tc.input)
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

func BenchmarkRemove_dups(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			input, result := make([]int, 0, bm.len*2), 0
			for i := 0; i < bm.len; i++ {
				input = append(input, i, i)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = remove_dups(input)
			}
			resultInt = result
		})
	}
}