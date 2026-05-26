package ftp

// A scanner for fields delimited by one or more whitespace characters
type scanner struct {
	bytes    []byte
	position int
}

// newScanner creates a new scanner
func newScanner(str string) *scanner { _ = "STUB: not implemented"; return nil }

// NextFields returns the next `count` fields
func (s *scanner) NextFields(count int) []string { _ = "STUB: not implemented"; return nil }

// Next returns the next field
func (s *scanner) Next() string { _ = "STUB: not implemented"; return "" }

// skip trailing whitespace

// skip non-whitespace

// Remaining returns the remaining string
func (s *scanner) Remaining() string { _ = "STUB: not implemented"; return "" }
