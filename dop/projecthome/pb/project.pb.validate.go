// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: project.proto

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

// Validate checks the field values on ProjectHome with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProjectHome) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProjectHome with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProjectHomeMultiError, or
// nil if none found.
func (m *ProjectHome) ValidateAll() error {
	return m.validate(true)
}

func (m *ProjectHome) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Readme

	for idx, item := range m.GetLinks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProjectHomeValidationError{
						field:  fmt.Sprintf("Links[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProjectHomeValidationError{
						field:  fmt.Sprintf("Links[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectHomeValidationError{
					field:  fmt.Sprintf("Links[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ProjectHomeMultiError(errors)
	}

	return nil
}

// ProjectHomeMultiError is an error wrapping multiple validation errors
// returned by ProjectHome.ValidateAll() if the designated constraints aren't met.
type ProjectHomeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProjectHomeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProjectHomeMultiError) AllErrors() []error { return m }

// ProjectHomeValidationError is the validation error returned by
// ProjectHome.Validate if the designated constraints aren't met.
type ProjectHomeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectHomeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectHomeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectHomeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectHomeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectHomeValidationError) ErrorName() string { return "ProjectHomeValidationError" }

// Error satisfies the builtin error interface
func (e ProjectHomeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectHome.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectHomeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectHomeValidationError{}

// Validate checks the field values on Link with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Link) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Link with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in LinkMultiError, or nil if none found.
func (m *Link) ValidateAll() error {
	return m.validate(true)
}

func (m *Link) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Url

	if len(errors) > 0 {
		return LinkMultiError(errors)
	}

	return nil
}

// LinkMultiError is an error wrapping multiple validation errors returned by
// Link.ValidateAll() if the designated constraints aren't met.
type LinkMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LinkMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LinkMultiError) AllErrors() []error { return m }

// LinkValidationError is the validation error returned by Link.Validate if the
// designated constraints aren't met.
type LinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LinkValidationError) ErrorName() string { return "LinkValidationError" }

// Error satisfies the builtin error interface
func (e LinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LinkValidationError{}

// Validate checks the field values on GetProjectHomeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProjectHomeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProjectHomeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProjectHomeRequestMultiError, or nil if none found.
func (m *GetProjectHomeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProjectHomeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectID

	if len(errors) > 0 {
		return GetProjectHomeRequestMultiError(errors)
	}

	return nil
}

// GetProjectHomeRequestMultiError is an error wrapping multiple validation
// errors returned by GetProjectHomeRequest.ValidateAll() if the designated
// constraints aren't met.
type GetProjectHomeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProjectHomeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProjectHomeRequestMultiError) AllErrors() []error { return m }

// GetProjectHomeRequestValidationError is the validation error returned by
// GetProjectHomeRequest.Validate if the designated constraints aren't met.
type GetProjectHomeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectHomeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectHomeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectHomeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectHomeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectHomeRequestValidationError) ErrorName() string {
	return "GetProjectHomeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectHomeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectHomeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectHomeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectHomeRequestValidationError{}

// Validate checks the field values on GetProjectHomeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProjectHomeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProjectHomeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProjectHomeResponseMultiError, or nil if none found.
func (m *GetProjectHomeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProjectHomeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProjectHomeResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProjectHomeResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProjectHomeResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetProjectHomeResponseMultiError(errors)
	}

	return nil
}

// GetProjectHomeResponseMultiError is an error wrapping multiple validation
// errors returned by GetProjectHomeResponse.ValidateAll() if the designated
// constraints aren't met.
type GetProjectHomeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProjectHomeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProjectHomeResponseMultiError) AllErrors() []error { return m }

// GetProjectHomeResponseValidationError is the validation error returned by
// GetProjectHomeResponse.Validate if the designated constraints aren't met.
type GetProjectHomeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectHomeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectHomeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectHomeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectHomeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectHomeResponseValidationError) ErrorName() string {
	return "GetProjectHomeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectHomeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectHomeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectHomeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectHomeResponseValidationError{}

// Validate checks the field values on CreateOrUpdateProjectHomeRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CreateOrUpdateProjectHomeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrUpdateProjectHomeRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreateOrUpdateProjectHomeRequestMultiError, or nil if none found.
func (m *CreateOrUpdateProjectHomeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrUpdateProjectHomeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectID

	// no validation rules for Readme

	for idx, item := range m.GetLinks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateOrUpdateProjectHomeRequestValidationError{
						field:  fmt.Sprintf("Links[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateOrUpdateProjectHomeRequestValidationError{
						field:  fmt.Sprintf("Links[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrUpdateProjectHomeRequestValidationError{
					field:  fmt.Sprintf("Links[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateOrUpdateProjectHomeRequestMultiError(errors)
	}

	return nil
}

// CreateOrUpdateProjectHomeRequestMultiError is an error wrapping multiple
// validation errors returned by
// CreateOrUpdateProjectHomeRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateOrUpdateProjectHomeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrUpdateProjectHomeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrUpdateProjectHomeRequestMultiError) AllErrors() []error { return m }

// CreateOrUpdateProjectHomeRequestValidationError is the validation error
// returned by CreateOrUpdateProjectHomeRequest.Validate if the designated
// constraints aren't met.
type CreateOrUpdateProjectHomeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrUpdateProjectHomeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrUpdateProjectHomeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrUpdateProjectHomeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrUpdateProjectHomeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrUpdateProjectHomeRequestValidationError) ErrorName() string {
	return "CreateOrUpdateProjectHomeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrUpdateProjectHomeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrUpdateProjectHomeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrUpdateProjectHomeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrUpdateProjectHomeRequestValidationError{}

// Validate checks the field values on CreateOrUpdateProjectHomeResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CreateOrUpdateProjectHomeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrUpdateProjectHomeResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CreateOrUpdateProjectHomeResponseMultiError, or nil if none found.
func (m *CreateOrUpdateProjectHomeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrUpdateProjectHomeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateOrUpdateProjectHomeResponseMultiError(errors)
	}

	return nil
}

// CreateOrUpdateProjectHomeResponseMultiError is an error wrapping multiple
// validation errors returned by
// CreateOrUpdateProjectHomeResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateOrUpdateProjectHomeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrUpdateProjectHomeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrUpdateProjectHomeResponseMultiError) AllErrors() []error { return m }

// CreateOrUpdateProjectHomeResponseValidationError is the validation error
// returned by CreateOrUpdateProjectHomeResponse.Validate if the designated
// constraints aren't met.
type CreateOrUpdateProjectHomeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrUpdateProjectHomeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrUpdateProjectHomeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrUpdateProjectHomeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrUpdateProjectHomeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrUpdateProjectHomeResponseValidationError) ErrorName() string {
	return "CreateOrUpdateProjectHomeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrUpdateProjectHomeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrUpdateProjectHomeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrUpdateProjectHomeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrUpdateProjectHomeResponseValidationError{}
