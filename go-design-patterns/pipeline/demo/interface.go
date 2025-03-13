package main

import "context"

type ISource interface {
	Process(ctx context.Context) (<-chan any, error)
}

type IProcessor interface {
	Process(ctx context.Context, params any) (any, error)
}

type ISink interface {
	Process(ctx context.Context, params any) error
}

// TimerSource source with timer
type TimerSource struct {
}

func main() {

	return
}
