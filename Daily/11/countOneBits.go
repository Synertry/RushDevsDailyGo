package Daily11

func CountOneBits(num int) (cnt int) {
	for num != 0 {
		cnt += num & 1
		num >>= 1
	}
	return
}

// CountOneBitsNonBitOps is like countOneBits but without bitwise operation
func CountOneBitsNonBitOps(num int) (cnt int) {
	for num != 0 {
		if num%2 == 1 {
			cnt++
		}
		num /= 2
	}
	return
}

// CountOneBitsO1 is the PopCnt alternative for Go
// Time complexity: O(1)
func CountOneBitsO1(x uint64) int {
	const m0 = 0x5555555555555555 // 01010101 ...
	const m1 = 0x3333333333333333 // 00110011 ...
	const m2 = 0x0f0f0f0f0f0f0f0f // 00001111 ...

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
	x += x >> 32
	return int(x) & (1<<7 - 1)
}
