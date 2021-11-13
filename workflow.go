package workflow

import (
	"context"
)

type WorkStatus int64

const (
	IDLE      WorkStatus = iota
	RUNNING   WorkStatus = iota + 1
	COMPLETED WorkStatus = iota + 2
	FAILED    WorkStatus = iota + 3
)

type Work interface {
	Name() string
	Execute(ctx context.Context) WorkReport
}

type WorkFlow interface {
	Work
	Build() WorkFlow
}

