package rpc

import (
	"github.com/basecomplextech/baselibrary/alloc"
	"github.com/basecomplextech/baselibrary/status"
	"github.com/basecomplextech/spec/mpx"
)

const (
	// CodeSkipResponse is used by oneway RPCs to skip sending a response.
	CodeSkipResponse status.Code = "skip_response"
)

type (
	// Context is an RPC context, which is an alias for mpx.Context.
	Context = mpx.Context

	// Options is RPC options, which are a type alias for mpx.Options.
	Options = mpx.Options
)

// SkipResponse instructs the server to skip a response for a oneway method.
var SkipResponse = status.Status{
	Code:    CodeSkipResponse,
	Message: "skip response for oneway method",
}

// Default returns default options.
func Default() Options {
	return mpx.Default()
}

// NewBuffer returns a new alloc.Buffer.
// The method is used in generated code.
func NewBuffer() alloc.Buffer {
	return alloc.NewBuffer()
}
