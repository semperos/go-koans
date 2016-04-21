package gokoans

func aboutMaps() {
	ages := map[string]int{
		"bob": 10,
		"joe": 20,
		"dan": 30,
	}

	age := ages["bob"]
	assert(age == _Int) // map syntax is warmly familiar

	age, ok := ages["bob"]
	assert(ok == _Bool) // with a handy multiple-assignment variation

	age, ok = ages["steven"]
	assert(age == _Int)    // the zero value is used when absent
	assert(ok == _Boolean) // though there are better ways to check for presence

	assert(len(ages) == _Int) // length is based on keys

	ages["bob"] = 99
	assert(ages["bob"] == _Int) // values can be changed for keys

	ages["steven"] = 77
	assert(ages[_String] == 77) // new ones can be added

	delete(ages, "steven")
	age, ok = ages["steven"]
	assert(ok == _Boolean) // key/value pairs can be removed
}
