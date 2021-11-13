package workflow

import (
	"context"
	"fmt"
	"testing"
)

func TestSequentialFlow_Execute(t *testing.T) {
	ctx := context.Background()
	workflow := NewSequentialFlow()
	workflow.
		Named("bar").
		ExecuteOne(NewPrintMessageWork("foo")).
		ThenOne(NewPrintMessageWork("span")).
		Build()

	engine := NewEngine()
	workReport := engine.Run(workflow, ctx)
	fmt.Println(workReport.Status())
}
