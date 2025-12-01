package toolkit

// Pmod computes the positive modulo of x by d.
func Pmod(x, d int) int {
	x = x % d

	if x < 0 {
		return x + d
	}

	return x
}
