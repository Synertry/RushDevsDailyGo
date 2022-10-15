package Daily31

import (
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/Math"
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
		inputNum := Math.IntPow(10, i)
		benchmarksPrimes[i] = benchmarksPrime{name: "10^" + strconv.Itoa(i), len: inputNum}
	}
}

func TestIsPrime(t *testing.T) {
	for name, tcP := range testsPrimes {
		t.Run(name, func(t *testing.T) {
			if got := isPrime(tcP.input); got != tcP.want {
				t.Errorf("input: %d, expected: %t, got: %t", tcP.input, tcP.want, got)
			}
		})
	}
}

func TestIsPrimeBaillePSW(t *testing.T) {
	for name, tcP := range testsPrimes {
		t.Run(name, func(t *testing.T) {
			if got := isPrime(tcP.input); got != tcP.want {
				t.Errorf("input: %d, expected: %t, got: %t", tcP.input, tcP.want, got)
			}
		})
	}
}

func TestIsPrimeMillerRabin(t *testing.T) {
	for name, tcP := range testsPrimes {
		t.Run(name, func(t *testing.T) {
			if got := isPrime(tcP.input); got != tcP.want {
				t.Errorf("input: %d, expected: %t, got: %t", tcP.input, tcP.want, got)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for _, bmP := range benchmarksPrimes {
		b.Run(bmP.name, func(b *testing.B) {
			var result bool
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = isPrime(bmP.len)
			}
			resultBool = result
		})
	}
}

func BenchmarkIsPrimeBaillePSW(b *testing.B) {
	for _, bmP := range benchmarksPrimes {
		b.Run(bmP.name, func(b *testing.B) {
			var result bool
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = isPrimeBaillePSW(bmP.len)
			}
			resultBool = result
		})
	}
}

func BenchmarkIsPrimeMillerRabin(b *testing.B) {
	for _, bmP := range benchmarksPrimes {
		b.Run(bmP.name, func(b *testing.B) {
			var result bool
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = isPrimeMillerRabin(bmP.len)
			}
			resultBool = result
		})
	}
}
