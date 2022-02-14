package domain

type TaskCreatedEvent struct {
	taskId TaskIdentifier
}

func (tce TaskCreatedEvent) String() string {
	return "TaskCreatedEvent: {taskId: " + tce.taskId.String() + "}"
}
