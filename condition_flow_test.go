package workflow

import (
	"context"
	"fmt"
	"testing"
)

func TestConditionFlow_Execute(t *testing.T) {
	ctx := context.Background()
	workflow := NewConditionFlow()
	workflow.
		Named("bar").
		ExecuteOne(NewPrintMessageWork("foo")).
		When(NewTimesPredicate(3)).
		Then(NewPrintMessageWork("spam")).
		Otherwise(NewPrintMessageWork("die"))

	engine := NewEngine()
	workReport := engine.Run(workflow, ctx)
	fmt.Println(workReport.Status())
}
