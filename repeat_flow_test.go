package workflow

import (
	"context"
	"fmt"
	"testing"
)

func TestRepeatFlow_Execute(t *testing.T) {
	ctx := context.Background()
	workflow := NewRepeatFlow()
	workflow.
		Named("bar").
		Repeat(NewPrintMessageWork("foo")).
		Times(3).
		Build()

	engine := NewEngine()
	workReport := engine.Run(workflow, ctx)
	fmt.Println(workReport.Status())
}