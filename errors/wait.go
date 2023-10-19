package errors

import (
	"errors"
	"fmt"
)

var ErrBadWaitState = errors.New("encountered bad state while waiting for resource")

type BadWaitState struct {
	State string
}

func (e BadWaitState) Error() string {
	return fmt.Sprintf("%s: %s", ErrBadWaitState.Error(), e.State)
}
func (e BadWaitState) Unwrap() error {
	return ErrBadWaitState
}

func NewBadWaitStateError(state string) BadWaitState {
	return BadWaitState{
		State: state,
	}
}
