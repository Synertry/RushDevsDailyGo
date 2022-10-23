package main

import (
	"fmt"
	"testing"

	"github.com/Synertry/GoSysUtils/IO"
)

var (
	benchmarks = []struct {
		name string
		num  int
	}{
		{name: "6", num: 6},
		{name: "28", num: 28},
		{name: "496", num: 496},
		{name: "8128", num: 8128},
		{name: "33550336", num: 33550336},
	}

	tests = map[string]struct {
		input int
		want  bool
	}{
		"perfect6":        {input: 6, want: true},
		"perfect28":       {input: 28, want: true},
		"perfect496":      {input: 496, want: true},
		"perfect8128":     {input: 8128, want: true},
		"perfect33550336": {input: 33550336, want: true},
		"notPerfect":      {input: 1, want: false},
		"zero":            {input: 0, want: false},
		"negative":        {input: -1, want: false},
	}

	resultBool bool
)

// ##### Tests #####

func TestMainFunc(t *testing.T) {
	want := fmt.Sprintf("Number %d is a perfect number:\n\t-> %t", num, checkPerfectNumber(num))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func testCheckPerfectNumber(t *testing.T, fn func(int) bool) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := fn(tc.input)
			if got != tc.want {
				t.Errorf("input: %d, expected: %t, got: %t", tc.input, tc.want, got)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func TestCheckPerfectNumberSqrt(t *testing.T) {
	testCheckPerfectNumber(t, checkPerfectNumberSqrt)
}

func TestCheckPerfectNumber(t *testing.T) {
	testCheckPerfectNumber(t, checkPerfectNumber)
}

// ##### Benchmarks #####

func benchmarkCheckPerfectNumber(b *testing.B, fn func(int) bool) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			var result bool
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = fn(bm.num)
			}
			resultBool = result
		})
	}
}

func BenchmarkCheckPerfectNumberSqrt(b *testing.B) {
	benchmarkCheckPerfectNumber(b, checkPerfectNumberSqrt)
}

func BenchmarkCheckPerfectNumber(b *testing.B) {
	benchmarkCheckPerfectNumber(b, checkPerfectNumber)
}
