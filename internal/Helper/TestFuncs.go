package Helper

import (
	"fmt"

	"github.com/Synertry/GoSysUtils/IO"
)

// CompareFuncStdout compares the output of a function to a string
func CompareFuncStdout(main func(), want string) bool {
	got := IO.GetOutput(main)
	if got != want {
		fmt.Printf("expected: %s, got: %s", want, got)
		return false
	}
	return true
}