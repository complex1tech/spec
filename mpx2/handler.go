package mpx

import (
	"github.com/basecomplextech/baselibrary/async"
	"github.com/basecomplextech/baselibrary/status"
)

// Handler is a server channel handler.
type Handler interface {
	// HandleChannel handles an incoming channel.
	HandleChannel(ctx async.Context, ch Channel) status.Status
}

// HandleFunc is a type adapter to allow use of ordinary functions as channel handlers.
type HandleFunc func(ctx async.Context, ch Channel) status.Status

// HandleChannel handles an incoming channel.
func (f HandleFunc) HandleChannel(ctx async.Context, ch Channel) status.Status {
	return f(ctx, ch)
}
