// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: issuerelation.proto

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

// Validate checks the field values on IssueRelation with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IssueRelation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IssueRelation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IssueRelationMultiError, or
// nil if none found.
func (m *IssueRelation) ValidateAll() error {
	return m.validate(true)
}

func (m *IssueRelation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ID

	if all {
		switch v := interface{}(m.GetTimeCreated()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IssueRelationValidationError{
					field:  "TimeCreated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IssueRelationValidationError{
					field:  "TimeCreated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssueRelationValidationError{
				field:  "TimeCreated",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimeUpdated()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IssueRelationValidationError{
					field:  "TimeUpdated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IssueRelationValidationError{
					field:  "TimeUpdated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeUpdated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssueRelationValidationError{
				field:  "TimeUpdated",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSoftDeletedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IssueRelationValidationError{
					field:  "SoftDeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IssueRelationValidationError{
					field:  "SoftDeletedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSoftDeletedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssueRelationValidationError{
				field:  "SoftDeletedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Relation

	// no validation rules for IssueID

	// no validation rules for Type

	// no validation rules for Extra

	if len(errors) > 0 {
		return IssueRelationMultiError(errors)
	}

	return nil
}

// IssueRelationMultiError is an error wrapping multiple validation errors
// returned by IssueRelation.ValidateAll() if the designated constraints
// aren't met.
type IssueRelationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IssueRelationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IssueRelationMultiError) AllErrors() []error { return m }

// IssueRelationValidationError is the validation error returned by
// IssueRelation.Validate if the designated constraints aren't met.
type IssueRelationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IssueRelationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IssueRelationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IssueRelationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IssueRelationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IssueRelationValidationError) ErrorName() string { return "IssueRelationValidationError" }

// Error satisfies the builtin error interface
func (e IssueRelationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIssueRelation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IssueRelationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IssueRelationValidationError{}

// Validate checks the field values on CreateIssueRelationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateIssueRelationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateIssueRelationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateIssueRelationRequestMultiError, or nil if none found.
func (m *CreateIssueRelationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateIssueRelationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Relation

	// no validation rules for IssueID

	// no validation rules for Type

	// no validation rules for Extra

	if len(errors) > 0 {
		return CreateIssueRelationRequestMultiError(errors)
	}

	return nil
}

// CreateIssueRelationRequestMultiError is an error wrapping multiple
// validation errors returned by CreateIssueRelationRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateIssueRelationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateIssueRelationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateIssueRelationRequestMultiError) AllErrors() []error { return m }

// CreateIssueRelationRequestValidationError is the validation error returned
// by CreateIssueRelationRequest.Validate if the designated constraints aren't met.
type CreateIssueRelationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateIssueRelationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateIssueRelationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateIssueRelationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateIssueRelationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateIssueRelationRequestValidationError) ErrorName() string {
	return "CreateIssueRelationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateIssueRelationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateIssueRelationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateIssueRelationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateIssueRelationRequestValidationError{}

// Validate checks the field values on CreateIssueRelationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateIssueRelationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateIssueRelationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateIssueRelationResponseMultiError, or nil if none found.
func (m *CreateIssueRelationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateIssueRelationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIssueRelation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateIssueRelationResponseValidationError{
					field:  "IssueRelation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateIssueRelationResponseValidationError{
					field:  "IssueRelation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIssueRelation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateIssueRelationResponseValidationError{
				field:  "IssueRelation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateIssueRelationResponseMultiError(errors)
	}

	return nil
}

// CreateIssueRelationResponseMultiError is an error wrapping multiple
// validation errors returned by CreateIssueRelationResponse.ValidateAll() if
// the designated constraints aren't met.
type CreateIssueRelationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateIssueRelationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateIssueRelationResponseMultiError) AllErrors() []error { return m }

// CreateIssueRelationResponseValidationError is the validation error returned
// by CreateIssueRelationResponse.Validate if the designated constraints
// aren't met.
type CreateIssueRelationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateIssueRelationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateIssueRelationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateIssueRelationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateIssueRelationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateIssueRelationResponseValidationError) ErrorName() string {
	return "CreateIssueRelationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateIssueRelationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateIssueRelationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateIssueRelationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateIssueRelationResponseValidationError{}

// Validate checks the field values on DeleteIssueRelationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteIssueRelationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteIssueRelationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteIssueRelationRequestMultiError, or nil if none found.
func (m *DeleteIssueRelationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteIssueRelationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RelationID

	if len(errors) > 0 {
		return DeleteIssueRelationRequestMultiError(errors)
	}

	return nil
}

// DeleteIssueRelationRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteIssueRelationRequest.ValidateAll() if
// the designated constraints aren't met.
type DeleteIssueRelationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteIssueRelationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteIssueRelationRequestMultiError) AllErrors() []error { return m }

// DeleteIssueRelationRequestValidationError is the validation error returned
// by DeleteIssueRelationRequest.Validate if the designated constraints aren't met.
type DeleteIssueRelationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteIssueRelationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteIssueRelationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteIssueRelationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteIssueRelationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteIssueRelationRequestValidationError) ErrorName() string {
	return "DeleteIssueRelationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteIssueRelationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteIssueRelationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteIssueRelationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteIssueRelationRequestValidationError{}

// Validate checks the field values on DeleteIssueRelationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteIssueRelationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteIssueRelationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteIssueRelationResponseMultiError, or nil if none found.
func (m *DeleteIssueRelationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteIssueRelationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteIssueRelationResponseMultiError(errors)
	}

	return nil
}

// DeleteIssueRelationResponseMultiError is an error wrapping multiple
// validation errors returned by DeleteIssueRelationResponse.ValidateAll() if
// the designated constraints aren't met.
type DeleteIssueRelationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteIssueRelationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteIssueRelationResponseMultiError) AllErrors() []error { return m }

// DeleteIssueRelationResponseValidationError is the validation error returned
// by DeleteIssueRelationResponse.Validate if the designated constraints
// aren't met.
type DeleteIssueRelationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteIssueRelationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteIssueRelationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteIssueRelationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteIssueRelationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteIssueRelationResponseValidationError) ErrorName() string {
	return "DeleteIssueRelationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteIssueRelationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteIssueRelationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteIssueRelationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteIssueRelationResponseValidationError{}

// Validate checks the field values on ListIssueRelationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListIssueRelationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListIssueRelationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListIssueRelationRequestMultiError, or nil if none found.
func (m *ListIssueRelationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListIssueRelationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	if len(errors) > 0 {
		return ListIssueRelationRequestMultiError(errors)
	}

	return nil
}

// ListIssueRelationRequestMultiError is an error wrapping multiple validation
// errors returned by ListIssueRelationRequest.ValidateAll() if the designated
// constraints aren't met.
type ListIssueRelationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListIssueRelationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListIssueRelationRequestMultiError) AllErrors() []error { return m }

// ListIssueRelationRequestValidationError is the validation error returned by
// ListIssueRelationRequest.Validate if the designated constraints aren't met.
type ListIssueRelationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListIssueRelationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListIssueRelationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListIssueRelationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListIssueRelationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListIssueRelationRequestValidationError) ErrorName() string {
	return "ListIssueRelationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListIssueRelationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListIssueRelationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListIssueRelationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListIssueRelationRequestValidationError{}

// Validate checks the field values on ListIssueRelationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListIssueRelationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListIssueRelationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListIssueRelationResponseMultiError, or nil if none found.
func (m *ListIssueRelationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListIssueRelationResponse) validate(all bool) error {
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
					errors = append(errors, ListIssueRelationResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListIssueRelationResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListIssueRelationResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListIssueRelationResponseMultiError(errors)
	}

	return nil
}

// ListIssueRelationResponseMultiError is an error wrapping multiple validation
// errors returned by ListIssueRelationResponse.ValidateAll() if the
// designated constraints aren't met.
type ListIssueRelationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListIssueRelationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListIssueRelationResponseMultiError) AllErrors() []error { return m }

// ListIssueRelationResponseValidationError is the validation error returned by
// ListIssueRelationResponse.Validate if the designated constraints aren't met.
type ListIssueRelationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListIssueRelationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListIssueRelationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListIssueRelationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListIssueRelationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListIssueRelationResponseValidationError) ErrorName() string {
	return "ListIssueRelationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListIssueRelationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListIssueRelationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListIssueRelationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListIssueRelationResponseValidationError{}
