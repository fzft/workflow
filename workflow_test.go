package workflow

import (
	"context"
	"fmt"
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
