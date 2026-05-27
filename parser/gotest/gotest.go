// Package gotest is a standard Go test output parser.
package gotest

import (
	"io"
	"regexp"
	"time"

	"github.com/jstemmer/go-junit-report/v2/gtr"
	"github.com/jstemmer/go-junit-report/v2/parser/gotest/internal/reader"
)

const (
	// maxLineSize is the maximum amount of bytes we'll read for a single line.
	// Lines longer than maxLineSize will be truncated.
	maxLineSize = 4 * 1024 * 1024
)

var (
	// regexBenchInfo captures 3-5 groups: benchmark name, number of times ran, ns/op (with or without decimal), MB/sec (optional), B/op (optional), and allocs/op (optional).
	regexBenchmark    = regexp.MustCompile(`^(Benchmark[^ -]+)$`)
	regexBenchSummary = regexp.MustCompile(`^(Benchmark[^ -]+)(?:-\d+\s+|\s+)(\d+)\s+(\d+|\d+\.\d+)\sns\/op(?:\s+(\d+|\d+\.\d+)\sMB\/s)?(?:\s+(\d+)\sB\/op)?(?:\s+(\d+)\sallocs/op)?`)
	regexCoverage     = regexp.MustCompile(`^coverage:\s+(\d+|\d+\.\d+)%\s+of\s+statements(?:\sin\s(.+))?$`)
	regexEndBenchmark = regexp.MustCompile(`^--- (BENCH|FAIL|SKIP): (Benchmark[^ -]+)(?:-\d+)?$`)
	regexEndTest      = regexp.MustCompile(`((?:    )*)--- (PASS|FAIL|SKIP): ([^ ]+) \((\d+\.\d+)(?: seconds|s)\)`)
	regexStatus       = regexp.MustCompile(`^(PASS|FAIL|SKIP)$`)
	regexSummary      = regexp.MustCompile(`` +
		// 1: result
		`^(\?|ok|FAIL)` +
		// 2: package name
		`\s+([^ \t]+)` +
		// 3: duration (optional)
		`(?:\s+(\d+\.\d+)s)?` +
		// 4: cached indicator (optional)
		`(?:\s+(\(cached\)))?` +
		// 5: coverage percentage (optional)
		// 6: coverage package list (optional)
		`(?:\s+coverage:\s+(?:\[no\sstatements\]|(\d+\.\d+)%\sof\sstatements(?:\sin\s(.+))?))?` +
		// 7: [status message] (optional)
		`(?:\s+(\[[^\]]+\]))?` +
		`$`)
)

// Option defines options that can be passed to gotest.New.
type Option func(*Parser)

// PackageName is an Option that sets the default package name to use when it
// cannot be determined from the test output.
func PackageName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// TimestampFunc is an Option that sets the timestamp function that is used to
// determine the current time when creating the Report. This can be used to
// override the default behaviour of using time.Now().
func TimestampFunc(f func() time.Time) Option { _ = "STUB: not implemented"; return *new(Option) }

// SubtestMode configures how Go subtests should be handled by the parser.
type SubtestMode string

const (
	// SubtestModeDefault is the default subtest mode. It treats tests with
	// subtests as any other tests.
	SubtestModeDefault SubtestMode = ""

	// IgnoreParentResults ignores test results for tests with subtests. Use
	// this mode if you use subtest parents for common setup/teardown, but are
	// not interested in counting them as failed tests. Ignoring their results
	// still preserves these tests and their captured output in the report.
	IgnoreParentResults SubtestMode = "ignore-parent-results"

	// ExcludeParents excludes tests that contain subtests from the report.
	// Note that the subtests themselves are not removed. Use this mode if you
	// use subtest parents for common setup/teardown, but are not actually
	// interested in their presence in the created report. If output was
	// captured for tests that are removed, the output is preserved in the
	// global report output.
	ExcludeParents SubtestMode = "exclude-parents"
)

// ParseSubtestMode returns a SubtestMode for the given string.
func ParseSubtestMode(in string) (SubtestMode, error) {
	_ = "STUB: not implemented"
	return *new(SubtestMode), nil
}

// SetSubtestMode is an Option to change how the parser handles tests with
// subtests. See the documentation for the individual SubtestModes for more
// information.
func SetSubtestMode(mode SubtestMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// Parser is a Go test output Parser.
type Parser struct {
	packageName string
	subtestMode SubtestMode

	timestampFunc func() time.Time

	events []Event
}

// NewParser returns a new Go test output parser.
func NewParser(options ...Option) *Parser { _ = "STUB: not implemented"; return nil }

// Parse parses Go test output from the given io.Reader r and returns
// gtr.Report.
func (p *Parser) Parse(r io.Reader) (gtr.Report, error) {
	_ = "STUB: not implemented"
	return *new(gtr.Report), nil
}

func (p *Parser) parse(r reader.LineReader) (gtr.Report, error) {
	_ = "STUB: not implemented"
	return *new(gtr.Report), nil
}

// Lines that exceed bufio.MaxScanTokenSize are not expected to contain
// any relevant test infrastructure output, so instead of parsing them
// we treat them as regular output to increase performance.
//
// Parser used a bufio.Scanner in the past, which only supported
// reading lines up to bufio.MaxScanTokenSize in length. Since this
// turned out to be fine in almost all cases, it seemed an appropriate
// value to use to decide whether or not to attempt parsing this line.

// Events returns the events created by the parser.
func (p *Parser) Events() []Event { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseLine(line string) (events []Event) { _ = "STUB: not implemented"; return nil }

// for compatibility with gotest 1.20+ https://go-review.git.corp.google.com/c/go/+/443596

func (p *Parser) runTest(name string) []Event { _ = "STUB: not implemented"; return nil }

func (p *Parser) pauseTest(name string) []Event { _ = "STUB: not implemented"; return nil }

func (p *Parser) contTest(name string) []Event { _ = "STUB: not implemented"; return nil }

func (p *Parser) endTest(line, indent, result, name, duration string) []Event {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) status(result string) []Event { _ = "STUB: not implemented"; return nil }

func (p *Parser) summary(result, name, duration, cached, status, covpct, packages string) []Event {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) coverage(percent, packages string) []Event { _ = "STUB: not implemented"; return nil }

func (p *Parser) runBench(name string) []Event { _ = "STUB: not implemented"; return nil }

func (p *Parser) benchSummary(name, iterations, nsPerOp, mbPerSec, bytesPerOp, allocsPerOp string) []Event {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) endBench(result, name string) []Event { _ = "STUB: not implemented"; return nil }

func (p *Parser) buildOutput(packageName string) []Event { _ = "STUB: not implemented"; return nil }

func (p *Parser) output(line string) []Event { _ = "STUB: not implemented"; return nil }

func parseSeconds(s string) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// ignore error

func parseFloat(s string) float64 { _ = "STUB: not implemented"; return 0 }

// ignore error

func parsePackages(pkgList string) []string { _ = "STUB: not implemented"; return nil }

func parseInt(s string) int64 {
	_ = "STUB: not implemented"
	// ignore error
	return 0
}

func stripIndent(line string) (string, int) { _ = "STUB: not implemented"; return "", 0 }
