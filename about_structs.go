package gokoans

func aboutStructs() {
	var bob struct {
		name string
		age  int
	}
	bob.name = "bob"
	bob.age = 30

	assert(bob.name == _String) // structs are collections of named variables
	assert(bob.age == _Int)     // each field has both setter and getter behavior

	type person struct {
		name string
		age  int
	}

	var john person
	john.name = "bob"
	john.age = _Int

	assert(bob == john) // assuredly, bob is certainly not john.. yet
}
