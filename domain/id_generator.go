package domain

type TaskIdGenerator interface {
	Generate() (TaskIdentifier, error)
}

type TaskIdentifier struct {
	Value string
}

func (ti *TaskIdentifier) String() string {
	return "TaskIdentifier: {Value: " + ti.Value + "}"
}
