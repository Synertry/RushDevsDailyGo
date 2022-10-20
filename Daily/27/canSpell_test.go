package Daily27

import (
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/Math"
	"github.com/Synertry/GoSysUtils/Slice"
	"github.com/Synertry/GoSysUtils/Str"
)

var resultBool bool

func TestCanSpell(t *testing.T) {
	tests := map[string]struct {
		magazine []string
		note     string
		want     bool
	}{
		"intro":       {magazine: []string{"a", "b", "c", "d", "e", "f"}, note: "bed", want: true},
		"intro2":      {magazine: []string{"a", "b", "c", "d", "e", "f"}, note: "cat", want: false},
		"explanation": {magazine: []string{"s", "w", "i", "m", "i", "n", "g"}, note: "swimming", want: false},
		"same":        {magazine: []string{"a", "a", "a"}, note: "a", want: true},
		"empty":       {magazine: []string{}, note: "", want: true},
		"emptyf":      {magazine: []string{}, note: "f", want: false},
		"emptyt":      {magazine: []string{"a"}, note: "", want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CanSpell(tc.magazine, tc.note)
			if got != tc.want {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func BenchmarkCanSpell(b *testing.B) {
	type benchmark struct {
		name string
		len  int
	}

	maxExpDictLen := 4
	benchmarks := make([]benchmark, maxExpDictLen+2) // + 2 for empty floor(10^-1) and single 10^0 -> 1

	for i := -1; i <= maxExpDictLen; i++ { // -1 as start, because substraction is more costly than addition
		dictLen := Math.IntPow(10, i)
		benchmarks[i+1] = benchmark{name: "DictLen10^" + strconv.Itoa(i), len: dictLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			input, word, result := Slice.GenRandomStringsLen(bm.len, 1), Str.GenRandom(5), false
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = CanSpell(input, word)
			}
			resultBool = result
		})
	}
}
