package codility

// BinaryGap - find the binary gap of an int
func BinaryGap(N int) int {
	var result int
	var resultTemp int
	var calc bool

	for N > 0 {
		// N is odd
		if N%2 == 1 {
			if !calc {
				calc = true
			} else {
				if resultTemp > result {
					result = resultTemp
				}
				resultTemp = 0
			}
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
