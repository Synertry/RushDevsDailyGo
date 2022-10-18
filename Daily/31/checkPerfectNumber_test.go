package Daily31

import "testing"

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

func TestCheckPerfectNumberSqrt(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CheckPerfectNumberSqrt(tc.input)
			if got != tc.want {
				t.Errorf("input: %d, expected: %t, got: %t", tc.input, tc.want, got)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func TestCheckPerfectNumberSqrtLog(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CheckPerfectNumberSqrtLog(tc.input)
			if got != tc.want {
				t.Errorf("input: %d, expected: %t, got: %t", tc.input, tc.want, got)
			}
		})
		if t.Failed() {
			break
		}
	}
}

// ##### Benchmarks #####

func BenchmarkCheckPerfectNumberSqrt(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			var result bool
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = CheckPerfectNumberSqrt(bm.num)
			}
			resultBool = result
		})
	}
}

func BenchmarkCheckPerfectNumberSqrtLog(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			var result bool
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = CheckPerfectNumberSqrtLog(bm.num)
			}
			resultBool = result
		})
	}
}
