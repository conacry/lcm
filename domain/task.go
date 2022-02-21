package domain

import (
	"log"
	"strconv"
)

const createdState = "CREATED"

type Task struct {
	Identifier TaskIdentifier
	N1         Param
	N2         Param
	Result     int //Init with -1. Result of a calculation cannot be negative
	State      string
}

func CreateTask(n1 Param, n2 Param, generator TaskIdGenerator) (*Task, *TaskCreatedEvent, error) {
	id, err := generator.Generate()

	if err != nil {
		log.Fatalf("Id generation error, %s", err)
		return nil, nil, err
	}

	task := Task{Identifier: id, N1: n1, N2: n2, Result: -1, State: createdState}
	event := TaskCreatedEvent{TaskId: id}

	return &task, &event, nil
}

func (t *Task) String() string {
	return "Task: {Identifier: " +
		t.Identifier.Value + " , N: " +
		strconv.Itoa(t.N1.Value) + " , M: " +
		strconv.Itoa(t.N2.Value) + " , Result: " +
		strconv.Itoa(t.Result) + " , State: " +
		t.State + "}"
}
