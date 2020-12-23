package errors

import "fmt"

// Er represents cusstom error
type Er struct {
	Status int
	Err    error
}

func (e *Er) Error() string {
	return e.Err.Error()
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

// Status codes
const (
	// OK Return on Success
	OK = iota + 1
	// CANCELLED The operation was cancelled, typically by the caller.

	CANCELLED
	// UNKNOWN For example, this error may be returned when a Status value received from another address space
	// belongs to an error-space that isn't known in this address space. Also errors raised by APIs that don't
	// return enough error information may be converted to this error.

	UNKNOWN
	// INVALID_ARGUMENT the client specified an invalid argument.

	INVALID_ARGUMENT
	// DEADLINE_EXCEEDED The deadline expired before the operation could complete.

	DEADLINE_EXCEEDED
	// NOT_FOUND Some requested entity was not found.

	NOT_FOUND
	// ALREADY_EXISTS The entity that a client attempted to create already exists.

	ALREADY_EXISTS
	// PERMISSION_DENIED The caller does not have permission to execute the specified operation.

	PERMISSION_DENIED
	// RESOURCE_EXHAUSTED Some resource has been exhausted.

	RESOURCE_EXHAUSTED
	// FAILED_PRECONDITION The operation was rejected because the system is not in a state required for the operation’s execution.

	FAILED_PRECONDITION
	// ABORTED The operation was aborted.

	ABORTED
	// OUT_OF_RANGE The operation was attempted past the valid range.

	OUT_OF_RANGE
	// UNIMPLEMENTED The operation is not implemented or is not supported/enabled in this service.

	UNIMPLEMENTED
	// INTERNAL Internal errors.

	INTERNAL
	// UNAVAILABLE The service is currently unavailable.

	UNAVAILABLE
	// DATA_LOSS Unrecoverable data loss or corruption.

	DATA_LOSS
	// UNAUTHENTICATED The request does not have valid authentication credentials for the operation.
	UNAUTHENTICATED
)
