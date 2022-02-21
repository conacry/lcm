package domain

type TaskCreatedEvent struct {
	TaskId TaskIdentifier
}

func (tce TaskCreatedEvent) String() string {
	return "TaskCreatedEvent: {taskId: " + tce.TaskId.String() + "}"
}
