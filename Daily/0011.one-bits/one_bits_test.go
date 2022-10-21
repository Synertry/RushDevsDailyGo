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
	for i := 0; i <= maxExp; i++ { // -1 as start, because substraction is more costly than addition
		inputNum := Int.Pow(2, i)
		benchmarks[i] = benchmark{name: "2^" + strconv.Itoa(i), input: inputNum}
	}
}

func TestMainFunc(t *testing.T) {
	want := fmt.Sprintf("Number of set bits in %d is:\n\t-> %d", num, one_bits(num))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestOne_bits(t *testing.T) {
	for i := 0; i < maxTestSize; i++ {
		std := bits.OnesCount64(uint64(i))
		res := one_bits(i)
		if std != res {
			t.Fatalf("input: %d, expected %d, got %d", i, std, res)
		}
	}
}

func TestOne_bitsNonBitOps(t *testing.T) {
	for i := 0; i < maxTestSize; i++ {
		std := bits.OnesCount64(uint64(i))
		res := one_bitsNonBitOps(i)
		if std != res {
			t.Fatalf("input: %d, expected %d, got %d", i, std, res)
		}
	}
}

func TestOne_bitsO1(t *testing.T) {
	for i := 0; i < maxTestSize; i++ {
		std := bits.OnesCount64(uint64(i))
		// res := one_bitsO1(uint64(i))
		res := one_bitsO1(i)
		if std != res {
			t.Fatalf("input: %d, expected %d, got %d", i, std, res)
		}
	}
}

func BenchmarkOne_bits(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cnt, input := 0, bm.input
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cnt = one_bits(input)
			}
			bitCount = cnt
		})
	}
}

func BenchmarkOne_bitsNonBitOps(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cnt, input := 0, bm.input
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cnt = one_bitsNonBitOps(input)
			}
			bitCount = cnt
		})
	}
}

func BenchmarkOne_bitsO1(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cnt, input := 0, bm.input
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cnt = one_bitsO1(input)
			}
			bitCount = cnt
		})
	}
}

func BenchmarkBitsOnesCount(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cnt, input := 0, uint(bm.input)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cnt = bits.OnesCount(input)
			}
			bitCount = cnt
		})
	}
}