package codility

import (
	"fmt"
	"strconv"
)

// BinaryGap - find the binary gap of an int
func BinaryGap(N int) int {
	var result int
	var resultTemp int
	var calc bool

	for N > 0 {
		binrep := strconv.FormatInt(int64(N), 2)
		fmt.Println(binrep)

		// N is odd
		if N%2 == 1 {
			// Don't count till the first '1' is seen
			if !calc {
				calc = true
				// Reset count every '1' thereafter
			} else {
				if resultTemp > result {
					result = resultTemp
				}
				resultTemp = 0
			}
			// N is even, thus a '0' at the end in binary
		} else {
			if calc {
				resultTemp++
			}
		}
		// powers of 2
		N = N / 2
	}

	return result
}
