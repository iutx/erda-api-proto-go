// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: configcenter.proto

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

// Validate checks the field values on GetGroupRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetGroupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetGroupRequestMultiError, or nil if none found.
func (m *GetGroupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetGroupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTenantID()) < 1 {
		err := GetGroupRequestValidationError{
			field:  "TenantID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Keyword

	if m.GetPageNum() <= 0 {
		err := GetGroupRequestValidationError{
			field:  "PageNum",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPageSize() <= 0 {
		err := GetGroupRequestValidationError{
			field:  "PageSize",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetGroupRequestMultiError(errors)
	}

	return nil
}

// GetGroupRequestMultiError is an error wrapping multiple validation errors
// returned by GetGroupRequest.ValidateAll() if the designated constraints
// aren't met.
type GetGroupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetGroupRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetGroupRequestMultiError) AllErrors() []error { return m }

// GetGroupRequestValidationError is the validation error returned by
// GetGroupRequest.Validate if the designated constraints aren't met.
type GetGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroupRequestValidationError) ErrorName() string { return "GetGroupRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroupRequestValidationError{}

// Validate checks the field values on GetGroupResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetGroupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetGroupResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetGroupResponseMultiError, or nil if none found.
func (m *GetGroupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetGroupResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetGroupResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetGroupResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetGroupResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetGroupResponseMultiError(errors)
	}

	return nil
}

// GetGroupResponseMultiError is an error wrapping multiple validation errors
// returned by GetGroupResponse.ValidateAll() if the designated constraints
// aren't met.
type GetGroupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetGroupResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetGroupResponseMultiError) AllErrors() []error { return m }

// GetGroupResponseValidationError is the validation error returned by
// GetGroupResponse.Validate if the designated constraints aren't met.
type GetGroupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroupResponseValidationError) ErrorName() string { return "GetGroupResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetGroupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroupResponseValidationError{}

// Validate checks the field values on Groups with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Groups) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Groups with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in GroupsMultiError, or nil if none found.
func (m *Groups) ValidateAll() error {
	return m.validate(true)
}

func (m *Groups) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return GroupsMultiError(errors)
	}

	return nil
}

// GroupsMultiError is an error wrapping multiple validation errors returned by
// Groups.ValidateAll() if the designated constraints aren't met.
type GroupsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GroupsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GroupsMultiError) AllErrors() []error { return m }

// GroupsValidationError is the validation error returned by Groups.Validate if
// the designated constraints aren't met.
type GroupsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GroupsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GroupsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GroupsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GroupsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GroupsValidationError) ErrorName() string { return "GroupsValidationError" }

// Error satisfies the builtin error interface
func (e GroupsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGroups.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GroupsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GroupsValidationError{}

// Validate checks the field values on GetGroupPropertiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetGroupPropertiesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetGroupPropertiesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetGroupPropertiesRequestMultiError, or nil if none found.
func (m *GetGroupPropertiesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetGroupPropertiesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTenantID()) < 1 {
		err := GetGroupPropertiesRequestValidationError{
			field:  "TenantID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetGroupID()) < 1 {
		err := GetGroupPropertiesRequestValidationError{
			field:  "GroupID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetGroupPropertiesRequestMultiError(errors)
	}

	return nil
}

// GetGroupPropertiesRequestMultiError is an error wrapping multiple validation
// errors returned by GetGroupPropertiesRequest.ValidateAll() if the
// designated constraints aren't met.
type GetGroupPropertiesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetGroupPropertiesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetGroupPropertiesRequestMultiError) AllErrors() []error { return m }

// GetGroupPropertiesRequestValidationError is the validation error returned by
// GetGroupPropertiesRequest.Validate if the designated constraints aren't met.
type GetGroupPropertiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroupPropertiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroupPropertiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroupPropertiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroupPropertiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroupPropertiesRequestValidationError) ErrorName() string {
	return "GetGroupPropertiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetGroupPropertiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroupPropertiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroupPropertiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroupPropertiesRequestValidationError{}

// Validate checks the field values on GetGroupPropertiesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetGroupPropertiesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetGroupPropertiesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetGroupPropertiesResponseMultiError, or nil if none found.
func (m *GetGroupPropertiesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetGroupPropertiesResponse) validate(all bool) error {
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
					errors = append(errors, GetGroupPropertiesResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetGroupPropertiesResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetGroupPropertiesResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetGroupPropertiesResponseMultiError(errors)
	}

	return nil
}

// GetGroupPropertiesResponseMultiError is an error wrapping multiple
// validation errors returned by GetGroupPropertiesResponse.ValidateAll() if
// the designated constraints aren't met.
type GetGroupPropertiesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetGroupPropertiesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetGroupPropertiesResponseMultiError) AllErrors() []error { return m }

// GetGroupPropertiesResponseValidationError is the validation error returned
// by GetGroupPropertiesResponse.Validate if the designated constraints aren't met.
type GetGroupPropertiesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroupPropertiesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroupPropertiesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroupPropertiesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroupPropertiesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroupPropertiesResponseValidationError) ErrorName() string {
	return "GetGroupPropertiesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetGroupPropertiesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroupPropertiesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroupPropertiesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroupPropertiesResponseValidationError{}

// Validate checks the field values on SaveGroupPropertiesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveGroupPropertiesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveGroupPropertiesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveGroupPropertiesRequestMultiError, or nil if none found.
func (m *SaveGroupPropertiesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveGroupPropertiesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTenantID()) < 1 {
		err := SaveGroupPropertiesRequestValidationError{
			field:  "TenantID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetGroupID()) < 1 {
		err := SaveGroupPropertiesRequestValidationError{
			field:  "GroupID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetProperties() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SaveGroupPropertiesRequestValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SaveGroupPropertiesRequestValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SaveGroupPropertiesRequestValidationError{
					field:  fmt.Sprintf("Properties[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SaveGroupPropertiesRequestMultiError(errors)
	}

	return nil
}

// SaveGroupPropertiesRequestMultiError is an error wrapping multiple
// validation errors returned by SaveGroupPropertiesRequest.ValidateAll() if
// the designated constraints aren't met.
type SaveGroupPropertiesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveGroupPropertiesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveGroupPropertiesRequestMultiError) AllErrors() []error { return m }

// SaveGroupPropertiesRequestValidationError is the validation error returned
// by SaveGroupPropertiesRequest.Validate if the designated constraints aren't met.
type SaveGroupPropertiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveGroupPropertiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveGroupPropertiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveGroupPropertiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveGroupPropertiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveGroupPropertiesRequestValidationError) ErrorName() string {
	return "SaveGroupPropertiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SaveGroupPropertiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveGroupPropertiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveGroupPropertiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveGroupPropertiesRequestValidationError{}

// Validate checks the field values on SaveGroupPropertiesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveGroupPropertiesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveGroupPropertiesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveGroupPropertiesResponseMultiError, or nil if none found.
func (m *SaveGroupPropertiesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveGroupPropertiesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return SaveGroupPropertiesResponseMultiError(errors)
	}

	return nil
}

// SaveGroupPropertiesResponseMultiError is an error wrapping multiple
// validation errors returned by SaveGroupPropertiesResponse.ValidateAll() if
// the designated constraints aren't met.
type SaveGroupPropertiesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveGroupPropertiesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveGroupPropertiesResponseMultiError) AllErrors() []error { return m }

// SaveGroupPropertiesResponseValidationError is the validation error returned
// by SaveGroupPropertiesResponse.Validate if the designated constraints
// aren't met.
type SaveGroupPropertiesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveGroupPropertiesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveGroupPropertiesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveGroupPropertiesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveGroupPropertiesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveGroupPropertiesResponseValidationError) ErrorName() string {
	return "SaveGroupPropertiesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SaveGroupPropertiesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveGroupPropertiesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveGroupPropertiesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveGroupPropertiesResponseValidationError{}

// Validate checks the field values on GroupProperties with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GroupProperties) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GroupProperties with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GroupPropertiesMultiError, or nil if none found.
func (m *GroupProperties) ValidateAll() error {
	return m.validate(true)
}

func (m *GroupProperties) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Group

	for idx, item := range m.GetProperties() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GroupPropertiesValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GroupPropertiesValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GroupPropertiesValidationError{
					field:  fmt.Sprintf("Properties[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GroupPropertiesMultiError(errors)
	}

	return nil
}

// GroupPropertiesMultiError is an error wrapping multiple validation errors
// returned by GroupProperties.ValidateAll() if the designated constraints
// aren't met.
type GroupPropertiesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GroupPropertiesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GroupPropertiesMultiError) AllErrors() []error { return m }

// GroupPropertiesValidationError is the validation error returned by
// GroupProperties.Validate if the designated constraints aren't met.
type GroupPropertiesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GroupPropertiesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GroupPropertiesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GroupPropertiesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GroupPropertiesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GroupPropertiesValidationError) ErrorName() string { return "GroupPropertiesValidationError" }

// Error satisfies the builtin error interface
func (e GroupPropertiesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGroupProperties.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GroupPropertiesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GroupPropertiesValidationError{}

// Validate checks the field values on Property with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Property) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Property with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PropertyMultiError, or nil
// if none found.
func (m *Property) ValidateAll() error {
	return m.validate(true)
}

func (m *Property) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	// no validation rules for Source

	if len(errors) > 0 {
		return PropertyMultiError(errors)
	}

	return nil
}

// PropertyMultiError is an error wrapping multiple validation errors returned
// by Property.ValidateAll() if the designated constraints aren't met.
type PropertyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PropertyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PropertyMultiError) AllErrors() []error { return m }

// PropertyValidationError is the validation error returned by
// Property.Validate if the designated constraints aren't met.
type PropertyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PropertyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PropertyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PropertyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PropertyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PropertyValidationError) ErrorName() string { return "PropertyValidationError" }

// Error satisfies the builtin error interface
func (e PropertyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProperty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PropertyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PropertyValidationError{}
