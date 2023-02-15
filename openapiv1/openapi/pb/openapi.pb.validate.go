// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: openapi.proto

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

// Validate checks the field values on DICE_METADATA_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DICE_METADATA_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DICE_METADATA_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DICE_METADATA_RequestMultiError, or nil if none found.
func (m *DICE_METADATA_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *DICE_METADATA_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DICE_METADATA_RequestMultiError(errors)
	}

	return nil
}

// DICE_METADATA_RequestMultiError is an error wrapping multiple validation
// errors returned by DICE_METADATA_Request.ValidateAll() if the designated
// constraints aren't met.
type DICE_METADATA_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DICE_METADATA_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DICE_METADATA_RequestMultiError) AllErrors() []error { return m }

// DICE_METADATA_RequestValidationError is the validation error returned by
// DICE_METADATA_Request.Validate if the designated constraints aren't met.
type DICE_METADATA_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DICE_METADATA_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DICE_METADATA_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DICE_METADATA_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DICE_METADATA_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DICE_METADATA_RequestValidationError) ErrorName() string {
	return "DICE_METADATA_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DICE_METADATA_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDICE_METADATA_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DICE_METADATA_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DICE_METADATA_RequestValidationError{}

// Validate checks the field values on DOC_JSON_Request with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DOC_JSON_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DOC_JSON_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DOC_JSON_RequestMultiError, or nil if none found.
func (m *DOC_JSON_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *DOC_JSON_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DOC_JSON_RequestMultiError(errors)
	}

	return nil
}

// DOC_JSON_RequestMultiError is an error wrapping multiple validation errors
// returned by DOC_JSON_Request.ValidateAll() if the designated constraints
// aren't met.
type DOC_JSON_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DOC_JSON_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DOC_JSON_RequestMultiError) AllErrors() []error { return m }

// DOC_JSON_RequestValidationError is the validation error returned by
// DOC_JSON_Request.Validate if the designated constraints aren't met.
type DOC_JSON_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DOC_JSON_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DOC_JSON_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DOC_JSON_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DOC_JSON_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DOC_JSON_RequestValidationError) ErrorName() string { return "DOC_JSON_RequestValidationError" }

// Error satisfies the builtin error interface
func (e DOC_JSON_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDOC_JSON_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DOC_JSON_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DOC_JSON_RequestValidationError{}

// Validate checks the field values on OPENAPI_DOC_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OPENAPI_DOC_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OPENAPI_DOC_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OPENAPI_DOC_RequestMultiError, or nil if none found.
func (m *OPENAPI_DOC_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *OPENAPI_DOC_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OPENAPI_DOC_RequestMultiError(errors)
	}

	return nil
}

// OPENAPI_DOC_RequestMultiError is an error wrapping multiple validation
// errors returned by OPENAPI_DOC_Request.ValidateAll() if the designated
// constraints aren't met.
type OPENAPI_DOC_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OPENAPI_DOC_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OPENAPI_DOC_RequestMultiError) AllErrors() []error { return m }

// OPENAPI_DOC_RequestValidationError is the validation error returned by
// OPENAPI_DOC_Request.Validate if the designated constraints aren't met.
type OPENAPI_DOC_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OPENAPI_DOC_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OPENAPI_DOC_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OPENAPI_DOC_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OPENAPI_DOC_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OPENAPI_DOC_RequestValidationError) ErrorName() string {
	return "OPENAPI_DOC_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e OPENAPI_DOC_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOPENAPI_DOC_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OPENAPI_DOC_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OPENAPI_DOC_RequestValidationError{}

// Validate checks the field values on OPENAPI_EVENT_DOC_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OPENAPI_EVENT_DOC_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OPENAPI_EVENT_DOC_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OPENAPI_EVENT_DOC_RequestMultiError, or nil if none found.
func (m *OPENAPI_EVENT_DOC_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *OPENAPI_EVENT_DOC_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OPENAPI_EVENT_DOC_RequestMultiError(errors)
	}

	return nil
}

// OPENAPI_EVENT_DOC_RequestMultiError is an error wrapping multiple validation
// errors returned by OPENAPI_EVENT_DOC_Request.ValidateAll() if the
// designated constraints aren't met.
type OPENAPI_EVENT_DOC_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OPENAPI_EVENT_DOC_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OPENAPI_EVENT_DOC_RequestMultiError) AllErrors() []error { return m }

// OPENAPI_EVENT_DOC_RequestValidationError is the validation error returned by
// OPENAPI_EVENT_DOC_Request.Validate if the designated constraints aren't met.
type OPENAPI_EVENT_DOC_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OPENAPI_EVENT_DOC_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OPENAPI_EVENT_DOC_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OPENAPI_EVENT_DOC_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OPENAPI_EVENT_DOC_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OPENAPI_EVENT_DOC_RequestValidationError) ErrorName() string {
	return "OPENAPI_EVENT_DOC_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e OPENAPI_EVENT_DOC_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOPENAPI_EVENT_DOC_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OPENAPI_EVENT_DOC_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OPENAPI_EVENT_DOC_RequestValidationError{}

// Validate checks the field values on OPENAPI_GEN_CLIENT_TOKEN_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *OPENAPI_GEN_CLIENT_TOKEN_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OPENAPI_GEN_CLIENT_TOKEN_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// OPENAPI_GEN_CLIENT_TOKEN_RequestMultiError, or nil if none found.
func (m *OPENAPI_GEN_CLIENT_TOKEN_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *OPENAPI_GEN_CLIENT_TOKEN_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OPENAPI_GEN_CLIENT_TOKEN_RequestMultiError(errors)
	}

	return nil
}

// OPENAPI_GEN_CLIENT_TOKEN_RequestMultiError is an error wrapping multiple
// validation errors returned by
// OPENAPI_GEN_CLIENT_TOKEN_Request.ValidateAll() if the designated
// constraints aren't met.
type OPENAPI_GEN_CLIENT_TOKEN_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OPENAPI_GEN_CLIENT_TOKEN_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OPENAPI_GEN_CLIENT_TOKEN_RequestMultiError) AllErrors() []error { return m }

// OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError is the validation error
// returned by OPENAPI_GEN_CLIENT_TOKEN_Request.Validate if the designated
// constraints aren't met.
type OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError) ErrorName() string {
	return "OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOPENAPI_GEN_CLIENT_TOKEN_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OPENAPI_GEN_CLIENT_TOKEN_RequestValidationError{}

// Validate checks the field values on OPENAPI_LIST_CLIENT_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OPENAPI_LIST_CLIENT_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OPENAPI_LIST_CLIENT_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OPENAPI_LIST_CLIENT_RequestMultiError, or nil if none found.
func (m *OPENAPI_LIST_CLIENT_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *OPENAPI_LIST_CLIENT_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OPENAPI_LIST_CLIENT_RequestMultiError(errors)
	}

	return nil
}

// OPENAPI_LIST_CLIENT_RequestMultiError is an error wrapping multiple
// validation errors returned by OPENAPI_LIST_CLIENT_Request.ValidateAll() if
// the designated constraints aren't met.
type OPENAPI_LIST_CLIENT_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OPENAPI_LIST_CLIENT_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OPENAPI_LIST_CLIENT_RequestMultiError) AllErrors() []error { return m }

// OPENAPI_LIST_CLIENT_RequestValidationError is the validation error returned
// by OPENAPI_LIST_CLIENT_Request.Validate if the designated constraints
// aren't met.
type OPENAPI_LIST_CLIENT_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OPENAPI_LIST_CLIENT_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OPENAPI_LIST_CLIENT_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OPENAPI_LIST_CLIENT_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OPENAPI_LIST_CLIENT_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OPENAPI_LIST_CLIENT_RequestValidationError) ErrorName() string {
	return "OPENAPI_LIST_CLIENT_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e OPENAPI_LIST_CLIENT_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOPENAPI_LIST_CLIENT_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OPENAPI_LIST_CLIENT_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OPENAPI_LIST_CLIENT_RequestValidationError{}

// Validate checks the field values on OPENAPI_METRICS_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OPENAPI_METRICS_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OPENAPI_METRICS_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OPENAPI_METRICS_RequestMultiError, or nil if none found.
func (m *OPENAPI_METRICS_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *OPENAPI_METRICS_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OPENAPI_METRICS_RequestMultiError(errors)
	}

	return nil
}

// OPENAPI_METRICS_RequestMultiError is an error wrapping multiple validation
// errors returned by OPENAPI_METRICS_Request.ValidateAll() if the designated
// constraints aren't met.
type OPENAPI_METRICS_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OPENAPI_METRICS_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OPENAPI_METRICS_RequestMultiError) AllErrors() []error { return m }

// OPENAPI_METRICS_RequestValidationError is the validation error returned by
// OPENAPI_METRICS_Request.Validate if the designated constraints aren't met.
type OPENAPI_METRICS_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OPENAPI_METRICS_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OPENAPI_METRICS_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OPENAPI_METRICS_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OPENAPI_METRICS_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OPENAPI_METRICS_RequestValidationError) ErrorName() string {
	return "OPENAPI_METRICS_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e OPENAPI_METRICS_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOPENAPI_METRICS_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OPENAPI_METRICS_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OPENAPI_METRICS_RequestValidationError{}

// Validate checks the field values on OPENAPI_NEW_CLIENT_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OPENAPI_NEW_CLIENT_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OPENAPI_NEW_CLIENT_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OPENAPI_NEW_CLIENT_RequestMultiError, or nil if none found.
func (m *OPENAPI_NEW_CLIENT_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *OPENAPI_NEW_CLIENT_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OPENAPI_NEW_CLIENT_RequestMultiError(errors)
	}

	return nil
}

// OPENAPI_NEW_CLIENT_RequestMultiError is an error wrapping multiple
// validation errors returned by OPENAPI_NEW_CLIENT_Request.ValidateAll() if
// the designated constraints aren't met.
type OPENAPI_NEW_CLIENT_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OPENAPI_NEW_CLIENT_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OPENAPI_NEW_CLIENT_RequestMultiError) AllErrors() []error { return m }

// OPENAPI_NEW_CLIENT_RequestValidationError is the validation error returned
// by OPENAPI_NEW_CLIENT_Request.Validate if the designated constraints aren't met.
type OPENAPI_NEW_CLIENT_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OPENAPI_NEW_CLIENT_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OPENAPI_NEW_CLIENT_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OPENAPI_NEW_CLIENT_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OPENAPI_NEW_CLIENT_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OPENAPI_NEW_CLIENT_RequestValidationError) ErrorName() string {
	return "OPENAPI_NEW_CLIENT_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e OPENAPI_NEW_CLIENT_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOPENAPI_NEW_CLIENT_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OPENAPI_NEW_CLIENT_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OPENAPI_NEW_CLIENT_RequestValidationError{}

// Validate checks the field values on OPENAPI_STAT_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OPENAPI_STAT_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OPENAPI_STAT_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OPENAPI_STAT_RequestMultiError, or nil if none found.
func (m *OPENAPI_STAT_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *OPENAPI_STAT_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OPENAPI_STAT_RequestMultiError(errors)
	}

	return nil
}

// OPENAPI_STAT_RequestMultiError is an error wrapping multiple validation
// errors returned by OPENAPI_STAT_Request.ValidateAll() if the designated
// constraints aren't met.
type OPENAPI_STAT_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OPENAPI_STAT_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OPENAPI_STAT_RequestMultiError) AllErrors() []error { return m }

// OPENAPI_STAT_RequestValidationError is the validation error returned by
// OPENAPI_STAT_Request.Validate if the designated constraints aren't met.
type OPENAPI_STAT_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OPENAPI_STAT_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OPENAPI_STAT_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OPENAPI_STAT_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OPENAPI_STAT_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OPENAPI_STAT_RequestValidationError) ErrorName() string {
	return "OPENAPI_STAT_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e OPENAPI_STAT_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOPENAPI_STAT_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OPENAPI_STAT_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OPENAPI_STAT_RequestValidationError{}

// Validate checks the field values on OPENAPI_VERSION_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OPENAPI_VERSION_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OPENAPI_VERSION_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OPENAPI_VERSION_RequestMultiError, or nil if none found.
func (m *OPENAPI_VERSION_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *OPENAPI_VERSION_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OPENAPI_VERSION_RequestMultiError(errors)
	}

	return nil
}

// OPENAPI_VERSION_RequestMultiError is an error wrapping multiple validation
// errors returned by OPENAPI_VERSION_Request.ValidateAll() if the designated
// constraints aren't met.
type OPENAPI_VERSION_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OPENAPI_VERSION_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OPENAPI_VERSION_RequestMultiError) AllErrors() []error { return m }

// OPENAPI_VERSION_RequestValidationError is the validation error returned by
// OPENAPI_VERSION_Request.Validate if the designated constraints aren't met.
type OPENAPI_VERSION_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OPENAPI_VERSION_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OPENAPI_VERSION_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OPENAPI_VERSION_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OPENAPI_VERSION_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OPENAPI_VERSION_RequestValidationError) ErrorName() string {
	return "OPENAPI_VERSION_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e OPENAPI_VERSION_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOPENAPI_VERSION_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OPENAPI_VERSION_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OPENAPI_VERSION_RequestValidationError{}
