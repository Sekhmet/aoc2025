package toolkit

// Pmod computes the positive modulo of x by d.
func Pmod(x, d int) int {
	x = x % d

	if x < 0 {
		return x + d
	}

	return x
}

// FindMaxIndex returns index of biggest element in values
func FindMaxIndex(values []int) int {
	max_index := 0

	for i, v := range values {
		if v > values[max_index] {
			max_index = i
		}
	}

	return max_index
}
