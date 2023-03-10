// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api_policy.proto

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

// Validate checks the field values on SetPolicyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SetPolicyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetPolicyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetPolicyResponseMultiError, or nil if none found.
func (m *SetPolicyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SetPolicyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SetPolicyResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SetPolicyResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetPolicyResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SetPolicyResponseMultiError(errors)
	}

	return nil
}

// SetPolicyResponseMultiError is an error wrapping multiple validation errors
// returned by SetPolicyResponse.ValidateAll() if the designated constraints
// aren't met.
type SetPolicyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetPolicyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetPolicyResponseMultiError) AllErrors() []error { return m }

// SetPolicyResponseValidationError is the validation error returned by
// SetPolicyResponse.Validate if the designated constraints aren't met.
type SetPolicyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetPolicyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetPolicyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetPolicyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetPolicyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetPolicyResponseValidationError) ErrorName() string {
	return "SetPolicyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SetPolicyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetPolicyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetPolicyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetPolicyResponseValidationError{}

// Validate checks the field values on SetPolicyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SetPolicyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetPolicyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetPolicyRequestMultiError, or nil if none found.
func (m *SetPolicyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SetPolicyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Category

	// no validation rules for PackageId

	// no validation rules for ApiId

	// no validation rules for Body

	if len(errors) > 0 {
		return SetPolicyRequestMultiError(errors)
	}

	return nil
}

// SetPolicyRequestMultiError is an error wrapping multiple validation errors
// returned by SetPolicyRequest.ValidateAll() if the designated constraints
// aren't met.
type SetPolicyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetPolicyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetPolicyRequestMultiError) AllErrors() []error { return m }

// SetPolicyRequestValidationError is the validation error returned by
// SetPolicyRequest.Validate if the designated constraints aren't met.
type SetPolicyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetPolicyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetPolicyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetPolicyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetPolicyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetPolicyRequestValidationError) ErrorName() string { return "SetPolicyRequestValidationError" }

// Error satisfies the builtin error interface
func (e SetPolicyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetPolicyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetPolicyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetPolicyRequestValidationError{}

// Validate checks the field values on GetPolicyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetPolicyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPolicyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPolicyRequestMultiError, or nil if none found.
func (m *GetPolicyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPolicyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Category

	// no validation rules for PackageId

	// no validation rules for ApiId

	if len(errors) > 0 {
		return GetPolicyRequestMultiError(errors)
	}

	return nil
}

// GetPolicyRequestMultiError is an error wrapping multiple validation errors
// returned by GetPolicyRequest.ValidateAll() if the designated constraints
// aren't met.
type GetPolicyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPolicyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPolicyRequestMultiError) AllErrors() []error { return m }

// GetPolicyRequestValidationError is the validation error returned by
// GetPolicyRequest.Validate if the designated constraints aren't met.
type GetPolicyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPolicyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPolicyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPolicyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPolicyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPolicyRequestValidationError) ErrorName() string { return "GetPolicyRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetPolicyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPolicyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPolicyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPolicyRequestValidationError{}

// Validate checks the field values on GetPolicyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetPolicyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPolicyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPolicyResponseMultiError, or nil if none found.
func (m *GetPolicyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPolicyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetPolicyResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetPolicyResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPolicyResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetPolicyResponseMultiError(errors)
	}

	return nil
}

// GetPolicyResponseMultiError is an error wrapping multiple validation errors
// returned by GetPolicyResponse.ValidateAll() if the designated constraints
// aren't met.
type GetPolicyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPolicyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPolicyResponseMultiError) AllErrors() []error { return m }

// GetPolicyResponseValidationError is the validation error returned by
// GetPolicyResponse.Validate if the designated constraints aren't met.
type GetPolicyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPolicyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPolicyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPolicyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPolicyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPolicyResponseValidationError) ErrorName() string {
	return "GetPolicyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPolicyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPolicyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPolicyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPolicyResponseValidationError{}
