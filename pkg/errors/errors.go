package errors

import "fmt"

// Er represents cusstom error
type Er struct {
	Status int
	Err    error
}

func (e *Er) Error() string {
	return fmt.Sprintf("Status %d: Error %s", e.Status, e.Err.Error())
}

// GetStatus returns status from error
func (e *Er) GetStatus() int {
	return e.Status
}

// New create new custom error with status and error
func New(status int, err error) *Er {
	return &Er{
		Status: status,
		Err:    err,
	}
}

// NewMsgln create new custom error with status and msg
func NewMsgln(status int, msg interface{}) *Er {
	return &Er{
		Status: status,
		Err:    fmt.Errorf("%v", msg),
	}
}

// NewMsgf creates new custom error with status and formated msg
func NewMsgf(status int, format string, msg ...interface{}) *Er {
	return &Er{
		Status: status,
		Err:    fmt.Errorf(format, msg...),
	}
}
