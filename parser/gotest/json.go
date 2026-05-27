package gotest

import (
	"io"

	"github.com/jstemmer/go-junit-report/v2/gtr"
)

// NewJSONParser returns a new Go test json output parser.
func NewJSONParser(options ...Option) *JSONParser { _ = "STUB: not implemented"; return nil }

// JSONParser is a `go test -json` output Parser.
type JSONParser struct {
	gp *Parser
}

// Parse parses Go test json output from the given io.Reader r and returns
// gtr.Report.
func (p *JSONParser) Parse(r io.Reader) (gtr.Report, error) {
	_ = "STUB: not implemented"
	return *new(gtr.Report), nil
}

// Events returns the events created by the parser.
func (p *JSONParser) Events() []Event { _ = "STUB: not implemented"; return nil }
