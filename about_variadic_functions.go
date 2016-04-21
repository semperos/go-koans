package gokoans

import "strings"

func concatNames(sep string, names ...string) string {
	return strings.Join(names, sep) // variadic parameters are really just slices
}

func aboutVariadicFunctions() {
	{
		str := concatNames(" ", "bob", "billy", "fred")
		assert(str == _String) // several values can be passed to variadic parameters
	}

	{
		names := []string{"bob", "billy", "fred"}
		str := concatNames("-", names...)
		assert(str == _String) // or a slice can be dotted in place of all of them
	}
}
