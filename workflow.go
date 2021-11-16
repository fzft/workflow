package workflow

import (
	"context"
)

type WorkStatus int64

const (
	IDLE      WorkStatus = iota
	RUNNING
	COMPLETED
	FAILED
)

type Work interface {
	Name() string
	Execute(ctx context.Context) WorkReport
}

type WorkFlow interface {
	Work
	Build() WorkFlow
}

