package Daily11

import (
	"math/bits"
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/Math"
)

type benchmark struct {
	name  string
	input int
}

const maxExp = 8

var (
	bitCount       int
	bitCountUint64 uint64
	maxTestSize    = Math.IntPow(10, maxExp)
	benchmarks     = make([]benchmark, maxExp+1) // do not use maps! Order will be randomized; + 1 for 2^0
)

func init() {
	// generate benchmark data
	for i := 0; i <= maxExp; i++ { // -1 as start, because substraction is more costly than addition
		inputNum := Math.IntPow(2, i)
		benchmarks[i] = benchmark{name: "2^" + strconv.Itoa(i), input: inputNum}
	}
}

func TestCountOneBits(t *testing.T) {
	for i := 0; i < maxTestSize; i++ {
		std := bits.OnesCount64(uint64(i))
		res := CountOneBits(i)
		if std != res {
			t.Fatalf("input: %d, expected %d, got %d", i, std, res)
		}
	}
}

func TestCountOneBitsNonBitOps(t *testing.T) {
	for i := 0; i < maxTestSize; i++ {
		std := bits.OnesCount64(uint64(i))
		res := CountOneBitsNonBitOps(i)
		if std != res {
			t.Fatalf("input: %d, expected %d, got %d", i, std, res)
		}
	}
}

func TestCountOneBitsO1(t *testing.T) {
	for i := 0; i < maxTestSize; i++ {
		std := bits.OnesCount64(uint64(i))
		res := CountOneBitsO1(uint64(i))
		if std != res {
			t.Fatalf("input: %d, expected %d, got %d", i, std, res)
		}
	}
}

func BenchmarkCountOneBits(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cnt, input := 0, bm.input
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cnt = CountOneBits(input)
			}
			bitCount = cnt
		})
	}
}

func BenchmarkCountOneBitsNonBitOps(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cnt, input := 0, bm.input
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cnt = CountOneBitsNonBitOps(input)
			}
			bitCount = cnt
		})
	}
}

func BenchmarkCountOneBitsO1(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cnt, input := 0, uint64(bm.input)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cnt = CountOneBitsO1(input)
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
