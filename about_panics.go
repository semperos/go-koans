package gokoans

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	assert(_DeleteMe) // panics are exceptional errors at runtime

	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
