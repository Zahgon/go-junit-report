// Package gtr defines a standard test report format and provides convenience
// methods to create and convert reports.
package gtr

import (
	"time"
)

// Result is the result of a test.
type Result int

// Test results.
const (
	Unknown Result = iota
	Pass
	Fail
	Skip
)

func (r Result) String() string { _ = "STUB: not implemented"; return "" }

// Report contains the build and test results of a collection of packages.
type Report struct {
	Packages []Package
}

// IsSuccessful returns true if none of the packages in this report have build
// or runtime errors and all tests passed without failures or were skipped.
func (r *Report) IsSuccessful() bool { _ = "STUB: not implemented"; return false }

// Package contains build and test results for a single package.
type Package struct {
	Name       string
	Timestamp  time.Time
	Duration   time.Duration
	Coverage   float64
	Output     []string
	Properties []Property

	Tests []Test

	BuildError Error
	RunError   Error
}

// SetProperty stores a key/value property in the current package. If a
// property with the given key already exists, its old value will be
// overwritten with the given value.
func (p *Package) SetProperty(key, value string) {
	_ = "STUB: not implemented"
	// TODO(jstemmer): Delete this method in the next major release.
	// Delete all the properties whose name is the specified key,
	// then add the specified key-value property.
	return
}

// AddProperty appends a name/value property in the current package.
func (p *Package) AddProperty(name, value string) { _ = "STUB: not implemented"; return }

// Property is a name/value property.
type Property struct {
	Name, Value string
}

// Test contains the results of a single test.
type Test struct {
	ID       int
	Name     string
	Duration time.Duration
	Result   Result
	Level    int
	Output   []string
	Data     map[string]interface{}
}

// NewTest creates a new Test with the given id and name.
func NewTest(id int, name string) Test { _ = "STUB: not implemented"; return *new(Test) }

// Error contains details of a build or runtime error.
type Error struct {
	ID       int
	Name     string
	Duration time.Duration
	Cause    string
	Output   []string
}

// TrimPrefixSpaces trims the leading whitespace of the given line using the
// indentation level of the test. Printing logs in a Go test is typically
// prepended by blocks of 4 spaces to align it with the rest of the test
// output. TrimPrefixSpaces intends to only trim the whitespace added by the Go
// test command, without inadvertently trimming whitespace added by the test
// author.
func TrimPrefixSpaces(line string, indent int) string {
	_ = "STUB: not implemented"
	// We only want to trim the whitespace prefix if it was part of the test
	// output. Test output is usually prefixed by a series of 4-space indents,
	// so we'll check for that to decide whether this output was likely to be
	// from a test.
	return ""
}

// Use the subtest level to trim a consistently sized prefix from the
// output lines.
