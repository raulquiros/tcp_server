package inmemory

import (
	"context"
	"github.com/raulquiros/tcp_server/kit/query"
)

// QueryBus is an in-memory implementation of the query.Bus.
type QueryBus struct {
	handlers map[query.Type]query.Handler
}

// NewQueryBus initializes a new instance of QueryBus.
func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

// Dispatch implements the query.Bus interface.
func (b *QueryBus) Dispatch(ctx context.Context, cmd query.Query) (query.Response, error) {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return query.Response{}, nil
	}

	resp, err := handler.Handle(ctx, cmd)
	if err != nil {
		panic(err)
	}

	return resp, nil
}

// Register implements the query.Bus interface.
func (b *QueryBus) Register(queries map[query.Type]query.Handler) {
	b.handlers = queries
}
