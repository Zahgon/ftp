package ftp

import "io"

type debugWrapper struct {
	conn io.ReadWriteCloser
	io.Reader
	io.Writer
}

func newDebugWrapper(conn io.ReadWriteCloser, w io.Writer) io.ReadWriteCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadWriteCloser)
}

func (w *debugWrapper) Close() error { _ = "STUB: not implemented"; return nil }

type streamDebugWrapper struct {
	io.Reader
	closer io.ReadCloser
}

func newStreamDebugWrapper(rd io.ReadCloser, w io.Writer) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (w *streamDebugWrapper) Close() error { _ = "STUB: not implemented"; return nil }
