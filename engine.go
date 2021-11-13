package workflow

import "context"

type Engine interface {
	Run(flow WorkFlow, ctx context.Context) WorkReport
}

type EngineImpl struct {
}

func NewEngine() *EngineImpl {
	return &EngineImpl{}
}

func (e *EngineImpl) Run(flow WorkFlow, ctx context.Context) WorkReport {
	return flow.Execute(ctx)
}
