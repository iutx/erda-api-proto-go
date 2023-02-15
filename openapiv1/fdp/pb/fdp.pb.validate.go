// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: fdp.proto

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

// Validate checks the field values on CDP_DELETE_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CDP_DELETE_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CDP_DELETE_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CDP_DELETE_RequestMultiError, or nil if none found.
func (m *CDP_DELETE_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CDP_DELETE_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CDP_DELETE_RequestMultiError(errors)
	}

	return nil
}

// CDP_DELETE_RequestMultiError is an error wrapping multiple validation errors
// returned by CDP_DELETE_Request.ValidateAll() if the designated constraints
// aren't met.
type CDP_DELETE_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CDP_DELETE_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CDP_DELETE_RequestMultiError) AllErrors() []error { return m }

// CDP_DELETE_RequestValidationError is the validation error returned by
// CDP_DELETE_Request.Validate if the designated constraints aren't met.
type CDP_DELETE_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CDP_DELETE_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CDP_DELETE_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CDP_DELETE_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CDP_DELETE_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CDP_DELETE_RequestValidationError) ErrorName() string {
	return "CDP_DELETE_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CDP_DELETE_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCDP_DELETE_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CDP_DELETE_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CDP_DELETE_RequestValidationError{}

// Validate checks the field values on CDP_GET_Request with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CDP_GET_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CDP_GET_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CDP_GET_RequestMultiError, or nil if none found.
func (m *CDP_GET_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CDP_GET_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CDP_GET_RequestMultiError(errors)
	}

	return nil
}

// CDP_GET_RequestMultiError is an error wrapping multiple validation errors
// returned by CDP_GET_Request.ValidateAll() if the designated constraints
// aren't met.
type CDP_GET_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CDP_GET_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CDP_GET_RequestMultiError) AllErrors() []error { return m }

// CDP_GET_RequestValidationError is the validation error returned by
// CDP_GET_Request.Validate if the designated constraints aren't met.
type CDP_GET_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CDP_GET_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CDP_GET_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CDP_GET_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CDP_GET_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CDP_GET_RequestValidationError) ErrorName() string { return "CDP_GET_RequestValidationError" }

// Error satisfies the builtin error interface
func (e CDP_GET_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCDP_GET_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CDP_GET_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CDP_GET_RequestValidationError{}

// Validate checks the field values on CDP_PATCH_Request with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CDP_PATCH_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CDP_PATCH_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CDP_PATCH_RequestMultiError, or nil if none found.
func (m *CDP_PATCH_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CDP_PATCH_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CDP_PATCH_RequestMultiError(errors)
	}

	return nil
}

// CDP_PATCH_RequestMultiError is an error wrapping multiple validation errors
// returned by CDP_PATCH_Request.ValidateAll() if the designated constraints
// aren't met.
type CDP_PATCH_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CDP_PATCH_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CDP_PATCH_RequestMultiError) AllErrors() []error { return m }

// CDP_PATCH_RequestValidationError is the validation error returned by
// CDP_PATCH_Request.Validate if the designated constraints aren't met.
type CDP_PATCH_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CDP_PATCH_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CDP_PATCH_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CDP_PATCH_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CDP_PATCH_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CDP_PATCH_RequestValidationError) ErrorName() string {
	return "CDP_PATCH_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CDP_PATCH_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCDP_PATCH_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CDP_PATCH_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CDP_PATCH_RequestValidationError{}

// Validate checks the field values on CDP_POST_Request with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CDP_POST_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CDP_POST_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CDP_POST_RequestMultiError, or nil if none found.
func (m *CDP_POST_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CDP_POST_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CDP_POST_RequestMultiError(errors)
	}

	return nil
}

// CDP_POST_RequestMultiError is an error wrapping multiple validation errors
// returned by CDP_POST_Request.ValidateAll() if the designated constraints
// aren't met.
type CDP_POST_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CDP_POST_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CDP_POST_RequestMultiError) AllErrors() []error { return m }

// CDP_POST_RequestValidationError is the validation error returned by
// CDP_POST_Request.Validate if the designated constraints aren't met.
type CDP_POST_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CDP_POST_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CDP_POST_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CDP_POST_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CDP_POST_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CDP_POST_RequestValidationError) ErrorName() string { return "CDP_POST_RequestValidationError" }

// Error satisfies the builtin error interface
func (e CDP_POST_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCDP_POST_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CDP_POST_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CDP_POST_RequestValidationError{}

// Validate checks the field values on CDP_PUT_Request with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CDP_PUT_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CDP_PUT_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CDP_PUT_RequestMultiError, or nil if none found.
func (m *CDP_PUT_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CDP_PUT_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CDP_PUT_RequestMultiError(errors)
	}

	return nil
}

// CDP_PUT_RequestMultiError is an error wrapping multiple validation errors
// returned by CDP_PUT_Request.ValidateAll() if the designated constraints
// aren't met.
type CDP_PUT_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CDP_PUT_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CDP_PUT_RequestMultiError) AllErrors() []error { return m }

// CDP_PUT_RequestValidationError is the validation error returned by
// CDP_PUT_Request.Validate if the designated constraints aren't met.
type CDP_PUT_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CDP_PUT_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CDP_PUT_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CDP_PUT_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CDP_PUT_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CDP_PUT_RequestValidationError) ErrorName() string { return "CDP_PUT_RequestValidationError" }

// Error satisfies the builtin error interface
func (e CDP_PUT_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCDP_PUT_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CDP_PUT_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CDP_PUT_RequestValidationError{}

// Validate checks the field values on FDP_WEBSOCKET_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FDP_WEBSOCKET_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FDP_WEBSOCKET_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FDP_WEBSOCKET_RequestMultiError, or nil if none found.
func (m *FDP_WEBSOCKET_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *FDP_WEBSOCKET_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FDP_WEBSOCKET_RequestMultiError(errors)
	}

	return nil
}

// FDP_WEBSOCKET_RequestMultiError is an error wrapping multiple validation
// errors returned by FDP_WEBSOCKET_Request.ValidateAll() if the designated
// constraints aren't met.
type FDP_WEBSOCKET_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FDP_WEBSOCKET_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FDP_WEBSOCKET_RequestMultiError) AllErrors() []error { return m }

// FDP_WEBSOCKET_RequestValidationError is the validation error returned by
// FDP_WEBSOCKET_Request.Validate if the designated constraints aren't met.
type FDP_WEBSOCKET_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FDP_WEBSOCKET_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FDP_WEBSOCKET_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FDP_WEBSOCKET_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FDP_WEBSOCKET_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FDP_WEBSOCKET_RequestValidationError) ErrorName() string {
	return "FDP_WEBSOCKET_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e FDP_WEBSOCKET_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFDP_WEBSOCKET_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FDP_WEBSOCKET_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FDP_WEBSOCKET_RequestValidationError{}