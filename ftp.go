// Package ftp implements a FTP client as described in RFC 959.
//
// A textproto.Error is returned for errors at the protocol level.
package ftp

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/textproto"
	"time"
)

const (
	// 30 seconds was chosen as it's the
	// same duration as http.DefaultTransport's timeout.
	DefaultDialTimeout = 30 * time.Second
)

// EntryType describes the different types of an Entry.
type EntryType int

// The differents types of an Entry
const (
	EntryTypeFile EntryType = iota
	EntryTypeFolder
	EntryTypeLink
)

// TransferType denotes the formats for transferring Entries.
type TransferType string

// The different transfer types
const (
	TransferTypeBinary = TransferType("I")
	TransferTypeASCII  = TransferType("A")
)

// Time format used by the MDTM and MFMT commands
const timeFormat = "20060102150405"

// ServerConn represents the connection to a remote FTP server.
// A single connection only supports one in-flight data connection.
// It is not safe to be called concurrently.
type ServerConn struct {
	options *dialOptions
	conn    *textproto.Conn // connection wrapper for text protocol
	netConn net.Conn        // underlying network connection
	host    string

	// Server capabilities discovered at runtime
	features      map[string]string
	skipEPSV      bool
	mlstSupported bool
	mfmtSupported bool
	mdtmSupported bool
	mdtmCanWrite  bool
	usePRET       bool
}

// DialOption represents an option to start a new connection with Dial
type DialOption struct {
	setup func(do *dialOptions)
}

// dialOptions contains all the options set by DialOption.setup
type dialOptions struct {
	context         context.Context
	dialer          net.Dialer
	tlsConfig       *tls.Config
	explicitTLS     bool
	disableEPSV     bool
	disableUTF8     bool
	disableMLSD     bool
	writingMDTM     bool
	forceListHidden bool
	location        *time.Location
	debugOutput     io.Writer
	dialFunc        func(network, address string) (net.Conn, error)
	shutTimeout     time.Duration // time to wait for data connection closing status
}

// Entry describes a file and is returned by List().
type Entry struct {
	Name   string
	Target string // target of symbolic link
	Type   EntryType
	Size   uint64
	Time   time.Time
}

// Response represents a data-connection
type Response struct {
	conn   net.Conn
	c      *ServerConn
	closed bool
}

// Dial connects to the specified address with optional options
func Dial(addr string, options ...DialOption) (*ServerConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use the resolved IP address in case addr contains a domain name
// If we use the domain name, we might not resolve to the same IP.

// DialWithTimeout returns a DialOption that configures the ServerConn with specified timeout
func DialWithTimeout(timeout time.Duration) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithShutTimeout returns a DialOption that configures the ServerConn with
// maximum time to wait for the data closing status on control connection
// and nudging the control connection deadline before reading status.
func DialWithShutTimeout(shutTimeout time.Duration) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithDialer returns a DialOption that configures the ServerConn with specified net.Dialer
func DialWithDialer(dialer net.Dialer) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithNetConn returns a DialOption that configures the ServerConn with the underlying net.Conn
//
// Deprecated: Use [DialWithDialFunc] instead
func DialWithNetConn(conn net.Conn) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

// DialWithDisabledEPSV returns a DialOption that configures the ServerConn with EPSV disabled
// Note that EPSV is only used when advertised in the server features.
func DialWithDisabledEPSV(disabled bool) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithDisabledUTF8 returns a DialOption that configures the ServerConn with UTF8 option disabled
func DialWithDisabledUTF8(disabled bool) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithDisabledMLSD returns a DialOption that configures the ServerConn with MLSD option disabled
//
// This is useful for servers which advertise MLSD (eg some versions
// of Serv-U) but don't support it properly.
func DialWithDisabledMLSD(disabled bool) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithWritingMDTM returns a DialOption making ServerConn use MDTM to set file time
//
// This option addresses a quirk in the VsFtpd server which doesn't support
// the MFMT command for setting file time like other servers but by default
// uses the MDTM command with non-standard arguments for that.
// See "mdtm_write" in https://security.appspot.com/vsftpd/vsftpd_conf.html
func DialWithWritingMDTM(enabled bool) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithForceListHidden returns a DialOption making ServerConn use LIST -a to include hidden files and folders in directory listings
//
// This is useful for servers that do not do this by default, but it forces the use of the LIST command
// even if the server supports MLST.
func DialWithForceListHidden(enabled bool) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithLocation returns a DialOption that configures the ServerConn with specified time.Location
// The location is used to parse the dates sent by the server which are in server's timezone
func DialWithLocation(location *time.Location) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithContext returns a DialOption that configures the ServerConn with specified context
// The context will be used for the initial connection setup
func DialWithContext(ctx context.Context) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithTLS returns a DialOption that configures the ServerConn with specified TLS config
//
// If called together with the DialWithDialFunc option, the DialWithDialFunc function
// will be used when dialing new connections but regardless of the function,
// the connection will be treated as a TLS connection.
func DialWithTLS(tlsConfig *tls.Config) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithExplicitTLS returns a DialOption that configures the ServerConn to be upgraded to TLS
// See DialWithTLS for general TLS documentation
func DialWithExplicitTLS(tlsConfig *tls.Config) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithDebugOutput returns a DialOption that configures the ServerConn to write to the Writer
// everything it reads from the server
func DialWithDebugOutput(w io.Writer) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialWithDialFunc returns a DialOption that configures the ServerConn to use the
// specified function to establish both control and data connections
//
// If used together with the DialWithNetConn option, the DialWithNetConn
// takes precedence for the control connection, while data connections will
// be established using function specified with the DialWithDialFunc option
func DialWithDialFunc(f func(network, address string) (net.Conn, error)) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

func (o *dialOptions) wrapConn(netConn net.Conn) io.ReadWriteCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadWriteCloser)
}

func (o *dialOptions) wrapStream(rd io.ReadCloser) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

// Connect is an alias to Dial, for backward compatibility
//
// Deprecated: Use [Dial] instead
func Connect(addr string) (*ServerConn, error) {
	_ = "STUB: not implemented"

	// DialTimeout initializes the connection to the specified ftp server address.
	//
	// Deprecated: Use [Dial] with [DialWithTimeout] option instead
	return nil, nil
}

func DialTimeout(addr string, timeout time.Duration) (*ServerConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Login authenticates the client with specified user and password.
//
// "anonymous"/"anonymous" is a common user/password scheme for FTP servers
// that allows anonymous read-only accounts.
func (c *ServerConn) Login(user, password string) error { _ = "STUB: not implemented"; return nil }

// Probe features

// Switch to binary mode

// Switch to UTF-8

// If using implicit TLS, make data connections also use TLS

// authTLS upgrades the connection to use TLS
func (c *ServerConn) authTLS() error { _ = "STUB: not implemented"; return nil }

// feat issues a FEAT FTP command to list the additional commands supported by
// the remote FTP server.
// FEAT is described in RFC 2389
func (c *ServerConn) feat() error { _ = "STUB: not implemented"; return nil }

// The server does not support the FEAT command. This is not an
// error: we consider that there is no additional feature.

// setUTF8 issues an "OPTS UTF8 ON" command.
func (c *ServerConn) setUTF8() error { _ = "STUB: not implemented"; return nil }

// Workaround for FTP servers, that does not support this option.

// The ftpd "filezilla-server" has FEAT support for UTF8, but always returns
// "202 UTF8 mode is always enabled. No need to send this command." when
// trying to use it. That's OK

// epsv issues an "EPSV" command to get a port number for a data connection.
func (c *ServerConn) epsv() (port int, err error) { _ = "STUB: not implemented"; return 0, nil }

// pasv issues a "PASV" command to get a port number for a data connection.
func (c *ServerConn) pasv() (host string, port int, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// PASV response format : 227 Entering Passive Mode (h1,h2,h3,h4,p1,p2).

// We have to split the response string

// Let's compute the port number

// Recompose port

// Make the IP address to connect to

func isBogusDataIP(cmdIP, dataIP net.IP) bool {
	_ = "STUB: not implemented"
	// Logic stolen from lftp (https://github.com/lavv17/lftp/blob/d67fc14d085849a6b0418bb3e912fea2e94c18d1/src/ftpclass.cc#L769)
	return false
}

// getDataConnPort returns a host, port for a new data connection
// it uses the best available method to do so
func (c *ServerConn) getDataConnPort() (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// if there is an error, skip EPSV for the next attempts

// openDataConn creates a new FTP data connection.
func (c *ServerConn) openDataConn() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// We don't use tls.DialWithDialer here (which does Dial, create
// the Client and then do the Handshake) because it seems to
// hang with some FTP servers, namely proftpd and pureftpd.
//
// Instead we do Dial, create the Client and wait for the first
// Read or Write to trigger the Handshake.
//
// This means that if we are uploading a zero sized file, we
// need to make sure we do the Handshake explicitly as Write
// won't have been called. This is done in StorFrom().
//
// See: https://github.com/jlaffaye/ftp/issues/282

// cmd is a helper function to execute a command and check for the expected FTP
// return code
func (c *ServerConn) cmd(expected int, format string, args ...interface{}) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

// cmdDataConnFrom executes a command which require a FTP data connection.
// Issues a REST FTP command to specify the number of bytes to skip for the transfer.
func (c *ServerConn) cmdDataConnFrom(offset uint64, format string, args ...interface{}) (net.Conn, error) {
	_ = "STUB: not implemented"
	// If server requires PRET send the PRET command to warm it up
	// See: https://tools.ietf.org/html/draft-dd-pret-00
	return *new(net.Conn), nil
}

// Type switches the transfer mode for the connection.
func (c *ServerConn) Type(transferType TransferType) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// NameList issues an NLST FTP command.
func (c *ServerConn) NameList(path string) (entries []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List issues a LIST FTP command.
func (c *ServerConn) List(path string) (entries []*Entry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetEntry issues a MLST FTP command which retrieves one single Entry using the
// control connection. The returnedEntry will describe the current directory
// when no path is given.
func (c *ServerConn) GetEntry(path string) (entry *Entry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The expected reply will look something like:
//
//    250-File details
//     Type=file;Size=1024;Modify=20220813133357; path
//    250 End
//
// Multiple lines are allowed though, so it can also be in the form:
//
//    250-File details
//     Type=file;Size=1024; path
//     Modify=20220813133357; path
//    250 End

// lines must be a multi-line message with a length of 3 or more, and we
// don't care about the first and last line

// According to RFC 3659, the entry lines must start with a space when passed over the
// control connection. Some servers don't seem to add that space though. Both forms are
// accepted here.

// Some severs seem to send a blank line at the end which we ignore

// IsTimePreciseInList returns true if client and server support the MLSD
// command so List can return time with 1-second precision for all files.
func (c *ServerConn) IsTimePreciseInList() bool { _ = "STUB: not implemented"; return false }

// ChangeDir issues a CWD FTP command, which changes the current directory to
// the specified path.
func (c *ServerConn) ChangeDir(path string) error { _ = "STUB: not implemented"; return nil }

// ChangeDirToParent issues a CDUP FTP command, which changes the current
// directory to the parent directory.  This is similar to a call to ChangeDir
// with a path set to "..".
func (c *ServerConn) ChangeDirToParent() error { _ = "STUB: not implemented"; return nil }

// CurrentDir issues a PWD FTP command, which Returns the path of the current
// directory.
func (c *ServerConn) CurrentDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

// FileSize issues a SIZE FTP command, which Returns the size of the file
func (c *ServerConn) FileSize(path string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// GetTime issues the MDTM FTP command to obtain the file modification time.
// It returns a UTC time.
func (c *ServerConn) GetTime(path string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// IsGetTimeSupported allows library callers to check in advance that they
// can use GetTime to get file time.
func (c *ServerConn) IsGetTimeSupported() bool { _ = "STUB: not implemented"; return false }

// SetTime issues the MFMT FTP command to set the file modification time.
// Also it can use a non-standard form of the MDTM command supported by
// the VsFtpd server instead of MFMT for the same purpose.
// See "mdtm_write" in https://security.appspot.com/vsftpd/vsftpd_conf.html
func (c *ServerConn) SetTime(path string, t time.Time) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// IsSetTimeSupported allows library callers to check in advance that they
// can use SetTime to set file time.
func (c *ServerConn) IsSetTimeSupported() bool { _ = "STUB: not implemented"; return false }

// Retr issues a RETR FTP command to fetch the specified file from the remote
// FTP server.
//
// The returned ReadCloser must be closed to cleanup the FTP data connection.
func (c *ServerConn) Retr(path string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil,

		// RetrFrom issues a RETR FTP command to fetch the specified file from the remote
		// FTP server, the server will not send the offset first bytes of the file.
		//
		// The returned ReadCloser must be closed to cleanup the FTP data connection.
		nil
}

func (c *ServerConn) RetrFrom(path string, offset uint64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stor issues a STOR FTP command to store a file to the remote FTP server.
// Stor creates the specified file with the content of the io.Reader.
//
// Hint: io.Pipe() can be used if an io.Writer is required.
func (c *ServerConn) Stor(path string, r io.Reader) error { _ = "STUB: not implemented"; return nil }

// checkDataShut reads the "closing data connection" status from the
// control connection. It is called after transferring a piece of data
// on the data connection during which the control connection was idle.
// This may result in the idle timeout triggering on the control connection
// right when we try to read the response.
// The ShutTimeout dial option will rescue here. It will nudge the control
// connection deadline right before checking the data closing status.
func (c *ServerConn) checkDataShut() error { _ = "STUB: not implemented"; return nil }

// StorFrom issues a STOR FTP command to store a file to the remote FTP server.
// Stor creates the specified file with the content of the io.Reader, writing
// on the server will start at the given file offset.
//
// Hint: io.Pipe() can be used if an io.Writer is required.
func (c *ServerConn) StorFrom(path string, r io.Reader, offset uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// if the upload fails we still need to try to read the server
// response otherwise if the failure is not due to a connection problem,
// for example the server denied the upload for quota limits, we miss
// the response and we cannot use the connection to send other commands.

// If we wrote no bytes and got no error, make sure we call
// tls.Handshake on the connection as it won't get called
// unless Write() is called. (See comment in openDataConn()).
//
// ProFTP doesn't like this and returns "Unable to build data
// connection: Operation not permitted" when trying to upload
// an empty file without this.

// Append issues a APPE FTP command to store a file to the remote FTP server.
// If a file already exists with the given path, then the content of the
// io.Reader is appended. Otherwise, a new file is created with that content.
//
// Hint: io.Pipe() can be used if an io.Writer is required.
func (c *ServerConn) Append(path string, r io.Reader) error { _ = "STUB: not implemented"; return nil }

// Rename renames a file on the remote FTP server.
func (c *ServerConn) Rename(from, to string) error { _ = "STUB: not implemented"; return nil }

// Delete issues a DELE FTP command to delete the specified file from the
// remote FTP server.
func (c *ServerConn) Delete(path string) error { _ = "STUB: not implemented"; return nil }

// RemoveDirRecur deletes a non-empty folder recursively using
// RemoveDir and Delete
func (c *ServerConn) RemoveDirRecur(path string) error { _ = "STUB: not implemented"; return nil }

// MakeDir issues a MKD FTP command to create the specified directory on the
// remote FTP server.
func (c *ServerConn) MakeDir(path string) error { _ = "STUB: not implemented"; return nil }

// RemoveDir issues a RMD FTP command to remove the specified directory from
// the remote FTP server.
func (c *ServerConn) RemoveDir(path string) error { _ = "STUB: not implemented"; return nil }

// Walk prepares the internal walk function so that the caller can begin traversing the directory
func (c *ServerConn) Walk(root string) *Walker { _ = "STUB: not implemented"; return nil }

// NoOp issues a NOOP FTP command.
// NOOP has no effects and is usually used to prevent the remote FTP server to
// close the otherwise idle connection.
func (c *ServerConn) NoOp() error { _ = "STUB: not implemented"; return nil }

// Logout issues a REIN FTP command to logout the current user.
func (c *ServerConn) Logout() error { _ = "STUB: not implemented"; return nil }

// Quit issues a QUIT FTP command to properly close the connection from the
// remote FTP server.
func (c *ServerConn) Quit() error { _ = "STUB: not implemented"; return nil }

// Read implements the io.Reader interface on a FTP data connection.
func (r *Response) Read(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Close implements the io.Closer interface on a FTP data connection.
		// After the first call, Close will do nothing and return nil.
		nil
}

func (r *Response) Close() error { _ = "STUB: not implemented"; return nil }

// SetDeadline sets the deadlines associated with the connection.
func (r *Response) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// String returns the string representation of EntryType t.
func (t EntryType) String() string { _ = "STUB: not implemented"; return "" }
