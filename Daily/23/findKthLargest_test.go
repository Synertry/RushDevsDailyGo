package Daily23

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var (
	resultInt int
	random    = rand.New(rand.NewSource(time.Now().UnixNano()))
)

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
				t.Errorf("expected: %d for target %d, got: %d", tc.want, tc.target, got)
				t.Error(diff)
				t.Fatalf("input: %#v", tc.input)
			}
		})
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	b.ReportAllocs()
	benchmarks := []struct { // do not use maps! Order will be randomized
		name string
		len  int
	}{
		{"ArrLen2", 2},
		{"ArrLen10^1", 10},
		{"ArrLen10^2", 100},
		{"ArrLen10^3", 1000},
		{"ArrLen10^4", 10000},
	}

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
	b.ReportAllocs()
	benchmarks := []struct { // do not use maps! Order will be randomized
		name string
		len  int
	}{
		{"ArrLen2", 2},
		{"ArrLen10^1", 10},
		{"ArrLen10^2", 100},
		{"ArrLen10^3", 1000},
		{"ArrLen10^4", 10000},
	}

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
