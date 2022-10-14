package Daily13

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

func TestFindMajorElem(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"intro":  {input: []int{3, 5, 3, 3, 2, 4, 3}, want: 3},
		"same":   {input: []int{2, 2, 2, 2, 2, 2}, want: 2},
		"single": {input: []int{1}, want: 1},
		"empty":  {input: []int{}, want: 0},
	}

	for i := 1; i < 1000; i++ { // random tests
		random.Seed(time.Now().UnixNano()) // ensure pseudo-randomness

		input, majorElem := make([]int, 0, i*2), random.Intn(i)
		for j := 0; j < i; j++ {
			input = append(input, j, majorElem)
		}

		tests["randomLen"+strconv.Itoa(i*2)] = struct {
			input []int
			want  int
		}{input: input, want: majorElem}
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := FindMajorElem(tc.input)
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

func BenchmarkFindMajorElemMixed(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			random.Seed(time.Now().UnixNano())
			input, majorElem, result := make([]int, 0, bm.len*2), random.Intn(bm.len)-1, 0
			for i := 0; i < bm.len; i++ {
				input = append(input, i, majorElem)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = FindMajorElem(input)
			}
			resultInt = result
		})
	}
}

func BenchmarkFindMajorElem1stHalf(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			random.Seed(time.Now().UnixNano())
			input, majorElem, result := make([]int, 0, bm.len*2), random.Intn(bm.len)-1, 0

			for i := 0; i < bm.len; i++ {
				input = append(input, majorElem)
			}
			input = append(input, random.Perm(bm.len)...)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = FindMajorElem(input)
			}
			resultInt = result
		})
	}
}

func BenchmarkFindMajorElem2ndHalf(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			random.Seed(time.Now().UnixNano())
			input, majorElem, result := make([]int, 0, bm.len*2), random.Intn(bm.len)-1, 0

			input = append(input, random.Perm(bm.len)...)
			for i := 0; i < bm.len; i++ {
				input = append(input, majorElem)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = FindMajorElem(input)
			}
			resultInt = result
		})
	}
}