package Daily21

import (
	"math/rand"
	"testing"
	"time"
	"unsafe"

	"github.com/google/go-cmp/cmp"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	sliceOfString []string
	src           = rand.NewSource(time.Now().UnixNano())
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []string
	}{
		"intro":     {input: []string{"tech", "lead", "techlead", "cat", "cats", "dog", "catsdog"}, want: []string{"techlead", "catsdog"}},
		"letters":   {input: []string{"a", "b", "ab", "abd"}, want: []string{"ab"}},
		"noConcats": {input: []string{"a", "ab", "c"}, want: []string{}},
		"empty":     {input: []string{}, want: []string{}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := FindAllConcatenatedWordsInADict(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
				t.Fatalf(diff)
			}
		})
	}
}

func BenchmarkFindAllConcatenatedWordsInADict(b *testing.B) {
	b.ReportAllocs()
	benchmarks := []struct { // do not use maps! Order will be randomized
		name string
		len  int
	}{
		{"DictLen0", 0},
		{"DictLen1", 1},
		{"DictLen10", 10},
		{"DictLen100", 100},
		{"DictLen1000", 1000},
		{"DictLen10000", 10000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			input, result := generateRandomStringSlice(bm.len), make([]string, bm.len)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = FindAllConcatenatedWordsInADict(input)
			}
			sliceOfString = result
		})
	}
}

func generateRandomStringSlice(length int) []string {
	slice := make([]string, length)
	for i := 0; i < length; i++ {
		slice[i] = RandStringBytesMaskImprSrcUnsafe((i + 1) / 2)
	}
	return slice
}

// RandStringBytesMaskImprSrcUnsafe generates a random string of length n
// Source: https://stackoverflow.com/a/31832326/5516320
func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
