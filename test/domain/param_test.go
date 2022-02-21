package domain

import (
	"github.com/conacry/lcm/domain"
	"reflect"
	"testing"
)

const validValue = 10

func TestCreateParam_ValidValue_CreateParam(t *testing.T) {
	expectParam := domain.Param{Value: validValue}

	actualParam, err := domain.CreateParam(validValue)

	if err != nil {
		t.Errorf("Param creation has finished with error")
	}

	if !reflect.DeepEqual(expectParam, actualParam) {
		t.Errorf("expected (%s), got (%s)", expectParam, actualParam)
	}
}
