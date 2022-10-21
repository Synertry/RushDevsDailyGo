package main

import "math/big"

// isPrime checks if a number is a prime number in O(√n)
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeBaillePSW(num int) bool {
	if num <= 1 {
		return false
	}
	return big.NewInt(int64(num)).ProbablyPrime(0)
}

func isPrimeMillerRabin(num int) bool {
	if num <= 1 {
		return false
	}
	return big.NewInt(int64(num)).ProbablyPrime(1)
}
