package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

const (
	colorRed  = "\x1b[31m"
	colorNone = "\x1b[0m"
)

// Equals fails the test if expected is not equal to actual.
// Taken from https://github.com/benbjohnson/testing
func Equals(tb testing.TB, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf(colorRed+"\n%s:%d:\n\n\texp: %#v\n\n\tgot: %#v"+colorNone+"\n\n", filepath.Base(file), line, expected, actual)
		tb.FailNow()
	}
}

// NoError fails the test if an err is not nil.
// Taken from https://github.com/benbjohnson/testing
func NoError(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf(colorRed+"\n%s:%d: unexpected error: %s"+colorNone+"\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// True fails the test if the condition is false.
// Taken from https://github.com/benbjohnson/testing
func True(tb testing.TB, condition bool, message string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf(colorRed+"\n%s:%d: "+message+colorNone+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// False fails the test if the condition is true.
func False(tb testing.TB, condition bool, message string, v ...interface{}) {
	if condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf(colorRed+"\n%s:%d: "+message+colorNone+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}
