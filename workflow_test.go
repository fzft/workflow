package workflow

import (
	"context"
	"fmt"
	"testing"
)

type PrintMessageWork struct {
	name string
}

func NewPrintMessageWork(name string) *PrintMessageWork {
	return &PrintMessageWork{
		name: name,
	}
}

func (pmw *PrintMessageWork) Name() string {
	return pmw.name
}

func (pmw *PrintMessageWork) Execute(ctx context.Context) WorkReport {
	fmt.Printf("this is execute %s work\n", pmw.name)
	return NewDefaultWorkReport(COMPLETED, ctx)
}


func TestWorkFlow_Execute(t *testing.T) {
	ctx := context.Background()
	workflow := NewSequentialFlow()

	workflow.Named("seq").
		ExecuteOne(NewRepeatFlow().Named("print foo 3 times").Repeat(NewPrintMessageWork("foo")).Times(3).Build()).
		ThenOne(NewParallelFlow().Named("print hello and world in parallel").ExecuteAll(NewPrintMessageWork("hello"), NewPrintMessageWork("world")).With(NewParallelFlowExecutor()).Build()).
		Build()


	engine := NewEngine()
	workReport := engine.Run(workflow, ctx)
	fmt.Println(workReport.Status())
}
