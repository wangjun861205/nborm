package nborm

import (
	"fmt"
	"regexp"
)

type validateErrorCode int

const (
	VErrCannotBeNull validateErrorCode = iota + 1
	VErrRegex
	VErrOverMaxLength
)

type ValidateError struct {
	code    validateErrorCode
	message string
}

func NewValidateError(code validateErrorCode, message string) *ValidateError {
	return &ValidateError{code, message}
}

func (e *ValidateError) Error() string {
	return e.message
}

type modelValidator interface {
	Validate(m Model) error
}

type fieldValidator interface {
	Validate(f Field) error
}

func ValidateNull(f Field) error {
	if f.IsNull() {
		return NewValidateError(VErrCannotBeNull, fmt.Sprintf("field cannot be null (col: %s)", f.rawFullColName()))
	}
	return nil
}

func ValidateRegex(f *String, regex string) error {
	re := regexp.MustCompile(regex)
	if !re.MatchString(f.AnyValue()) {
		return NewValidateError(VErrCannotBeNull, fmt.Sprintf("field value not match regexp (col: %s, regex: %s)", f.rawFullColName(), regex))
	}
	return nil
}

func ValidateMaxLenght(f *String, length int) error {
	if len([]rune(f.AnyValue())) > length {
		return NewValidateError(VErrOverMaxLength, fmt.Sprintf("field string length is over max length (col: %s, length: %d)", f.rawFullColName(), length))
	}
	return nil
}
