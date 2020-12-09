package codility

import (
	"fmt"
	"math"
	"strconv"
)

// BinaryGap - find the binary gap of an int
func BinaryGap(N int) int {
	var result int
	var resultTemp int
	var firstOne bool

	for N > 0 {
		binrep := strconv.FormatInt(int64(N), 2)
		fmt.Println(binrep)

		// N is odd
		if N%2 == 1 {
			// Don't count till the first '1' is seen
			if !firstOne {
				firstOne = true
				// Reset count every '1' thereafter
			} else {
				if resultTemp > result {
					result = resultTemp
				}
				resultTemp = 0
			}
			// N is even, thus a '0' at the end in binary
		} else {
			if firstOne {
				resultTemp++
			}
		}
		// powers of 2
		N = N / 2
	}

	return result
}

// Rotate an array
func Rotate(A []int, K int) []int {
	// Perform K rotations to the right on input

	// The easy ones
	if A == nil || len(A) == 0 || len(A) == 1 {
		return A
	}

	// Only need to rotate what's left after a full rotation
	k := K % len(A)
	if len(A) == K || k == 0 {
		return A
	}

	a := make([]int, len(A))

	// slice the last k items
	if k > len(A) {
		panic("Can't have a negative index")
	}
	backK := A[len(A)-k:]

	// Shift by k
	for i := 0; (i + k) < len(A); i++ {
		a[i+k] = A[i]
	}

	// Insert the saved k elements at the front
	for i := 0; i < k; i++ {
		a[i] = backK[i]
	}

	return a
}

// OddManOut - find the value of the element that does not have a partner
func OddManOut(A []int) int {

	oddMan := 0

	for _, currentVal := range A {
		oddMan ^= currentVal
	}

	return oddMan
}

// FrogJump - Count the minimal number of jumps that the small frog must perform to reach its target.
func FrogJump(X int, Y int, D int) int {
	startPoint := float64(Y - X)
	distance := float64(D)
	return int(math.Ceil(startPoint / distance))
}
