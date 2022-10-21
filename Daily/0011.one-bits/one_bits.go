package main

import "fmt"

var num = 23

func main() {
	fmt.Printf("Number of set bits in %d is:\n\t-> %d", num, one_bits(num))
	// Add line break in above print if you uncomment next lines
	// fmt.Printf("Number of set bits in %d is (without bitwise operation):\n\t-> %d\n\n", num, one_bitsNonBitOps(num))
	// fmt.Printf("Number of set bits in %d is (O(1)):\n\t-> %d\n\n", num, one_bitsO1(num))
}

func one_bits(num int) (cnt int) {
	for num != 0 {
		cnt += num & 1
		num >>= 1
	}
	return
}

// one_bitsNonBitOps is like one_bits but without bitwise operation
func one_bitsNonBitOps(num int) (cnt int) {
	for num != 0 {
		if num%2 == 1 {
			cnt++
		}
		num /= 2
	}
	return
}

// one_bitsO1 is the PopCnt alternative for Go
// Time complexity: O(1)
func one_bitsO1(x int) int {
	const m0 = 0x55555555 // 01010101 ...
	const m1 = 0x33333333 // 00110011 ...
	const m2 = 0x0f0f0f0f // 00001111 ...

	// Implementation: Parallel summing of adjacent bits.
	// See "Hacker's Delight", Chap. 5: Counting Bits.
	//
	// Masking (& operations) can be omitted when the field's sum won't carry over into the next field
	// Since the result cannot be > 64, 8 bits is enough, and we can ignore the masks for the shifts by 8 and up.

	const m = 1<<64 - 1
	x = x - x>>1&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	return x & (1<<7 - 1)
}