package workflow

import (
	"context"
	"sync"
)

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

type FlowExecutor interface {
	Execute(ctx context.Context, workUnits []Work) []WorkReport
}

type ParallelFlowExecutor struct {
	mu sync.Mutex
}

func NewParallelFlowExecutor() *ParallelFlowExecutor {
	return &ParallelFlowExecutor{}
}

func (e *ParallelFlowExecutor) Execute(ctx context.Context, workUnits []Work) []WorkReport {
	wg := sync.WaitGroup{}
	wrs := make([]WorkReport, len(workUnits))
	var m sync.Map
	for i, w := range workUnits {
		wg.Add(1)
		go func(work Work, idx int) {
			defer wg.Done()
			wr := work.Execute(ctx)
			m.Store(idx, wr)
		}(w, i)
	}
	wg.Wait()
	m.Range(func(k, v interface{}) bool {
		idx := k.(int)
		wpv := v.(WorkReport)
		wrs[idx] = wpv
		return true
	})
	return wrs
}
