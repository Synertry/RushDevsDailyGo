package Daily21

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

// TODO: Implement fuzz testing or random string slice generator
