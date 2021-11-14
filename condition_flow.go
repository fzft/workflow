package workflow

import "context"

type ConditionNameStep interface {
	Named(name string) ConditionExecuteStep
}

type ConditionExecuteStep interface {
	ExecuteOne(initialWorkUnit Work) ConditionWhenStep
}

type ConditionWhenStep interface {
	When(p WorkReportPredicate) ConditionThenStep
}

type ConditionThenStep interface {
	Then(w Work) ConditionOtherwiseStep
}

type ConditionOtherwiseStep interface {
	Otherwise(w Work) ConditionOtherwiseStep
	Build() WorkFlow
}

type ConditionFlow struct {
	name                   string
	initialWorkUnit        Work
	nextOnPredicateSuccess Work
	nextOnPredicateFailure Work
	workReportPredicate    WorkReportPredicate
}

func NewConditionFlow() *ConditionFlow {
	return &ConditionFlow{}
}

func (f *ConditionFlow) Execute(ctx context.Context) WorkReport {
	wr := f.initialWorkUnit.Execute(ctx)
	if f.workReportPredicate.Apply(wr) {
		wr = f.nextOnPredicateSuccess.Execute(ctx)
	} else {
		if f.nextOnPredicateFailure != nil {
			wr = f.nextOnPredicateFailure.Execute(ctx)
		}
	}
	return wr
}

func (f *ConditionFlow) Name() string {
	return f.name
}

func (f *ConditionFlow) Named(name string) ConditionExecuteStep {
	f.name = name
	return f
}

func (f *ConditionFlow) ExecuteOne(initialWorkUnit Work) ConditionWhenStep {
	f.initialWorkUnit = initialWorkUnit
	return f
}

func (f *ConditionFlow) When(p WorkReportPredicate) ConditionThenStep {
	f.workReportPredicate = p
	return f
}

func (f *ConditionFlow) Then(w Work) ConditionOtherwiseStep {
	f.nextOnPredicateSuccess = w
	return f
}

func (f *ConditionFlow) Otherwise(w Work) ConditionOtherwiseStep {
	f.nextOnPredicateFailure = w
	return f
}

func (f *ConditionFlow) Build() WorkFlow {
	return f
}
