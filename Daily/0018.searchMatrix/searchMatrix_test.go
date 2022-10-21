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

const maxExpArrLen = 9

var (
	resultBool bool
	random     = rand.New(rand.NewSource(time.Now().UnixNano()))
	benchmarks = make([]benchmark, maxExpArrLen) // do not use maps! Order will be randomized
)

func init() {
	// generate benchmark data
	for i := 1; i <= maxExpArrLen; i++ {
		arrLen := Math.IntPow(2, i)
		benchmarks[i-1] = benchmark{name: "ArrLen2^" + strconv.Itoa(i), len: arrLen}
	}
}

func TestSearchMatrix(t *testing.T) {
	tests := map[string]struct {
		input  [][]int
		target int
		want   bool
	}{
		"intro":  {input: [][]int{{1, 3, 5, 8}, {10, 11, 15, 16}, {24, 27, 30, 31}}, target: 4, want: false},
		"intro2": {input: [][]int{{1, 3, 5, 8}, {10, 11, 15, 16}, {24, 27, 30, 31}}, target: 10, want: true},
		"single": {input: [][]int{{1}}, target: 1, want: true},
		"empty":  {input: [][]int{}, target: 0, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := searchMatrix(tc.input, tc.target)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("target: %d, expected: %t, got: %t", tc.target, tc.want, got)
				t.Log(diff)
				t.Logf("input: %#v\n", tc.input)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func BenchmarkSearchMatrix(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			random.Seed(time.Now().UnixNano())
			var (
				target = random.Intn(bm.len + 1)
				input  [][]int
				result bool
			)
			for i := 0; i < bm.len; i++ { // generate sorted matrix
				row := make([]int, bm.len/Math.IntPow(2, countDigits(bm.len)))
				for j := 0; j < len(row); j++ { // fill row with ascending numbers from i
					row[j] = i
				}
				input = append(input, row)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = searchMatrix(input, target)
			}
			resultBool = result
		})
	}
}

// countDigits returns the number of digits in a number
func countDigits(num int) (count int) {
	num = abs(num)
	for num > 0 {
		num /= 10
		count++
	}
	return
}

// abs returns the absolute value of a number
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
