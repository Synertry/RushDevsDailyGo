package Daily23

import (
	"math/rand"
	"sort"
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

const maxExpArrLen = 4

var (
	resultInt  int
	random     = rand.New(rand.NewSource(time.Now().UnixNano()))
	benchmarks = make([]benchmark, maxExpArrLen+1) // do not use maps! Order will be randomized; + 1 for 2^0
)

func init() {
	// generate benchmark data
	benchmarks[0] = benchmark{name: "ArrLen2", len: 2} // start case

	for i := 1; i <= maxExpArrLen; i++ {
		arrLen := Math.IntPow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}
}

func TestFindKthLargest(t *testing.T) {
	tests := map[string]struct {
		input  []int
		target int
		want   int
	}{
		"intro":         {input: []int{8, 7, 2, 3, 4, 1, 5, 9, 0}, target: 3, want: 7},
		"same":          {input: []int{3, 3, 3, 3, 3, 3, 3, 3, 3}, target: 1, want: 3},
		"single":        {input: []int{1}, target: 1, want: 1},
		"empty":         {input: []int{}, target: 0, want: 0},
		"sorted":        {input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 1, want: 9},
		"sortedReverse": {input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, target: 1, want: 9},
		"sortedKHigh":   {input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 10, want: 0},
		"largeDiff":     {input: []int{1, 30, 81, 2, 20000, 6}, target: 2, want: 81},
	}

	for i := 2; i < 1000; i++ { // random tests
		random.Seed(time.Now().UnixNano()) // ensure pseudo-randomness
		input, target := random.Perm(i), random.Intn(i-1)+1
		tests["randomLen"+strconv.Itoa(i+1)] = struct {
			input  []int
			target int
			want   int
		}{input: input, target: target, want: simpleSolution(input, target)}
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := FindKthLargest(tc.input, tc.target)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("target, %d, expected: %d, got: %d", tc.target, tc.want, got)
				t.Log(diff)
				t.Logf("input: %#v\n", tc.input)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			input, target, result := random.Perm(bm.len), random.Intn(bm.len-1)+1, 0
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = FindKthLargest(input, target)
			}
			resultInt = result
		})
	}
}

func BenchmarkFindKthLargestSimple(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			input, target, result := random.Perm(bm.len), random.Intn(bm.len-1)+1, 0
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = simpleSolution(input, target)
			}
			resultInt = result
		})
	}
}

func simpleSolution(arr []int, k int) int {
	sort.Ints(arr)
	return arr[len(arr)-k]
}
