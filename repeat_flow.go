package workflow

import "context"


type RepeatNameStep interface {
	Named(name string) RepeatStep
}

type RepeatStep interface {
	Repeat(work Work) UntilStep
}

type UntilStep interface {
	Until(p WorkReportPredicate) UntilStep
	Times(times uint64) UntilStep
	Build() WorkFlow
}

type RepeatFlow struct {
	name                string
	work                Work
	workReportPredicate WorkReportPredicate
}

func NewRepeatFlow() *RepeatFlow {
	return &RepeatFlow{

	}
}

func (f *RepeatFlow) Name() string {
	return f.name
}

func (f *RepeatFlow) Execute(ctx context.Context) WorkReport {
	var wr WorkReport
	for f.workReportPredicate.Apply(wr) {
		wr = f.work.Execute(ctx)
	}
	return wr
}

func (f *RepeatFlow) Named(name string) RepeatStep {
	f.name = name
	return f
}

func (f *RepeatFlow) Repeat(work Work) UntilStep {
	f.work = work
	return f
}

func (f *RepeatFlow) Until(p WorkReportPredicate) UntilStep {
	f.workReportPredicate = p
	return f
}

func (f *RepeatFlow) Times(times uint64) UntilStep {
	f.Until(NewTimesPredicate(times))
	return f
}

func (f *RepeatFlow) Build() WorkFlow {
	return f
}


