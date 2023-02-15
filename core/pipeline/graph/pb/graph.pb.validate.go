// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: graph.proto

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

// Validate checks the field values on PipelineYmlGraphRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PipelineYmlGraphRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PipelineYmlGraphRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PipelineYmlGraphRequestMultiError, or nil if none found.
func (m *PipelineYmlGraphRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PipelineYmlGraphRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PipelineYmlContent

	// no validation rules for GlobalSnippetConfigLabels

	if all {
		switch v := interface{}(m.GetSnippetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PipelineYmlGraphRequestValidationError{
					field:  "SnippetConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PipelineYmlGraphRequestValidationError{
					field:  "SnippetConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSnippetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PipelineYmlGraphRequestValidationError{
				field:  "SnippetConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PipelineYmlGraphRequestMultiError(errors)
	}

	return nil
}

// PipelineYmlGraphRequestMultiError is an error wrapping multiple validation
// errors returned by PipelineYmlGraphRequest.ValidateAll() if the designated
// constraints aren't met.
type PipelineYmlGraphRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PipelineYmlGraphRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PipelineYmlGraphRequestMultiError) AllErrors() []error { return m }

// PipelineYmlGraphRequestValidationError is the validation error returned by
// PipelineYmlGraphRequest.Validate if the designated constraints aren't met.
type PipelineYmlGraphRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PipelineYmlGraphRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PipelineYmlGraphRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PipelineYmlGraphRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PipelineYmlGraphRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PipelineYmlGraphRequestValidationError) ErrorName() string {
	return "PipelineYmlGraphRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PipelineYmlGraphRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPipelineYmlGraphRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PipelineYmlGraphRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PipelineYmlGraphRequestValidationError{}

// Validate checks the field values on PipelineYmlGraphResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PipelineYmlGraphResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PipelineYmlGraphResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PipelineYmlGraphResponseMultiError, or nil if none found.
func (m *PipelineYmlGraphResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PipelineYmlGraphResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PipelineYmlGraphResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PipelineYmlGraphResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PipelineYmlGraphResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PipelineYmlGraphResponseMultiError(errors)
	}

	return nil
}

// PipelineYmlGraphResponseMultiError is an error wrapping multiple validation
// errors returned by PipelineYmlGraphResponse.ValidateAll() if the designated
// constraints aren't met.
type PipelineYmlGraphResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PipelineYmlGraphResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PipelineYmlGraphResponseMultiError) AllErrors() []error { return m }

// PipelineYmlGraphResponseValidationError is the validation error returned by
// PipelineYmlGraphResponse.Validate if the designated constraints aren't met.
type PipelineYmlGraphResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PipelineYmlGraphResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PipelineYmlGraphResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PipelineYmlGraphResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PipelineYmlGraphResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PipelineYmlGraphResponseValidationError) ErrorName() string {
	return "PipelineYmlGraphResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PipelineYmlGraphResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPipelineYmlGraphResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PipelineYmlGraphResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PipelineYmlGraphResponseValidationError{}
