// package assert provides simple assertion functions for a test.
package assert

import (
	"fmt"
	"io/ioutil"
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// Skips is the number of lines from stack-trace to remove to print the
// point of failure. It generally should be set to 1.
// Change the value if the line reported on failure is not the actual failing line.
// Set a different value inside an init() block.
//
//     import "github.com/chiku/assert"
//
//     func init() {
//         assert.Skips = 2
//     }
var Skips = 1

// RequireNoError verifies that err is nil.
// It prints the given message if err is not nil.
// The test is aborted on failure.
func RequireNoError(t *testing.T, err error, msg string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(Skips)
		fileBase := path.Base(file)

		fmt.Printf("\t%v:%v: %s\n", fileBase, line, msg)
		fmt.Printf("\t%v:%v: %s\n\n", fileBase, line, err.Error())
		t.FailNow()
	}
}

// RequireError verifies that err not nil.
// It prints the given message if err is nil.
// The test is aborted on failure.
func RequireError(t *testing.T, err error, msg string) {
	if err == nil {
		_, file, line, _ := runtime.Caller(Skips)
		fileBase := path.Base(file)

		fmt.Printf("\t%v:%v: %s\n", fileBase, line, msg)
		t.FailNow()
	}
}

// AssertEqual verifies the actual equal expected.
// It prints the given message if the two aren't equal.
// The equality is checked using reflect.DeepEquals. The test continues on failure.
func AssertEqual(t *testing.T, actual, expected, msg interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		_, file, line, _ := runtime.Caller(Skips)
		fileBase := path.Base(file)

		fmt.Printf("\t%v:%v: %s\n", fileBase, line, msg)
		fmt.Printf("\t%v:%v: %#v != %#v\n\n", fileBase, line, actual, expected)
		t.Fail()
	}
}

// AssertContains verifies part is a sub-string of total.
// It prints the given message if it isn't.
// The test continues on failure.
func AssertContains(t *testing.T, total, part, msg string) {
	if !strings.Contains(total, part) {
		_, file, line, _ := runtime.Caller(Skips)
		fileBase := path.Base(file)

		fmt.Printf("\t%v:%v: %s\n", fileBase, line, msg)
		fmt.Printf("\t%v:%v: %#v doesn't contain %#v\n\n", fileBase, line, total, part)
		t.Fail()
	}
}

// CreateFile creates a temporary file with the given contents.
// It returns the file-name of the created file.
// The caller is expected to delete the created file.
// The test aborts on failure.
func CreateFile(t *testing.T, content string) string {
	tmpfile, err := ioutil.TempFile("", "example")
	RequireNoError(t, err, "Expected no error creating temporary file")
	_, err = tmpfile.Write([]byte(content))
	RequireNoError(t, err, "Expected no error writing to temporary file")
	err = tmpfile.Close()
	RequireNoError(t, err, "Expected no error closing temporary file")

	return tmpfile.Name()
}
