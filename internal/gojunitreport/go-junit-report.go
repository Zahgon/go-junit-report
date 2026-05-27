package gojunitreport

import (
	"io"
	"time"

	"github.com/jstemmer/go-junit-report/v2/gtr"
	"github.com/jstemmer/go-junit-report/v2/parser/gotest"
)

type parser interface {
	Parse(r io.Reader) (gtr.Report, error)
	Events() []gotest.Event
}

// Config contains the go-junit-report command configuration.
type Config struct {
	Parser        string
	Hostname      string
	PackageName   string
	SkipXMLHeader bool
	SubtestMode   gotest.SubtestMode
	Properties    map[string]string
	TimestampFunc func() time.Time

	// For debugging
	PrintEvents bool
}

// Run runs the go-junit-report command and returns the generated report.
func (c Config) Run(input io.Reader, output io.Writer) (*gtr.Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) writeJunitXML(w io.Writer, report gtr.Report) error {
	_ = "STUB: not implemented"
	return nil
}

func (c Config) gotestOptions() []gotest.Option { _ = "STUB: not implemented"; return nil }
