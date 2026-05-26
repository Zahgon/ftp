package ftp

// Walker traverses the directory tree of a remote FTP server
type Walker struct {
	serverConn *ServerConn
	root       string
	cur        *item
	stack      []*item
	descend    bool
}

type item struct {
	path  string
	entry *Entry
	err   error
}

// Next advances the Walker to the next file or directory,
// which will then be available through the Path, Stat, and Err methods.
// It returns false when the walk stops at the end of the tree.
func (w *Walker) Next() bool {
	_ = "STUB: not implemented"
	// check if we need to init cur, maybe this should be inside Walk
	return false
}

// an error occurred, drop out and stop walking

// update cur

// reset SkipDir

// SkipDir tells the Next function to skip the currently processed directory
func (w *Walker) SkipDir() {
	_ = "STUB: not implemented"

	// Err returns the error, if any, for the most recent attempt by Next to
	// visit a file or a directory. If a directory has an error, the walker
	// will not descend in that directory
	return
}

func (w *Walker) Err() error {
	_ = "STUB: not implemented"

	// Stat returns info for the most recent file or directory
	// visited by a call to Next.
	return nil
}

func (w *Walker) Stat() *Entry {
	_ = "STUB: not implemented"

	// Path returns the path to the most recent file or directory
	// visited by a call to Next. It contains the argument to Walk
	// as a prefix; that is, if Walk is called with "dir", which is
	// a directory containing the file "a", Path will return "dir/a".
	return nil
}

func (w *Walker) Path() string { _ = "STUB: not implemented"; return "" }
