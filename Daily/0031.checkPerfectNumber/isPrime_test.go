package main

import (
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/Math/Int"
)

type benchmarksPrime struct {
	name string
	len  int
}

const maxArrExpLen = 10

var (
	benchmarksPrimes = make([]benchmarksPrime, maxArrExpLen+1)
	testsPrimes      = map[string]struct {
		input int
		want  bool
	}{
		"2":   {input: 2, want: true},
		"7":   {input: 7, want: true},
		"17":  {input: 17, want: true},
		"18":  {input: 18, want: false},
		"100": {input: 100, want: false},
		"0":   {input: 0, want: false},
		"1":   {input: 1, want: false},
	}
)

func init() {
	for i := 0; i <= maxArrExpLen; i++ {
		inputNum := Int.Pow(10, i)
		benchmarksPrimes[i] = benchmarksPrime{name: "10^" + strconv.Itoa(i), len: inputNum}
	}
}

func coreTestIsPrime(t *testing.T, fn func(int) bool) {
	for name, tcP := range testsPrimes {
		t.Run(name, func(t *testing.T) {
			if got := fn(tcP.input); got != tcP.want {
				t.Errorf("input: %d, expected: %t, got: %t", tcP.input, tcP.want, got)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	coreTestIsPrime(t, isPrime)
}

func TestIsPrimeBaillePSW(t *testing.T) {
	coreTestIsPrime(t, isPrimeBaillePSW)
}

func TestIsPrimeMillerRabin(t *testing.T) {
	coreTestIsPrime(t, isPrimeMillerRabin)
}

func coreBenchmarkIsPrime(b *testing.B, fn func(int) bool) {
	for _, bmP := range benchmarksPrimes {
		b.Run(bmP.name, func(b *testing.B) {
			var result bool
			b.ResetTimer()
			for i := 0; i < 100; i++ {
				result = fn(bmP.len)
			}
			resultBool = result
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	coreBenchmarkIsPrime(b, isPrime)
}

func BenchmarkIsPrimeBaillePSW(b *testing.B) {
	coreBenchmarkIsPrime(b, isPrimeBaillePSW)
}

func BenchmarkIsPrimeMillerRabin(b *testing.B) {
	coreBenchmarkIsPrime(b, isPrimeMillerRabin)
}
