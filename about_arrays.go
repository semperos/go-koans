package gokoans

func aboutArrays() {
	fruits := [4]string{"apple", "orange", "mango"}

	assert(fruits[0] == _String) // indexes begin at 0
	assert(fruits[1] == _String) // one is indeed the loneliest number
	assert(fruits[2] == _String) // it takes two to ...tango?
	assert(fruits[3] == _String) // there is no spoon, only an empty value

	assert(len(fruits) == _Int) // the length is what the length is
	assert(cap(fruits) == _Int) // it can hold no more

	assert(fruits == [4]string{}) // comparing arrays is not like comparing apples and oranges

	tasty_fruits := fruits[1:3]        // defining oneself as a variation of another
	assert(tasty_fruits[0] == _String) // slices of arrays share some data
	assert(tasty_fruits[1] == _String) // albeit slightly askewed

	assert(len(tasty_fruits) == _Int) // its length is manifest
	assert(cap(tasty_fruits) == _Int) // but its capacity is surprising!

	tasty_fruits[0] = "lemon" // are their shared roots truly identical?

	assert(fruits[0] == _String) // has this element remained the same?
	assert(fruits[1] == _String) // how about the second?
	assert(fruits[2] == _String) // surely one of these must have changed
	assert(fruits[3] == _String) // but who can know these things

	veggies := [...]string{"carrot", "pea"}

	assert(len(veggies) == _Int) // array literals need not repeat an obvious length
}
