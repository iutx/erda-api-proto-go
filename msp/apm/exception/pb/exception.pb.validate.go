// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: exception.proto

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

// Validate checks the field values on GetExceptionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExceptionsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExceptionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExceptionsRequestMultiError, or nil if none found.
func (m *GetExceptionsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExceptionsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetStartTime() <= 0 {
		err := GetExceptionsRequestValidationError{
			field:  "StartTime",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEndTime() <= 0 {
		err := GetExceptionsRequestValidationError{
			field:  "EndTime",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetScopeID()) < 1 {
		err := GetExceptionsRequestValidationError{
			field:  "ScopeID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Debug

	if len(errors) > 0 {
		return GetExceptionsRequestMultiError(errors)
	}

	return nil
}

// GetExceptionsRequestMultiError is an error wrapping multiple validation
// errors returned by GetExceptionsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetExceptionsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExceptionsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExceptionsRequestMultiError) AllErrors() []error { return m }

// GetExceptionsRequestValidationError is the validation error returned by
// GetExceptionsRequest.Validate if the designated constraints aren't met.
type GetExceptionsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExceptionsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExceptionsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExceptionsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExceptionsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExceptionsRequestValidationError) ErrorName() string {
	return "GetExceptionsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetExceptionsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExceptionsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExceptionsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExceptionsRequestValidationError{}

// Validate checks the field values on GetExceptionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExceptionsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExceptionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExceptionsResponseMultiError, or nil if none found.
func (m *GetExceptionsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExceptionsResponse) validate(all bool) error {
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
					errors = append(errors, GetExceptionsResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetExceptionsResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetExceptionsResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetExceptionsResponseMultiError(errors)
	}

	return nil
}

// GetExceptionsResponseMultiError is an error wrapping multiple validation
// errors returned by GetExceptionsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetExceptionsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExceptionsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExceptionsResponseMultiError) AllErrors() []error { return m }

// GetExceptionsResponseValidationError is the validation error returned by
// GetExceptionsResponse.Validate if the designated constraints aren't met.
type GetExceptionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExceptionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExceptionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExceptionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExceptionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExceptionsResponseValidationError) ErrorName() string {
	return "GetExceptionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetExceptionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExceptionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExceptionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExceptionsResponseValidationError{}

// Validate checks the field values on GetExceptionEventIdsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExceptionEventIdsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExceptionEventIdsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExceptionEventIdsRequestMultiError, or nil if none found.
func (m *GetExceptionEventIdsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExceptionEventIdsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetExceptionID()) < 1 {
		err := GetExceptionEventIdsRequestValidationError{
			field:  "ExceptionID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetScopeID()) < 1 {
		err := GetExceptionEventIdsRequestValidationError{
			field:  "ScopeID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetExceptionEventIdsRequestMultiError(errors)
	}

	return nil
}

// GetExceptionEventIdsRequestMultiError is an error wrapping multiple
// validation errors returned by GetExceptionEventIdsRequest.ValidateAll() if
// the designated constraints aren't met.
type GetExceptionEventIdsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExceptionEventIdsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExceptionEventIdsRequestMultiError) AllErrors() []error { return m }

// GetExceptionEventIdsRequestValidationError is the validation error returned
// by GetExceptionEventIdsRequest.Validate if the designated constraints
// aren't met.
type GetExceptionEventIdsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExceptionEventIdsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExceptionEventIdsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExceptionEventIdsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExceptionEventIdsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExceptionEventIdsRequestValidationError) ErrorName() string {
	return "GetExceptionEventIdsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetExceptionEventIdsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExceptionEventIdsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExceptionEventIdsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExceptionEventIdsRequestValidationError{}

// Validate checks the field values on GetExceptionEventIdsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExceptionEventIdsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExceptionEventIdsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExceptionEventIdsResponseMultiError, or nil if none found.
func (m *GetExceptionEventIdsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExceptionEventIdsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetExceptionEventIdsResponseMultiError(errors)
	}

	return nil
}

// GetExceptionEventIdsResponseMultiError is an error wrapping multiple
// validation errors returned by GetExceptionEventIdsResponse.ValidateAll() if
// the designated constraints aren't met.
type GetExceptionEventIdsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExceptionEventIdsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExceptionEventIdsResponseMultiError) AllErrors() []error { return m }

// GetExceptionEventIdsResponseValidationError is the validation error returned
// by GetExceptionEventIdsResponse.Validate if the designated constraints
// aren't met.
type GetExceptionEventIdsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExceptionEventIdsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExceptionEventIdsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExceptionEventIdsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExceptionEventIdsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExceptionEventIdsResponseValidationError) ErrorName() string {
	return "GetExceptionEventIdsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetExceptionEventIdsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExceptionEventIdsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExceptionEventIdsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExceptionEventIdsResponseValidationError{}

// Validate checks the field values on GetExceptionEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExceptionEventRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExceptionEventRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExceptionEventRequestMultiError, or nil if none found.
func (m *GetExceptionEventRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExceptionEventRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetExceptionEventID()) < 1 {
		err := GetExceptionEventRequestValidationError{
			field:  "ExceptionEventID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetScopeID()) < 1 {
		err := GetExceptionEventRequestValidationError{
			field:  "ScopeID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetExceptionEventRequestMultiError(errors)
	}

	return nil
}

// GetExceptionEventRequestMultiError is an error wrapping multiple validation
// errors returned by GetExceptionEventRequest.ValidateAll() if the designated
// constraints aren't met.
type GetExceptionEventRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExceptionEventRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExceptionEventRequestMultiError) AllErrors() []error { return m }

// GetExceptionEventRequestValidationError is the validation error returned by
// GetExceptionEventRequest.Validate if the designated constraints aren't met.
type GetExceptionEventRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExceptionEventRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExceptionEventRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExceptionEventRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExceptionEventRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExceptionEventRequestValidationError) ErrorName() string {
	return "GetExceptionEventRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetExceptionEventRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExceptionEventRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExceptionEventRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExceptionEventRequestValidationError{}

// Validate checks the field values on GetExceptionEventResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExceptionEventResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExceptionEventResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExceptionEventResponseMultiError, or nil if none found.
func (m *GetExceptionEventResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExceptionEventResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetExceptionEventResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetExceptionEventResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetExceptionEventResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetExceptionEventResponseMultiError(errors)
	}

	return nil
}

// GetExceptionEventResponseMultiError is an error wrapping multiple validation
// errors returned by GetExceptionEventResponse.ValidateAll() if the
// designated constraints aren't met.
type GetExceptionEventResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExceptionEventResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExceptionEventResponseMultiError) AllErrors() []error { return m }

// GetExceptionEventResponseValidationError is the validation error returned by
// GetExceptionEventResponse.Validate if the designated constraints aren't met.
type GetExceptionEventResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExceptionEventResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExceptionEventResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExceptionEventResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExceptionEventResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExceptionEventResponseValidationError) ErrorName() string {
	return "GetExceptionEventResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetExceptionEventResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExceptionEventResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExceptionEventResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExceptionEventResponseValidationError{}

// Validate checks the field values on Exception with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Exception) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Exception with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExceptionMultiError, or nil
// if none found.
func (m *Exception) ValidateAll() error {
	return m.validate(true)
}

func (m *Exception) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ClassName

	// no validation rules for Method

	// no validation rules for Type

	// no validation rules for EventCount

	// no validation rules for ExceptionMessage

	// no validation rules for File

	// no validation rules for ApplicationID

	// no validation rules for RuntimeID

	// no validation rules for ServiceName

	// no validation rules for ScopeID

	// no validation rules for CreateTime

	// no validation rules for UpdateTime

	if len(errors) > 0 {
		return ExceptionMultiError(errors)
	}

	return nil
}

// ExceptionMultiError is an error wrapping multiple validation errors returned
// by Exception.ValidateAll() if the designated constraints aren't met.
type ExceptionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExceptionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExceptionMultiError) AllErrors() []error { return m }

// ExceptionValidationError is the validation error returned by
// Exception.Validate if the designated constraints aren't met.
type ExceptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExceptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExceptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExceptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExceptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExceptionValidationError) ErrorName() string { return "ExceptionValidationError" }

// Error satisfies the builtin error interface
func (e ExceptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sException.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExceptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExceptionValidationError{}

// Validate checks the field values on Stacks with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Stacks) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Stacks with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StacksMultiError, or nil if none found.
func (m *Stacks) ValidateAll() error {
	return m.validate(true)
}

func (m *Stacks) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetStack()))
		i := 0
		for key := range m.GetStack() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetStack()[key]
			_ = val

			// no validation rules for Stack[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, StacksValidationError{
							field:  fmt.Sprintf("Stack[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, StacksValidationError{
							field:  fmt.Sprintf("Stack[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return StacksValidationError{
						field:  fmt.Sprintf("Stack[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return StacksMultiError(errors)
	}

	return nil
}

// StacksMultiError is an error wrapping multiple validation errors returned by
// Stacks.ValidateAll() if the designated constraints aren't met.
type StacksMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StacksMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StacksMultiError) AllErrors() []error { return m }

// StacksValidationError is the validation error returned by Stacks.Validate if
// the designated constraints aren't met.
type StacksValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StacksValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StacksValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StacksValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StacksValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StacksValidationError) ErrorName() string { return "StacksValidationError" }

// Error satisfies the builtin error interface
func (e StacksValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStacks.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StacksValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StacksValidationError{}

// Validate checks the field values on ExceptionEvent with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExceptionEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExceptionEvent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExceptionEventMultiError,
// or nil if none found.
func (m *ExceptionEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *ExceptionEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ExceptionID

	// no validation rules for Metadata

	// no validation rules for RequestContext

	// no validation rules for RequestHeaders

	// no validation rules for RequestID

	for idx, item := range m.GetStacks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExceptionEventValidationError{
						field:  fmt.Sprintf("Stacks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExceptionEventValidationError{
						field:  fmt.Sprintf("Stacks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExceptionEventValidationError{
					field:  fmt.Sprintf("Stacks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Tags

	// no validation rules for Timestamp

	// no validation rules for RequestSampled

	if len(errors) > 0 {
		return ExceptionEventMultiError(errors)
	}

	return nil
}

// ExceptionEventMultiError is an error wrapping multiple validation errors
// returned by ExceptionEvent.ValidateAll() if the designated constraints
// aren't met.
type ExceptionEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExceptionEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExceptionEventMultiError) AllErrors() []error { return m }

// ExceptionEventValidationError is the validation error returned by
// ExceptionEvent.Validate if the designated constraints aren't met.
type ExceptionEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExceptionEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExceptionEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExceptionEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExceptionEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExceptionEventValidationError) ErrorName() string { return "ExceptionEventValidationError" }

// Error satisfies the builtin error interface
func (e ExceptionEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExceptionEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExceptionEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExceptionEventValidationError{}
