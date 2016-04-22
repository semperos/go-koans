package tour

// Simple function
func add(a int, b int) int {
	// Implement addition here.
	return 0
}

// Example of multiple results
func swap(alpha, beta string) (string, string) {
	// Implement swap here.
	return "multiple", "results"
}

func learnFunctions() {
	// Implement `add` above
	assert(add(21, 21) == _Int)

	// Now rewrite the signature to `add` so that you don't repeat the `int`
	// type. A signature of `a int, b int` can be shortened to `a, b int`. When
	// ready, remove the _DeleteMe assertion and run the tests again.
	assert(_DeleteMe)

	// Implement `swap`, such that it takes two string arguments,
	// and returns them in swapped order.
	alpha, beta := swap("work", "hard")
	assert(alpha == "hard" && beta == "work")

	// Now rewrite the signature to `swap` and make it used named return values.
	// Instead of plain `(string, string)`, you can use something like
	// `(left, right string)`, define those in the body of the function, and then
	// put `return` on a line by itself at the very end of the function.
	// See https://golang.org/doc/effective_go.html#named-results for more details.
	assert(_DeleteMe)
}
