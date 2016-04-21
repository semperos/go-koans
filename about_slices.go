package gokoans

func aboutSlices() {
	fruits := []string{"apple", "orange", "mango"}

	assert(fruits[0] == _String) // slices seem like arrays
	assert(len(fruits) == _Int)  // in nearly all respects

	tastyFruits := fruits[1:3]        // we can even slice slices
	assert(tastyFruits[0] == _String) // slices of slices also share the underlying data

	pregnancySlots := []string{"baby", "baby", "lemon"}
	assert(cap(pregnancySlots) == _Int) // the capacity is initially the length

	pregnancySlots = append(pregnancySlots, "baby!")
	assert(len(pregnancySlots) == _Int) // slices can be extended with append(), much like realloc in C
	assert(cap(pregnancySlots) == _Int) // but with better optimizations

	pregnancySlots = append(pregnancySlots, "another baby!?", "yet another, oh dear!", "they must be Catholic")

	assert(len(pregnancySlots) == _Int) // append() can take N arguments to append to the slice
	assert(cap(pregnancySlots) == _Int) // the capacity optimizations have a guessable algorithm
}
