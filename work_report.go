package workflow

import (
	"context"
	"sync/atomic"
)

type WorkReport interface {
	Status() WorkStatus
	WorkContext() context.Context
	Error() error
}

type DefaultWorkReport struct {
	status WorkStatus
	ctx    context.Context
	err    error
}

func NewDefaultWorkReport(status WorkStatus, ctx context.Context) DefaultWorkReport {
	return DefaultWorkReport{
		status: status,
		ctx:    ctx,
	}
}

func (dwr DefaultWorkReport) Status() WorkStatus {
	return dwr.status
}

func (dwr DefaultWorkReport) WorkContext() context.Context {
	return dwr.ctx
}

func (dwr DefaultWorkReport) Error() error {
	return dwr.err
}

type WorkReportPredicate interface {
	Apply(report WorkReport) bool
}

type TimesPredicate struct {
	times   uint64
	counter *uint64
}

func NewTimesPredicate(times uint64) *TimesPredicate {
	return &TimesPredicate{
		times:   times,
		counter: new(uint64),
	}
}

func (p *TimesPredicate) Apply(report WorkReport) bool {
	b := atomic.LoadUint64(p.counter) != p.times
	atomic.AddUint64(p.counter, 1)
	return b
}

type ParallelFlowReport struct {
	ctx     context.Context
	reports []WorkReport
}

func NewParallelFlowReport(reports []WorkReport) ParallelFlowReport {
	return ParallelFlowReport{
		reports: reports,
	}
}

func (r ParallelFlowReport) Status() WorkStatus {
	for _, report := range r.reports {
		if report.Status() == FAILED {
			return FAILED
		}
	}
	return COMPLETED
}

func (r ParallelFlowReport) WorkContext() context.Context {
	return r.ctx
}

func (r ParallelFlowReport) Error() error {
	for _, report := range r.reports {
		if err := report.Error(); err != nil {
			return report.Error()
		}
	}
	return nil
}

