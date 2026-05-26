package ftp

import (
	"errors"
	"time"
)

var errUnsupportedListLine = errors.New("unsupported LIST line")
var errUnsupportedListDate = errors.New("unsupported LIST date")
var errUnknownListEntryType = errors.New("unknown entry type")

type parseFunc func(string, time.Time, *time.Location) (*Entry, error)

var listLineParsers = []parseFunc{
	parseRFC3659ListLine,
	parseLsListLine,
	parseDirListLine,
	parseHostedFTPLine,
}

var dirTimeFormats = []string{
	"01-02-06  03:04PM",
	"2006-01-02  15:04",
	"01-02-2006  03:04PM",
	"01-02-2006  15:04",
}

// parseRFC3659ListLine parses the style of directory line defined in RFC 3659.
func parseRFC3659ListLine(line string, _ time.Time, loc *time.Location) (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseNextRFC3659ListLine(line string, loc *time.Location, e *Entry) (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// All lines must have the same name

// parseLsListLine parses a directory line in a format based on the output of
// the UNIX ls command.
func parseLsListLine(line string, now time.Time, loc *time.Location) (*Entry, error) {
	_ = "STUB: not implemented"

	// Has the first field a length of exactly 10 bytes
	// - or 10 bytes with an additional '+' character for indicating ACLs?
	// If not, return.
	return nil, nil
}

// Read two more fields

// Split link name and target

// parseDirListLine parses a directory line in a format based on the output of
// the MS-DOS DIR command.
func parseDirListLine(line string, now time.Time, loc *time.Location) (*Entry, error) {
	_ = "STUB: not implemented"
	return nil,

		// Try various time formats that DIR might use, and stop when one works.
		nil
}

// None of the time formats worked.

// parseHostedFTPLine parses a directory line in the non-standard format used
// by hostedftp.com
// -r--------   0 user group     65222236 Feb 24 00:39 UABlacklistingWeek8.csv
// (The link count is inexplicably 0)
func parseHostedFTPLine(line string, now time.Time, loc *time.Location) (*Entry, error) {
	_ = "STUB: not implemented"
	// Has the first field a length of 10 bytes?
	return nil, nil
}

// Set link count to 1 and attempt to parse as Unix.

// parseListLine parses the various non-standard format returned by the LIST
// FTP command.
func parseListLine(line string, now time.Time, loc *time.Location) (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Entry) setSize(str string) (err error) { _ = "STUB: not implemented"; return nil }

func (e *Entry) setTime(fields []string, now time.Time, loc *time.Location) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// contains time

/*
	On unix, `info ls` shows:

	10.1.6 Formatting file timestamps
	---------------------------------

	A timestamp is considered to be “recent” if it is less than six
	months old, and is not dated in the future.  If a timestamp dated today
	is not listed in recent form, the timestamp is in the future, which
	means you probably have clock skew problems which may break programs
	like ‘make’ that rely on file timestamps.
*/

// only the date
