package rpc

import (
	"github.com/basecomplextech/baselibrary/alloc"
	"github.com/basecomplextech/baselibrary/status"
)

// Handler is an RPC handler.
type Handler interface {
	// Handle handles a request and returns its result and status.
	// Result is ignored if status is not OK.
	Handle(cancel <-chan struct{}, ch ServerChannel) (*alloc.Buffer, status.Status)
}

// HandleFunc is a type adapter to allow use of ordinary functions as RPC handlers.
type HandleFunc func(cancel <-chan struct{}, ch ServerChannel) (*alloc.Buffer, status.Status)

// Handle handles a request and returns its result and status.
// Result is ignored if status is not OK.
func (f HandleFunc) Handle(cancel <-chan struct{}, ch ServerChannel) (*alloc.Buffer, status.Status) {
	return f(cancel, ch)
}
