package Daily24

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var resultInts []int

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
			got := Closest_3sum(tc.input, tc.target)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
				t.Fatalf(diff)
			}
		})
	}
}

func BenchmarkClosest_3sum(b *testing.B) {
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
			input, target, result := rand.Perm(bm.len), rand.Intn(bm.len-1)+1, make([]int, 3)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = Closest_3sum(input, target)
			}
			resultInts = result
		})
	}
}
