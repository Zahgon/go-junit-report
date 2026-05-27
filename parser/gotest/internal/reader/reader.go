package reader

import (
	"bufio"
	"io"
	"time"
)

// LineReader is an interface to read lines with optional Metadata.
type LineReader interface {
	ReadLine() (string, *Metadata, error)
}

// Metadata contains metadata that belongs to a line.
type Metadata struct {
	Package string
}

// LimitedLineReader reads lines from an io.Reader object with a configurable
// line size limit. Lines exceeding the limit will be truncated, but read
// completely from the underlying io.Reader.
type LimitedLineReader struct {
	r     *bufio.Reader
	limit int
}

var _ LineReader = &LimitedLineReader{}

// NewLimitedLineReader returns a LimitedLineReader to read lines from r with a
// maximum line size of limit.
func NewLimitedLineReader(r io.Reader, limit int) *LimitedLineReader {
	_ = "STUB: not implemented"
	return nil
}

// ReadLine returns the next line from the underlying reader. The length of the
// line will not exceed the configured limit. ReadLine either returns a line or
// it returns an error, never both.
func (r *LimitedLineReader) ReadLine() (string, *Metadata, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Line is incomplete, keep reading until we reach the end of the line.

// ignore err, always nil

// Stop writing to buf if we exceed the limit. We continue reading
// however to make sure we consume the entire line.

// ignore err, always nil

// Event represents a JSON event emitted by `go test -json`.
type Event struct {
	Time    time.Time
	Action  string
	Package string
	Test    string
	Elapsed float64 // seconds
	Output  string
}

// JSONEventReader reads JSON events from an io.Reader object.
type JSONEventReader struct {
	r *LimitedLineReader
}

var _ LineReader = &JSONEventReader{}

// jsonLineLimit is the maximum size of a single JSON line emitted by `go test
// -json`.
const jsonLineLimit = 64 * 1024

// NewJSONEventReader returns a JSONEventReader to read the data in JSON
// events from r.
func NewJSONEventReader(r io.Reader) *JSONEventReader { _ = "STUB: not implemented"; return nil }

// ReadLine returns the next line from the underlying reader.
func (r *JSONEventReader) ReadLine() (string, *Metadata, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Skip events without output
