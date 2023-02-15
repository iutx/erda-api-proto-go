// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: trigger.proto

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

// Validate checks the field values on PipelineTriggerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PipelineTriggerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PipelineTriggerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PipelineTriggerRequestMultiError, or nil if none found.
func (m *PipelineTriggerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PipelineTriggerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EventName

	// no validation rules for Label

	if len(errors) > 0 {
		return PipelineTriggerRequestMultiError(errors)
	}

	return nil
}

// PipelineTriggerRequestMultiError is an error wrapping multiple validation
// errors returned by PipelineTriggerRequest.ValidateAll() if the designated
// constraints aren't met.
type PipelineTriggerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PipelineTriggerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PipelineTriggerRequestMultiError) AllErrors() []error { return m }

// PipelineTriggerRequestValidationError is the validation error returned by
// PipelineTriggerRequest.Validate if the designated constraints aren't met.
type PipelineTriggerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PipelineTriggerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PipelineTriggerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PipelineTriggerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PipelineTriggerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PipelineTriggerRequestValidationError) ErrorName() string {
	return "PipelineTriggerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PipelineTriggerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPipelineTriggerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PipelineTriggerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PipelineTriggerRequestValidationError{}

// Validate checks the field values on PipelineTriggerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PipelineTriggerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PipelineTriggerResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PipelineTriggerResponseMultiError, or nil if none found.
func (m *PipelineTriggerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PipelineTriggerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PipelineTriggerResponseMultiError(errors)
	}

	return nil
}

// PipelineTriggerResponseMultiError is an error wrapping multiple validation
// errors returned by PipelineTriggerResponse.ValidateAll() if the designated
// constraints aren't met.
type PipelineTriggerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PipelineTriggerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PipelineTriggerResponseMultiError) AllErrors() []error { return m }

// PipelineTriggerResponseValidationError is the validation error returned by
// PipelineTriggerResponse.Validate if the designated constraints aren't met.
type PipelineTriggerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PipelineTriggerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PipelineTriggerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PipelineTriggerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PipelineTriggerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PipelineTriggerResponseValidationError) ErrorName() string {
	return "PipelineTriggerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PipelineTriggerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPipelineTriggerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PipelineTriggerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PipelineTriggerResponseValidationError{}