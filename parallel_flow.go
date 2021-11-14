package workflow

import "context"

type ParallelNameStep interface {
	Named(name string) ParallelExecuteStep
}

type ParallelExecuteStep interface {
	ExecuteAll(works ...Work) ParallelWithStep
}

type ParallelWithStep interface {
	With(executor FlowExecutor) ParallelWithStep
	Build() WorkFlow
}

type ParallelFlow struct {
	name         string
	workUnits    []Work
	workExecutor FlowExecutor
}

func NewParallelFlow() *ParallelFlow {
	return &ParallelFlow{
	}
}

func (f *ParallelFlow) Execute(ctx context.Context) WorkReport {
	wrs := f.workExecutor.Execute(ctx, f.workUnits)
	pfr := NewParallelFlowReport(wrs)
	return pfr
}

func (f *ParallelFlow) Name() string {
	return f.name
}

func (f *ParallelFlow) Named(name string) ParallelExecuteStep {
	f.name = name
	return f
}

func (f *ParallelFlow) ExecuteAll(works ...Work) ParallelWithStep {
	for _, w := range works {
		f.workUnits = append(f.workUnits, w)
	}
	return f
}

func (f *ParallelFlow) With(executor FlowExecutor) ParallelWithStep {
	f.workExecutor = executor
	return f
}

func (f *ParallelFlow) Build() WorkFlow {
	return f
}
