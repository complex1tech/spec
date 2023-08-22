package tcp

import (
	"errors"
	"io"
	"net"

	"github.com/basecomplextech/baselibrary/status"
)

const (
	codeError        status.Code = "tcp_error"
	codeConnClosed   status.Code = "tcp_conn_closed"
	codeStreamClosed status.Code = "tcp_stream_closed"
)

var (
	statusConnClosed   = status.New(codeConnClosed, "tcp connection closed")
	statusStreamClosed = status.New(codeStreamClosed, "tcp stream closed")
)

func tcpError(err error) status.Status {
	if err == nil {
		return status.OK
	}

	switch err {
	case io.EOF:
		return status.End
	case io.ErrUnexpectedEOF:
		return status.WrapError(err).WithCode(codeError)
	}

	if errors.Is(err, net.ErrClosed) {
		return status.WrapError(err).WithCode(codeConnClosed)
	}

	ne, ok := (err).(net.Error)
	switch {
	case !ok:
		return status.WrapError(err).WithCode(codeError)
	case ne.Timeout():
		return status.Timeout
	}
	return status.WrapError(err).WithCode(codeError)
}

func tcpErrorf(format string, args ...any) status.Status {
	return status.Newf(codeError, format, args...)
}
