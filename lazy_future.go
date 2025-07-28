package future

import "context"

type LazyFuture[T any] struct {
	n *node[T, any]
}

func Lazy[T any](fn func(ctx context.Context) (T, error)) *LazyFuture[T] {
	lf := &LazyFuture[T]{
		&node[T, any]{
			nodeState: &state[T]{},
			fn: &loadFunc[T]{
				fn: fn,
			},
		},
	}
	return lf
}

type node[T, P any] struct {
	nodeState *state[T]
	prev      *node[T, P]
	fn        *loadFunc[T]
}

type loadFunc[T any] struct {
	fn func(ctx context.Context) (T, error)

	mark uint32
}
