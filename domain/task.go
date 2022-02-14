package domain

import (
	"log"
	"strconv"
)

const createdState = "CREATED"

type TaskIdentifier struct {
	Value string
}

func (ti TaskIdentifier) String() string {
	return "TaskIdentifier: {Valuse: " + ti.Value + "}"
}

type Task struct {
	Identifier TaskIdentifier
	N          Param
	M          Param
	Result     int //Init with -1. Result of a calculation cannot be negative
	State      string
}

func CreateTask(n1 Param, n2 Param, generate func() (TaskIdentifier, error)) (Task, TaskCreatedEvent, error) {
	id, err := generate()

	if err != nil {
		log.Fatalf("Id generation error, %s", err)
		return Task{}, TaskCreatedEvent{}, err
	}

	task := Task{id, n1, n2, -1, createdState}
	event := TaskCreatedEvent{taskId: id}

	return task, event, nil
}

func (t Task) String() string {
	return "Task: {Identifier: " + t.Identifier.Value + " , N: " + strconv.Itoa(t.N.Value) + " , M: " + strconv.Itoa(t.M.Value) + " , Result: " + strconv.Itoa(t.Result) + " , State: " + t.State + "}"
}
