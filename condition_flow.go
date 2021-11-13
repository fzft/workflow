package workflow

import "context"

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



