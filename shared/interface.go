package shared

import "context"

type IQueryHandler[Q, R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}

type ICommandHandler[C any] interface {
	Handle(ctx context.Context, command C) error
}

type ICommandResultHandler[C any, R any] interface {
	Handle(ctx context.Context, command C) (R, error)
}