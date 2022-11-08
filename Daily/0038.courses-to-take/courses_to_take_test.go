package main

import (
	"fmt"
	"testing"

	"github.com/Synertry/GoSysUtils/IO"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var (
	// resultSliceOfStrings []string
	less = func(a, b string) bool { return a < b } // sorts slice
)

func TestMainFunc(t *testing.T) {
	want := fmt.Sprintf("For the courses %v the recommended order is:\n\t-> %v", mapOfSlicesOfStrings, courses_to_take(mapOfSlicesOfStrings))
	got := IO.GetOutput(main)
	if got != want {
		t.Errorf("expected:\n%s\n, got:\n%s", want, got) // <- easier debugging with distinct line breaks
	}
}

func TestCourses_to_take(t *testing.T) {
	tests := map[string]struct {
		input map[string][]string
		want  []string
	}{
		"intro": {input: map[string][]string{
			"CSC300": {"CSC100", "CSC200"},
			"CSC200": {"CSC100"},
			"CSC100": {},
		}, want: []string{"CSC100", "CSC200", "CSC300"}},
		"empty":        {input: map[string][]string{}, want: []string{}},
		"single":       {input: map[string][]string{"CSC100": {}}, want: []string{"CSC100"}},
		"notAvailable": {input: map[string][]string{"CSC100": {"CSC200"}}, want: nil},
		"reqCycle":     {input: map[string][]string{"CSC100": {"CSC200"}, "CSC200": {"CSC100"}}, want: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := courses_to_take(tc.input)
			diff := cmp.Diff(tc.want, got, cmpopts.SortSlices(less))
			if diff != "" {
				t.Errorf("expected: %v, got: %v", tc.want, got)
				t.Log(diff)
				t.Logf("input: %v\n", tc.input)
			}
		})
		if t.Failed() {
			break
		}
	}
}

// TODO: add benchmark by generating semi-random map of slices of strings, which strings reference the keys of the map
// like how would that even work?
// I would need to generate a slice of strings for the keys of the map.
// and for each key reference other keys in the map chosen randomly with random length.
// Sounds harder than my non-existing matrix generator
