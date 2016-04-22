package easy

func aboutAllocation() {
	a := new(int)
	*a = 3
	assert(*a == _Int) // new() creates a pointer to the given type, like malloc() in C

	type person struct {
		name string
		age  int
	}
	bob := new(person)
	assert(bob.age == _Int) // it can allocate memory for custom types as well

	slice := make([]int, 3)
	assert(len(slice) == _Int) // make() creates slices of a given length

	slice = make([]int, 3, _PositiveInt) // but can also take an optional capacity
	assert(cap(slice) == 20)

	m := make(map[int]string)
	assert(len(m) == _Int) // make() also creates maps
}
