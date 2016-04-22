package tour

// Make the `fmt` and `math/rand` packages available
import (
	"fmt"
	"math/rand"
)

func learnPackages() {
	// The fmt package has both printing and string formatting functions
	assert(fmt.Sprintf("%s, %s!", "Hello", "Gopher") == _String)

	// TODO Add simpler math.sqrt(7) example here.

	// Run `go doc math/rand.Intn` at the command line to see what it does
	var randInt = rand.Intn(10)
	assert(_Int > randInt && randInt > _Int)
}
