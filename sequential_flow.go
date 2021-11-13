package workflow

import "context"

type NameStep interface {
	Named(name string) ExecuteStep
}

type ThenStep interface {
	ThenOne(nextWork Work) ThenStep
	ThenUnits(nextWorkUnits []Work) ThenStep
	Build() WorkFlow
}

type ExecuteStep interface {
	ExecuteOne(initialWork Work) ThenStep
	ExecuteUnits(initialWorkUnits []Work) ThenStep
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


func (s *SequentialFlow) Named(name string) ExecuteStep {
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

func (s *SequentialFlow) ExecuteOne(w Work) ThenStep {
	s.WorkUints = append(s.WorkUints, w)
	return s
}

func (s *SequentialFlow) ExecuteUnits(ws []Work) ThenStep {
	for _, w := range ws {
		s.WorkUints = append(s.WorkUints, w)
	}
	return s
}

func (s *SequentialFlow) ThenOne(initialWork Work) ThenStep {
	s.WorkUints = append(s.WorkUints, initialWork)
	return s
}

func (s *SequentialFlow) ThenUnits(nextWorkUnits []Work) ThenStep {
	for _, w := range nextWorkUnits {
		s.WorkUints = append(s.WorkUints, w)
	}
	return s
}

func (s *SequentialFlow) Build() WorkFlow {
	return s
}
