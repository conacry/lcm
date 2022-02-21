package domain

import (
	"github.com/conacry/lcm/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type IdGeneratorMock struct {
	mock.Mock
}

func (i *IdGeneratorMock) Generate() (domain.TaskIdentifier, error) {
	args := i.Called()
	return args.Get(0).(domain.TaskIdentifier), args.Error(1)
}

func TestCreateTask_ValidParam_CreateTask(t *testing.T) {
	n1 := domain.Param{Value: 10}
	n2 := domain.Param{Value: 20}

	generator := new(IdGeneratorMock)
	generator.On("Generate").Return(domain.TaskIdentifier{Value: "str"}, nil)

	wantTask := &domain.Task{
		Identifier: domain.TaskIdentifier{Value: "str"},
		N1:         n1,
		N2:         n2,
		Result:     -1,
		State:      "CREATED",
	}

	wantEvent := &domain.TaskCreatedEvent{TaskId: domain.TaskIdentifier{Value: "str"}}

	gotTask, gotEvent, err := domain.CreateTask(n1, n2, generator)

	a := assert.New(t)
	a.Equal(wantTask, gotTask, "expected (%s), got (%s)", wantTask, gotTask)
	a.Equal(wantEvent, gotEvent, "expected (%s), got (%s)", wantEvent, gotEvent)
	a.Nil(err, "Creation has finished with error")
}
