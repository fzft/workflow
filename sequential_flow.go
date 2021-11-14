package workflow

import "context"

type SequentialNameStep interface {
	Named(name string) SequentialExecuteStep
}

type SequentialThenStep interface {
	ThenOne(nextWork Work) SequentialThenStep
	ThenUnits(nextWorkUnits []Work) SequentialThenStep
	Build() WorkFlow
}

type SequentialExecuteStep interface {
	ExecuteOne(initialWork Work) SequentialThenStep
	ExecuteUnits(initialWorkUnits []Work) SequentialThenStep
}

type SequentialFlow struct {
	WorkUints []Work
	name      string
}

func NewSequentialFlow() *SequentialFlow {
	return &SequentialFlow{}
}

func (s *SequentialFlow) Name() string {
	return s.name
}


func (s *SequentialFlow) Named(name string) SequentialExecuteStep {
	s.name = name
	return s
}

func (s *SequentialFlow) Execute(ctx context.Context) WorkReport {
	var wr WorkReport
	for _, work := range s.WorkUints {
		wr = work.Execute(ctx)
		if wr.Status() == FAILED {
			break
		}
	}
	return wr
}

func (s *SequentialFlow) ExecuteOne(w Work) SequentialThenStep {
	s.WorkUints = append(s.WorkUints, w)
	return s
}

func (s *SequentialFlow) ExecuteUnits(ws []Work) SequentialThenStep {
	for _, w := range ws {
		s.WorkUints = append(s.WorkUints, w)
	}
	return s
}

func (s *SequentialFlow) ThenOne(initialWork Work) SequentialThenStep {
	s.WorkUints = append(s.WorkUints, initialWork)
	return s
}

func (s *SequentialFlow) ThenUnits(nextWorkUnits []Work) SequentialThenStep {
	for _, w := range nextWorkUnits {
		s.WorkUints = append(s.WorkUints, w)
	}
	return s
}

func (s *SequentialFlow) Build() WorkFlow {
	return s
}
