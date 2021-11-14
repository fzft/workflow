package workflow

import (
	"context"
	"testing"
)

func TestParallelFlow_Execute(t *testing.T) {
	ctx := context.Background()
	workflow := NewParallelFlow()
	workflow.
		Named("bar").
		ExecuteAll(NewPrintMessageWork("hello"),
				NewPrintMessageWork("foo")).
		With(NewParallelFlowExecutor()).
		Build()
	engine := NewEngine()
	engine.Run(workflow, ctx)
	//fmt.Println(workReport.Status())
}