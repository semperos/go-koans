package easy

import "fmt"

func aboutStrings() {
	assert("a"+_String == "abc") // string concatenation need not be difficult
	assert(len("abc") == _Int)   // and bounds are thoroughly checked

	assert("abc"[0] == _Byte) // their contents are reminiscent of C

	assert("smith"[2:] == _String)  // slicing may omit the end point
	assert("smith"[:4] == _String)  // or the beginning
	assert("smith"[2:4] == _String) // or neither
	assert("smith"[:] == _String)   // or both

	assert("smith" == _String) // they can be compared directly
	assert("smith" < _String)  // i suppose maybe this could be useful.. someday

	bytes := []byte{'a', 'b', 'c'}
	assert(string(bytes) == _String) // strings can be created from byte-slices

	bytes[0] = 'z'
	assert(string(bytes) == _String) // byte-slices can be mutated, although strings cannot

	assert(fmt.Sprintf("hello %s", _String) == "hello world") // our old friend sprintf returns
	assert(fmt.Sprintf("hello \"%s\"", "world") == _String)   // quoting is familiar
	assert(fmt.Sprintf("hello %q", "world") == _String)       // although it can be done more easily

	assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == _String) // "the root of all evil" is actually a misquotation, by the way
}
