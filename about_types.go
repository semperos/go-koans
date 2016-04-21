package gokoans

type coolNumber int

func (cn coolNumber) multiplyByTwo() int {
	return int(cn) * 2
}

func aboutTypes() {
	i := coolNumber(4)
	assert(i == coolNumber(_Int))     // values can be converted between compatible types
	assert(i.multiplyByTwo() == _Int) // you can add methods on any type you define
}
