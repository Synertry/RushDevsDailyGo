package main

import (
	"fmt"
	"strconv"
)

var nominator = 1
var denominator = 6

func main() {
	fmt.Printf("Fraction to decimal of %d/%d is:\n\t-> %s", nominator, denominator, frac_to_dec(nominator, denominator))
}

func frac_to_dec(n int, d int) string {
	if n == 0 {
		return "0"
	}

	var intSign string // prefix sign
	if n*d < 0 {
		intSign = "-"
	}

	// easier string handling without integer sign
	n, d = abs(n), abs(d)

	if n >= d {
		ds := frac_to_dec(n%d, d) // recurse the following decimals
		return intSign + strconv.Itoa(n/d) + ds[1:]
	}

	// fraction part
	digits := make([]byte, 2, 1024) // digits stores the result of n/d
	digits[0] = '0'
	digits[1] = '.'
	idx := 2                       // index number of the result of n/d in digits
	rec := make(map[int]int, 1024) // rec[n] = idx
	for {
		if i, ok := rec[n]; ok {
			// if n is repeated, then n/d is the start of the next loop
			// starting point of the loop section is the last occurrence of idx value of n
			return fmt.Sprintf("%s%s(%s)", intSign, string(digits[:i]), string(digits[i:]))
		}

		rec[n] = idx

		n *= 10
		idx++

		digits = append(digits, byte(n/d)+'0')
		n %= d

		if n == 0 {
			// loop break
			return intSign + string(digits)
		}
	}
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
