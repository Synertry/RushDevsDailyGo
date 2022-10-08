package Daily23

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var resultInt int

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	tests := map[string]struct {
		input  []int
		target int
		want   int
	}{
		"intro":     {input: []int{8, 7, 2, 3, 4, 1, 5, 9, 0}, target: 3, want: 7},
		"same":      {input: []int{3, 3, 3, 3, 3, 3, 3, 3, 3}, target: 1, want: 3},
		"single":    {input: []int{1}, target: 1, want: 1},
		"sorted":    {input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 4, want: 6},
		"largeDiff": {input: []int{1, 30, 81, 2, 20000, 6}, target: 2, want: 81},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := FindKthLargest(tc.input, tc.target)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
				t.Fatalf(diff)
			}
		})
	}
}

func BenchmarkFindAllConcatenatedWordsInADict(b *testing.B) {
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
			rand.Seed(time.Now().UnixNano())
			input, target, result := rand.Perm(bm.len), rand.Intn(bm.len-1)+1, 0
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = FindKthLargest(input, target)
			}
			resultInt = result
		})
	}
}
