package gotest

import (
	"time"

	"github.com/jstemmer/go-junit-report/v2/gtr"
	"github.com/jstemmer/go-junit-report/v2/parser/gotest/internal/collector"
)

const (
	globalID = 0
)

// reportBuilder helps build a test Report from a collection of events.
//
// The reportBuilder delegates to the packageBuilder for creating packages from
// basic test events, but keeps track of build errors itself. The reportBuilder
// is also responsible for generating unique test id's.
//
// Test output is collected by the output collector, which also keeps track of
// the currently active test so output is automatically associated with the
// correct test.
type reportBuilder struct {
	packageBuilders map[string]*packageBuilder
	buildErrors     map[int]gtr.Error

	nextID   int               // next free unused id
	output   *collector.Output // output collected for each id
	packages []gtr.Package     // completed packages

	// options
	packageName   string
	subtestMode   SubtestMode
	timestampFunc func() time.Time
}

// newReportBuilder creates a new reportBuilder.
func newReportBuilder() *reportBuilder { _ = "STUB: not implemented"; return nil }

// getPackageBuilder returns the packageBuilder for the given packageName. If
// no packageBuilder exists for the given package, a new one is created.
func (b *reportBuilder) getPackageBuilder(packageName string) *packageBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ProcessEvent takes a test event and adds it to the report.
func (b *reportBuilder) ProcessEvent(ev Event) { _ = "STUB: not implemented"; return }

// The summary marks the end of a package. We can now create the actual
// package from all the events we've processed so far for this package.

// This shouldn't happen, but just in case print a warning and ignore
// this event.

// newID returns a new unique id.
func (b *reportBuilder) generateID() int { _ = "STUB: not implemented"; return 0 }

// Build returns the new Report containing all the tests, build errors and
// their output created from the processed events.
func (b *reportBuilder) Build() gtr.Report {
	_ = "STUB: not implemented"
	// Create packages for any leftover package builders.
	return *new(gtr.Report)
}

// Create packages for any leftover build errors.

// CreateBuildError creates a new build error and marks it as active.
func (b *reportBuilder) CreateBuildError(packageName string) { _ = "STUB: not implemented"; return }

// CreatePackage returns a new package containing all the build errors, output,
// tests and benchmarks created so far. The optional packageName is used to
// find the correct reportBuilder. The newPackageName is the actual package
// name that will be given to the returned package, which should be used in
// case the packageName was unknown until this point.
func (b *reportBuilder) CreatePackage(packageName, newPackageName, result string, duration time.Duration, data string) gtr.Package {
	_ = "STUB: not implemented"
	return *new(gtr.Package)
}

// First check if this package contained a build error. If that's the case,
// we won't find any tests in this package.

// Get the packageBuilder for this package and make sure it's deleted, so
// future events for this package will use a new packageBuilder.

// If the packageBuilder is empty, we never received any events for this
// package so there's no need to continue.

// However, we should at least report an error if the result says we
// failed.

// If we've collected output, but there were no tests, then this package
// had a runtime error or it simply didn't have any tests.

// If the summary result says we failed, but there were no failing tests
// then something else must have failed.

// Collect tests for this package

// Sort packages by id to ensure we maintain insertion order.

// parseResult returns a gtr.Result for the given result string r.
func parseResult(r string) gtr.Result { _ = "STUB: not implemented"; return *new(gtr.Result) }

// groupBenchmarksByName groups tests with the Benchmark prefix if they have
// the same name and combines their output.
func groupBenchmarksByName(tests []gtr.Test, output *collector.Output) []gtr.Test {
	_ = "STUB: not implemented"
	return nil
}

// If this test is not a benchmark, we won't group it by name but
// just add it to the final result.

// combinedDuration returns the sum of the durations of the given tests.
func combinedDuration(tests []gtr.Test) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// groupResults returns the result we should use for a collection of tests.
func groupResults(tests []gtr.Test) gtr.Result { _ = "STUB: not implemented"; return *new(gtr.Result) }

// packageBuilder helps build a gtr.Package from a collection of test events.
type packageBuilder struct {
	generateID func() int
	output     *collector.Output

	tests     map[int]gtr.Test
	parentIDs map[int]struct{} // set of test id's that contain subtests
	coverage  float64          // coverage percentage
}

// newPackageBuilder creates a new packageBuilder. New tests will be assigned
// an ID returned by the generateID function. The activeIDSetter is called to
// set or reset the active test id.
func newPackageBuilder(generateID func() int, output *collector.Output) *packageBuilder {
	_ = "STUB: not implemented"
	return nil
}

// IsEmpty returns true if this package builder does not have any tests and has
// not collected any global output.
func (b packageBuilder) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// CreateTest adds a test with the given name to the package, marks it as
// active and returns its generated id.
func (b *packageBuilder) CreateTest(name string) int { _ = "STUB: not implemented"; return 0 }

// PauseTest marks the test with the given name no longer active. Any results
// or output added to the package after calling PauseTest will no longer be
// associated with this test.
func (b *packageBuilder) PauseTest(name string) { _ = "STUB: not implemented"; return }

// ContinueTest finds the test with the given name and marks it as active. If
// more than one test exist with this name, the most recently created test will
// be used.
func (b *packageBuilder) ContinueTest(name string) { _ = "STUB: not implemented"; return }

// EndTest finds the test with the given name, sets the result, duration and
// level. If more than one test exists with this name, the most recently
// created test will be used. If no test exists with this name, a new test is
// created. The test is then marked as no longer active.
func (b *packageBuilder) EndTest(name, result string, duration time.Duration, level int) {
	_ = "STUB: not implemented"
	return
}

// test did not exist, create one
// TODO: Likely reason is that the user ran go test without the -v
// flag, should we report this somewhere?

// End resets the active test.
func (b *packageBuilder) End() { _ = "STUB: not implemented"; return }

// BenchmarkResult updates an existing or adds a new test with the given
// results and marks it as active. If an existing test with this name exists
// but without result, then that one is updated. Otherwise a new one is added
// to the report.
func (b *packageBuilder) BenchmarkResult(name string, iterations int64, nsPerOp, mbPerSec float64, bytesPerOp, allocsPerOp int64) {
	_ = "STUB: not implemented"
	return
}

// Coverage sets the code coverage percentage.
func (b *packageBuilder) Coverage(pct float64, packages []string) {
	_ = "STUB: not implemented"

	// Output appends data to the output of this package.
	return
}

func (b *packageBuilder) Output(data string) { _ = "STUB: not implemented"; return }

// findTest returns the id of the most recently created test with the given
// name if it exists.
func (b *packageBuilder) findTest(name string) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// findTestParentID searches the existing tests in this package for a parent of
// the test with the given name, and returns its id if one is found.
func (b *packageBuilder) findTestParentID(name string) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// isParent returns true if the test with the given id has sub tests.
func (b *packageBuilder) isParent(id int) bool { _ = "STUB: not implemented"; return false }

// dropLastSegment strips the last `/` and everything following it from the
// given name. If no `/` was found, the empty string is returned.
func dropLastSegment(name string) string { _ = "STUB: not implemented"; return "" }

// containsFailures return true if this package contains at least one failing
// test or a test with an unknown result.
func (b *packageBuilder) containsFailures() bool { _ = "STUB: not implemented"; return false }
