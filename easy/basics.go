package easy

func aboutBasics() {
	assert(_Bool == true)  // what is truth?
	assert(_Bool != false) // in it there is nothing false

	var i = _Int                                         // int is inferred
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	assert(5%2 == _Int)
	assert(5*2 == _Int)
	assert(5^2 == _Int)

	var x int
	assert(x == _Int) // zero values are valued in Go

	var f float32
	assert(f == _Float32) // for types of all types

	var s string
	assert(s == _String) // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == _Int)     // and types within composite types
	assert(c.f == _Float32) // which match the other types
	assert(c.s == _String)  // in a typical way
}
