***

<div align="center">
    <b><em>Easy Flows</em></b><br>
    The simple, stupid workflow engine for golang, Enlightened by Java Project;
</div>

<div align="center">
</div>

## How does it work?

```go

	ctx := context.Background()
	workflow := NewSequentialFlow()

	workflow.Named("seq").
		ExecuteOne(NewRepeatFlow().Named("print foo 3 times").
			Repeat(NewPrintMessageWork("foo")).
			Times(3).
			Build()).
		ThenOne(NewParallelFlow().
			Named("print hello and world in parallel").
			ExecuteAll(NewPrintMessageWork("hello"), NewPrintMessageWork("world")).
			With(NewParallelFlowExecutor()).
			Build()).
		Build()


	engine := NewEngine()
	workReport := engine.Run(workflow, ctx)
	fmt.Println(workReport.Status())
```