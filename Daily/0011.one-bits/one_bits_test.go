package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/IO"
	"github.com/Synertry/GoSysUtils/Math/Int"
)

type benchmark struct {
	name  string
	input int
}

const maxExp = 6

var (
	bitCount    int
	maxTestSize = Int.Pow(10, maxExp)
	benchmarks  = make([]benchmark, maxExp+1) // do not use maps! Order will be randomized; + 1 for 2^0
)

func init() {
	// generate benchmark data
	for i := 0; i <= maxExp; i++ {
		inputNum := Int.Pow(2, i)
		benchmarks[i] = benchmark{name: "2^" + strconv.Itoa(i), input: inputNum}
	}
}

// ##### Tests #####

func TestMainFunc(t *testing.T) {
	want := fmt.Sprintf("Number of set bits in %d is:\n\t-> %d", num, one_bits(num))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func coreTestOne_bits(t *testing.T, fn func(int) int) {
	for i := 0; i < maxTestSize; i++ {
		std := bits.OnesCount64(uint64(i))
		res := fn(i)
		if std != res {
			t.Fatalf("input: %d, expected %d, got %d", i, std, res)
		}
	}
}

func TestOne_bits(t *testing.T) {
	coreTestOne_bits(t, one_bits)
}

func TestOne_bitsNonBitOps(t *testing.T) {
	coreTestOne_bits(t, one_bitsNonBitOps)
}

func TestOne_bitsO1(t *testing.T) {
	coreTestOne_bits(t, one_bitsO1)
}

// ##### Benchmarks #####

func coreBenchmarkOne_bits(b *testing.B, fn func(int) int) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cnt, input := 0, bm.input
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cnt = fn(input)
			}
			bitCount = cnt
		})
	}
}

func BenchmarkOne_bits(b *testing.B) {
	coreBenchmarkOne_bits(b, one_bits)
}

func BenchmarkOne_bitsNonBitOps(b *testing.B) {
	coreBenchmarkOne_bits(b, one_bitsNonBitOps)
}

func BenchmarkOne_bitsO1(b *testing.B) {
	coreBenchmarkOne_bits(b, one_bitsO1)
}

func BenchmarkBitsOnesCount(b *testing.B) {
	bitsOnesCount := func(num int) int {
		return bits.OnesCount(uint(num))
	}
	coreBenchmarkOne_bits(b, bitsOnesCount)
}
