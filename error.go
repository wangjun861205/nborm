package nborm

import "fmt"

type errorCode int

const (
	ErrCodeIO errorCode = iota + 1
	ErrCodeSerialize
	ErrCodeInvalidField
	ErrCodeInvalidValueType
	ErrCodeExecute
	ErrCodeInvalidValueFormat
)

type Error struct {
	code errorCode
	msg  string
	err  error
}

func newErr(code errorCode, msg string, err error) *Error {
	return &Error{code, msg, err}
}

func (e *Error) Error() string {
	return fmt.Sprintf("nborm error: %s(code: %d) caused by %v", e.msg, e.code, e.err)
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Is(code errorCode) bool {
	return e.code == code
}

func IsErr(err error, code errorCode) bool {
	if e, ok := err.(*Error); !ok {
		return false
	} else {
		return e.code == code
	}
}
