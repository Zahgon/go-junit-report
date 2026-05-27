// Package collector collects output lines grouped by id and provides ways to
// retrieve and merge output ordered by the time each line was added.
package collector

import (
	"time"
)

// line is a single line of output captured at some point in time.
type line struct {
	Timestamp time.Time
	Text      string
}

// Output stores output lines grouped by id. Output can be retrieved for one or
// more ids and output for different ids can be merged together, while
// preserving their insertion original order based on the time it was
// collected.
// Output also tracks the active id, so you can append output without providing
// an id.
type Output struct {
	m  map[int][]line
	id int // active id
}

// New returns a new output collector.
func New() *Output { _ = "STUB: not implemented"; return nil }

// Clear deletes all output for the given id.
func (o *Output) Clear(id int) {
	_ = "STUB: not implemented"

	// Append appends the given line of text to the output of the currently active
	// id.
	return
}

func (o *Output) Append(text string) { _ = "STUB: not implemented"; return }

// AppendToID appends the given line of text to the output of the given id.
func (o *Output) AppendToID(id int, text string) { _ = "STUB: not implemented"; return }

// Contains returns true if any output lines were collected for the given id.
func (o *Output) Contains(id int) bool { _ = "STUB: not implemented"; return false }

// Get returns the output lines for the given id.
func (o *Output) Get(id int) []string { _ = "STUB: not implemented"; return nil }

// GetAll returns the output lines for all ids sorted by the collection
// timestamp of each line of output.
func (o *Output) GetAll(ids ...int) []string { _ = "STUB: not implemented"; return nil }

// Merge merges the output lines from fromID into intoID, and sorts the output
// by the collection timestamp of each line of output.
func (o *Output) Merge(fromID, intoID int) { _ = "STUB: not implemented"; return }

// SetActiveID sets the active id. Text appended to this output will be
// associated with the active id.
func (o *Output) SetActiveID(id int) { _ = "STUB: not implemented"; return }
