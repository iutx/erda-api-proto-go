// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cmp_alert.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetAlertConditionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAlertConditionsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAlertConditionsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAlertConditionsRequestMultiError, or nil if none found.
func (m *GetAlertConditionsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAlertConditionsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ScopeType

	if len(errors) > 0 {
		return GetAlertConditionsRequestMultiError(errors)
	}

	return nil
}

// GetAlertConditionsRequestMultiError is an error wrapping multiple validation
// errors returned by GetAlertConditionsRequest.ValidateAll() if the
// designated constraints aren't met.
type GetAlertConditionsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAlertConditionsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAlertConditionsRequestMultiError) AllErrors() []error { return m }

// GetAlertConditionsRequestValidationError is the validation error returned by
// GetAlertConditionsRequest.Validate if the designated constraints aren't met.
type GetAlertConditionsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAlertConditionsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAlertConditionsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAlertConditionsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAlertConditionsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAlertConditionsRequestValidationError) ErrorName() string {
	return "GetAlertConditionsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAlertConditionsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAlertConditionsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAlertConditionsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAlertConditionsRequestValidationError{}

// Validate checks the field values on GetAlertConditionsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAlertConditionsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAlertConditionsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAlertConditionsResponseMultiError, or nil if none found.
func (m *GetAlertConditionsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAlertConditionsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetAlertConditionsResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetAlertConditionsResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAlertConditionsResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetAlertConditionsResponseMultiError(errors)
	}

	return nil
}

// GetAlertConditionsResponseMultiError is an error wrapping multiple
// validation errors returned by GetAlertConditionsResponse.ValidateAll() if
// the designated constraints aren't met.
type GetAlertConditionsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAlertConditionsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAlertConditionsResponseMultiError) AllErrors() []error { return m }

// GetAlertConditionsResponseValidationError is the validation error returned
// by GetAlertConditionsResponse.Validate if the designated constraints aren't met.
type GetAlertConditionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAlertConditionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAlertConditionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAlertConditionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAlertConditionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAlertConditionsResponseValidationError) ErrorName() string {
	return "GetAlertConditionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAlertConditionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAlertConditionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAlertConditionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAlertConditionsResponseValidationError{}

// Validate checks the field values on GetAlertConditionsValueRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAlertConditionsValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAlertConditionsValueRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetAlertConditionsValueRequestMultiError, or nil if none found.
func (m *GetAlertConditionsValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAlertConditionsValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetConditions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetAlertConditionsValueRequestValidationError{
						field:  fmt.Sprintf("Conditions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetAlertConditionsValueRequestValidationError{
						field:  fmt.Sprintf("Conditions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAlertConditionsValueRequestValidationError{
					field:  fmt.Sprintf("Conditions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetAlertConditionsValueRequestMultiError(errors)
	}

	return nil
}

// GetAlertConditionsValueRequestMultiError is an error wrapping multiple
// validation errors returned by GetAlertConditionsValueRequest.ValidateAll()
// if the designated constraints aren't met.
type GetAlertConditionsValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAlertConditionsValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAlertConditionsValueRequestMultiError) AllErrors() []error { return m }

// GetAlertConditionsValueRequestValidationError is the validation error
// returned by GetAlertConditionsValueRequest.Validate if the designated
// constraints aren't met.
type GetAlertConditionsValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAlertConditionsValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAlertConditionsValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAlertConditionsValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAlertConditionsValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAlertConditionsValueRequestValidationError) ErrorName() string {
	return "GetAlertConditionsValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAlertConditionsValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAlertConditionsValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAlertConditionsValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAlertConditionsValueRequestValidationError{}

// Validate checks the field values on GetAlertConditionsValueResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAlertConditionsValueResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAlertConditionsValueResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetAlertConditionsValueResponseMultiError, or nil if none found.
func (m *GetAlertConditionsValueResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAlertConditionsValueResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetAlertConditionsValueResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetAlertConditionsValueResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAlertConditionsValueResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetAlertConditionsValueResponseMultiError(errors)
	}

	return nil
}

// GetAlertConditionsValueResponseMultiError is an error wrapping multiple
// validation errors returned by GetAlertConditionsValueResponse.ValidateAll()
// if the designated constraints aren't met.
type GetAlertConditionsValueResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAlertConditionsValueResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAlertConditionsValueResponseMultiError) AllErrors() []error { return m }

// GetAlertConditionsValueResponseValidationError is the validation error
// returned by GetAlertConditionsValueResponse.Validate if the designated
// constraints aren't met.
type GetAlertConditionsValueResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAlertConditionsValueResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAlertConditionsValueResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAlertConditionsValueResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAlertConditionsValueResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAlertConditionsValueResponseValidationError) ErrorName() string {
	return "GetAlertConditionsValueResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAlertConditionsValueResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAlertConditionsValueResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAlertConditionsValueResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAlertConditionsValueResponseValidationError{}
