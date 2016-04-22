package tour

// Make the `fmt` and `math/rand` packages available
import (
	"fmt"
	"math"
	"math/rand"
)

func learnPackages() {
	// The fmt package has both printing and string formatting functions
	assert(fmt.Sprintf("%s, %s!", "Hello", "Gopher") == _String)

	// Exported names are capitalized
	assert(math.Sqrt(9) == _Float64)                // Too easy
	assert(fmt.Sprintf("%.2f", math.Pi) == _String) // Keep it short

	// Run `go doc math/rand.Intn` at the command line to see what it does
	var randInt = rand.Intn(10)
	assert(_Int > randInt && randInt > _Int)
}
