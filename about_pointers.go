package gokoans

func aboutPointers() {
	{
		a := 3
		b := a // 'b' is a copy of 'a' (all assignments are copy-operations)

		b++

		assert(a == _Int) // variables are independent of one another
	}

	{
		a := 3
		b := &a // 'b' is the address of 'a'

		*b = *b + 2          // de-referencing 'b' means acting like a mutable copy of 'a'
		assert(a == _Int) // pointers seem complicated at first but are actually simple
	}

	{
		increment := func(i int) {
			i++
		}

		a := 3
		increment(a)
		assert(a == _Int) // variables are always passed by value, and so a copy is made
	}

	{
		realIncrement := func(i *int) {
			(*i)++
		}

		b := 3
		realIncrement(&b)
		assert(b == _Int) // but passing a pointer allows others to mutate the value pointed to
	}
}
