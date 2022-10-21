package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/IO"
	"github.com/Synertry/GoSysUtils/Math/Int"
)

var resultString string

func TestMainFunc(t *testing.T) {
	want := fmt.Sprintf("Fraction to decimal of %d/%d is:\n\t-> %s", nominator, denominator, frac_to_dec(nominator, denominator))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestFrac_To_Dec(t *testing.T) {
	tests := map[string]struct {
		nominator   int
		denominator int
		want        string
	}{
		"intro":  {nominator: -3, denominator: 2, want: "-1.5"},
		"intro2": {nominator: 4, denominator: 3, want: "1.(3)"},
		"intro3": {nominator: 1, denominator: 6, want: "0.1(6)"},
		"1/2":    {nominator: 1, denominator: 2, want: "0.5"},
		"2/1":    {nominator: 2, denominator: 1, want: "2"},
		"2/3":    {nominator: 2, denominator: 3, want: "0.(6)"},
		"4/333":  {nominator: 4, denominator: 333, want: "0.(012)"},
		"1/5":    {nominator: 1, denominator: 5, want: "0.2"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := frac_to_dec(tc.nominator, tc.denominator)
			if got != tc.want {
				t.Errorf("input: %d/%d, expected: %s, got: %s", tc.nominator, tc.denominator, tc.want, got)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func BenchmarkFrac_To_Dec(b *testing.B) {
	type benchmark struct {
		name  string
		input int
	}

	const maxExp = 10

	benchmarks := make([]benchmark, maxExp+1)

	// generate benchmark data
	for i := 0; i <= maxExp; i++ {
		inputNum := Int.Pow(2, i)
		benchmarks[i] = benchmark{name: "1÷" + strconv.Itoa(inputNum), input: inputNum}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			inputn, inputd, result := 1, bm.input, ""
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = frac_to_dec(inputn, inputd)
			}
			resultString = result
		})
	}
}
