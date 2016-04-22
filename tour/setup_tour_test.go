package tour

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

const (
	_String      string  = "impossibly lame value"
	_Int         int     = -1
	_PositiveInt int     = 42
	_Byte        byte    = 255
	_Bool        bool    = false // ugh
	_Boolean     bool    = true  // oh well
	_Float32     float32 = -1.0
	_DeleteMe    bool    = false
)

type runner interface {
	run()
}

var _Runner runner

func TestTourKoans(t *testing.T) {
	//learnHello()
	learnPackages()

	fmt.Printf("\n%c[32;1mYou finished the tour. Great job!%c[0m\n\n", 27, 27)
}

func assert(o bool) {
	if !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, _GetRecentLine(), 27)
		os.Exit(1)
	}
}

func _GetRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
