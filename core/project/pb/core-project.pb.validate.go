// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: core-project.proto

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

// Validate checks the field values on CheckProjectExistReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckProjectExistReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckProjectExistReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckProjectExistReqMultiError, or nil if none found.
func (m *CheckProjectExistReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckProjectExistReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CheckProjectExistReqMultiError(errors)
	}

	return nil
}

// CheckProjectExistReqMultiError is an error wrapping multiple validation
// errors returned by CheckProjectExistReq.ValidateAll() if the designated
// constraints aren't met.
type CheckProjectExistReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckProjectExistReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckProjectExistReqMultiError) AllErrors() []error { return m }

// CheckProjectExistReqValidationError is the validation error returned by
// CheckProjectExistReq.Validate if the designated constraints aren't met.
type CheckProjectExistReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckProjectExistReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckProjectExistReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckProjectExistReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckProjectExistReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckProjectExistReqValidationError) ErrorName() string {
	return "CheckProjectExistReqValidationError"
}

// Error satisfies the builtin error interface
func (e CheckProjectExistReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckProjectExistReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckProjectExistReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckProjectExistReqValidationError{}

// Validate checks the field values on CheckProjectExistResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckProjectExistResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckProjectExistResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckProjectExistRespMultiError, or nil if none found.
func (m *CheckProjectExistResp) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckProjectExistResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return CheckProjectExistRespMultiError(errors)
	}

	return nil
}

// CheckProjectExistRespMultiError is an error wrapping multiple validation
// errors returned by CheckProjectExistResp.ValidateAll() if the designated
// constraints aren't met.
type CheckProjectExistRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckProjectExistRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckProjectExistRespMultiError) AllErrors() []error { return m }

// CheckProjectExistRespValidationError is the validation error returned by
// CheckProjectExistResp.Validate if the designated constraints aren't met.
type CheckProjectExistRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckProjectExistRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckProjectExistRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckProjectExistRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckProjectExistRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckProjectExistRespValidationError) ErrorName() string {
	return "CheckProjectExistRespValidationError"
}

// Error satisfies the builtin error interface
func (e CheckProjectExistRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckProjectExistResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckProjectExistRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckProjectExistRespValidationError{}

// Validate checks the field values on GetProjectByIDReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProjectByIDReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProjectByIDReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProjectByIDReqMultiError, or nil if none found.
func (m *GetProjectByIDReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProjectByIDReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.UserID != nil {
		// no validation rules for UserID
	}

	if len(errors) > 0 {
		return GetProjectByIDReqMultiError(errors)
	}

	return nil
}

// GetProjectByIDReqMultiError is an error wrapping multiple validation errors
// returned by GetProjectByIDReq.ValidateAll() if the designated constraints
// aren't met.
type GetProjectByIDReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProjectByIDReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProjectByIDReqMultiError) AllErrors() []error { return m }

// GetProjectByIDReqValidationError is the validation error returned by
// GetProjectByIDReq.Validate if the designated constraints aren't met.
type GetProjectByIDReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectByIDReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectByIDReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectByIDReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectByIDReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectByIDReqValidationError) ErrorName() string {
	return "GetProjectByIDReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectByIDReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectByIDReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectByIDReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectByIDReqValidationError{}

// Validate checks the field values on ProjectDto with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProjectDto) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProjectDto with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProjectDtoMultiError, or
// nil if none found.
func (m *ProjectDto) ValidateAll() error {
	return m.validate(true)
}

func (m *ProjectDto) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for DisplayName

	// no validation rules for OrgID

	// no validation rules for CreatorID

	// no validation rules for Logo

	// no validation rules for Desc

	if all {
		switch v := interface{}(m.GetActiveTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectDtoValidationError{
					field:  "ActiveTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectDtoValidationError{
					field:  "ActiveTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetActiveTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectDtoValidationError{
				field:  "ActiveTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsPublic

	if all {
		switch v := interface{}(m.GetCreatedTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectDtoValidationError{
					field:  "CreatedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectDtoValidationError{
					field:  "CreatedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectDtoValidationError{
				field:  "CreatedTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectDtoValidationError{
					field:  "UpdatedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectDtoValidationError{
					field:  "UpdatedTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectDtoValidationError{
				field:  "UpdatedTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Type

	if len(errors) > 0 {
		return ProjectDtoMultiError(errors)
	}

	return nil
}

// ProjectDtoMultiError is an error wrapping multiple validation errors
// returned by ProjectDto.ValidateAll() if the designated constraints aren't met.
type ProjectDtoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProjectDtoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProjectDtoMultiError) AllErrors() []error { return m }

// ProjectDtoValidationError is the validation error returned by
// ProjectDto.Validate if the designated constraints aren't met.
type ProjectDtoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDtoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDtoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDtoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDtoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDtoValidationError) ErrorName() string { return "ProjectDtoValidationError" }

// Error satisfies the builtin error interface
func (e ProjectDtoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDto.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDtoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDtoValidationError{}
