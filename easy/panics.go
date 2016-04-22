package easy

func divideFourBy(i int) int {
	return 4 / i
}

const _Divisor = 0

func aboutPanics() {
	assert(_DeleteMe) // panics are exceptional errors at runtime

	n := divideFourBy(_Divisor)
	assert(n == 2) // panics are exceptional errors at runtime
}
