package domain

import (
	"reflect"
	"testing"
)

func mockGenerate() (TaskIdentifier, error) {
	return TaskIdentifier{Value: "identifier"}, nil
}

func TestCreateTask_ValidParam_CreateTask(t *testing.T) {
	n1 := Param{Value: 10}
	n2 := Param{Value: 20}

	wantId, _ := mockGenerate()
	wantTask := Task{
		Identifier: wantId,
		N:          n1,
		M:          n2,
		Result:     -1,
		State:      "CREATED",
	}

	wantEvent := TaskCreatedEvent{taskId: wantId}

	gotTask, gotEvent, err := CreateTask(n1, n2, mockGenerate)

	if !reflect.DeepEqual(wantTask, gotTask) {
		t.Errorf("expected (%s), got (%s)", wantTask, gotTask)
	}

	if !reflect.DeepEqual(wantEvent, gotEvent) {
		t.Errorf("expected (%s), got (%s)", wantEvent, gotEvent)
	}

	if err != nil {
		t.Errorf("Creation has finished with error")
	}
}
