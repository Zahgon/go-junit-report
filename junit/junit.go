// Package junit defines a JUnit XML report and includes convenience methods
// for working with these reports.
package junit

import (
	"encoding/xml"
	"io"
	"time"

	"github.com/jstemmer/go-junit-report/v2/gtr"
)

// Testsuites is a collection of JUnit testsuites.
type Testsuites struct {
	XMLName xml.Name `xml:"testsuites"`

	Name     string `xml:"name,attr,omitempty"`
	Time     string `xml:"time,attr,omitempty"` // total duration in seconds
	Tests    int    `xml:"tests,attr,omitempty"`
	Errors   int    `xml:"errors,attr,omitempty"`
	Failures int    `xml:"failures,attr,omitempty"`
	Skipped  int    `xml:"skipped,attr,omitempty"`
	Disabled int    `xml:"disabled,attr,omitempty"`

	Suites []Testsuite `xml:"testsuite,omitempty"`
}

// AddSuite adds a Testsuite and updates this testssuites' totals.
func (t *Testsuites) AddSuite(ts Testsuite) { _ = "STUB: not implemented"; return }

// WriteXML writes the XML representation of Testsuites t to writer w.
func (t *Testsuites) WriteXML(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Testsuite is a single JUnit testsuite containing testcases.
type Testsuite struct {
	// required attributes
	Name     string `xml:"name,attr"`
	Tests    int    `xml:"tests,attr"`
	Failures int    `xml:"failures,attr"`
	Errors   int    `xml:"errors,attr"`
	ID       int    `xml:"id,attr"`

	// optional attributes
	Disabled  int    `xml:"disabled,attr,omitempty"`
	Hostname  string `xml:"hostname,attr,omitempty"`
	Package   string `xml:"package,attr,omitempty"`
	Skipped   int    `xml:"skipped,attr,omitempty"`
	Time      string `xml:"time,attr"`                // duration in seconds
	Timestamp string `xml:"timestamp,attr,omitempty"` // date and time in ISO8601
	File      string `xml:"file,attr,omitempty"`

	Properties *[]Property `xml:"properties>property,omitempty"`
	Testcases  []Testcase  `xml:"testcase,omitempty"`
	SystemOut  *Output     `xml:"system-out,omitempty"`
	SystemErr  *Output     `xml:"system-err,omitempty"`
}

// AddProperty adds a property with the given name and value to this Testsuite.
func (t *Testsuite) AddProperty(name, value string) { _ = "STUB: not implemented"; return }

// AddTestcase adds Testcase tc to this Testsuite.
func (t *Testsuite) AddTestcase(tc Testcase) { _ = "STUB: not implemented"; return }

// SetTimestamp sets the timestamp in this Testsuite.
func (t *Testsuite) SetTimestamp(timestamp time.Time) { _ = "STUB: not implemented"; return }

// Testcase represents a single test with its results.
type Testcase struct {
	// required attributes
	Name      string `xml:"name,attr"`
	Classname string `xml:"classname,attr"`

	// optional attributes
	Time   string `xml:"time,attr,omitempty"` // duration in seconds
	Status string `xml:"status,attr,omitempty"`

	Skipped   *Result `xml:"skipped,omitempty"`
	Error     *Result `xml:"error,omitempty"`
	Failure   *Result `xml:"failure,omitempty"`
	SystemOut *Output `xml:"system-out,omitempty"`
	SystemErr *Output `xml:"system-err,omitempty"`
}

// Property represents a key/value pair.
type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// Result represents the result of a single test.
type Result struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr,omitempty"`
	Data    string `xml:",cdata"`
}

// Output represents output written to stdout or sderr.
type Output struct {
	Data string `xml:",cdata"`
}

// CreateFromReport creates a JUnit representation of the given gtr.Report.
func CreateFromReport(report gtr.Report, hostname string) Testsuites {
	_ = "STUB: not implemented"
	return *new(Testsuites)
}

// JUnit doesn't have a good way of dealing with build or runtime
// errors that happen before a test has started, so we create a single
// failing test that contains the build error details.

func createTestcaseForTest(pkgName string, test gtr.Test) Testcase {
	_ = "STUB: not implemented"
	return *new(Testcase)
}

// formatDuration returns the JUnit string representation of the given
// duration.
func formatDuration(d time.Duration) string { _ = "STUB: not implemented"; return "" }

// formatOutput combines the lines from the given output into a single string.
func formatOutput(output []string) string { _ = "STUB: not implemented"; return "" }

func escapeIllegalChars(str string) string { _ = "STUB: not implemented"; return "" }

// Decide whether the given rune is in the XML Character Range, per
// the Char production of https://www.xml.com/axml/testaxml.htm,
// Section 2.2 Characters.
// From: encoding/xml/xml.go
func isInCharacterRange(r rune) (inrange bool) { _ = "STUB: not implemented"; return false }
