package domain

import (
	"errors"
	"fmt"
	"strconv"
)

type Param struct {
	Value int
}

func CreateParam(v int) (Param, error) {
	if v < 0 {
		msg := fmt.Sprintf("Input param: {%d} cannot be negative", v)
		return Param{}, errors.New(msg)
	}

	return Param{Value: v}, nil
}

func (p Param) String() string {
	return "Param: {Value: " + strconv.Itoa(p.Value) + "}"
}
