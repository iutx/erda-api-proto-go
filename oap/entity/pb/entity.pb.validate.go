// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: entity.proto

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

// Validate checks the field values on Entity with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Entity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Entity with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EntityMultiError, or nil if none found.
func (m *Entity) ValidateAll() error {
	return m.validate(true)
}

func (m *Entity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetType()) < 1 {
		err := EntityValidationError{
			field:  "Type",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := EntityValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	{
		sorted_keys := make([]string, len(m.GetValues()))
		i := 0
		for key := range m.GetValues() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetValues()[key]
			_ = val

			// no validation rules for Values[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, EntityValidationError{
							field:  fmt.Sprintf("Values[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, EntityValidationError{
							field:  fmt.Sprintf("Values[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return EntityValidationError{
						field:  fmt.Sprintf("Values[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for Labels

	// no validation rules for CreateTimeUnixNano

	// no validation rules for UpdateTimeUnixNano

	if len(errors) > 0 {
		return EntityMultiError(errors)
	}

	return nil
}

// EntityMultiError is an error wrapping multiple validation errors returned by
// Entity.ValidateAll() if the designated constraints aren't met.
type EntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntityMultiError) AllErrors() []error { return m }

// EntityValidationError is the validation error returned by Entity.Validate if
// the designated constraints aren't met.
type EntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityValidationError) ErrorName() string { return "EntityValidationError" }

// Error satisfies the builtin error interface
func (e EntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityValidationError{}

// Validate checks the field values on SetEntityRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SetEntityRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetEntityRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetEntityRequestMultiError, or nil if none found.
func (m *SetEntityRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SetEntityRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetData() == nil {
		err := SetEntityRequestValidationError{
			field:  "Data",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SetEntityRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SetEntityRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetEntityRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SetEntityRequestMultiError(errors)
	}

	return nil
}

// SetEntityRequestMultiError is an error wrapping multiple validation errors
// returned by SetEntityRequest.ValidateAll() if the designated constraints
// aren't met.
type SetEntityRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetEntityRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetEntityRequestMultiError) AllErrors() []error { return m }

// SetEntityRequestValidationError is the validation error returned by
// SetEntityRequest.Validate if the designated constraints aren't met.
type SetEntityRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetEntityRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetEntityRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetEntityRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetEntityRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetEntityRequestValidationError) ErrorName() string { return "SetEntityRequestValidationError" }

// Error satisfies the builtin error interface
func (e SetEntityRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetEntityRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetEntityRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetEntityRequestValidationError{}

// Validate checks the field values on SetEntityResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SetEntityResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetEntityResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetEntityResponseMultiError, or nil if none found.
func (m *SetEntityResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SetEntityResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return SetEntityResponseMultiError(errors)
	}

	return nil
}

// SetEntityResponseMultiError is an error wrapping multiple validation errors
// returned by SetEntityResponse.ValidateAll() if the designated constraints
// aren't met.
type SetEntityResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetEntityResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetEntityResponseMultiError) AllErrors() []error { return m }

// SetEntityResponseValidationError is the validation error returned by
// SetEntityResponse.Validate if the designated constraints aren't met.
type SetEntityResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetEntityResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetEntityResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetEntityResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetEntityResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetEntityResponseValidationError) ErrorName() string {
	return "SetEntityResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SetEntityResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetEntityResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetEntityResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetEntityResponseValidationError{}

// Validate checks the field values on RemoveEntityRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveEntityRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveEntityRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveEntityRequestMultiError, or nil if none found.
func (m *RemoveEntityRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveEntityRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetType()) < 1 {
		err := RemoveEntityRequestValidationError{
			field:  "Type",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := RemoveEntityRequestValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RemoveEntityRequestMultiError(errors)
	}

	return nil
}

// RemoveEntityRequestMultiError is an error wrapping multiple validation
// errors returned by RemoveEntityRequest.ValidateAll() if the designated
// constraints aren't met.
type RemoveEntityRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveEntityRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveEntityRequestMultiError) AllErrors() []error { return m }

// RemoveEntityRequestValidationError is the validation error returned by
// RemoveEntityRequest.Validate if the designated constraints aren't met.
type RemoveEntityRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveEntityRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveEntityRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveEntityRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveEntityRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveEntityRequestValidationError) ErrorName() string {
	return "RemoveEntityRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveEntityRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveEntityRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveEntityRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveEntityRequestValidationError{}

// Validate checks the field values on RemoveEntityResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveEntityResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveEntityResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveEntityResponseMultiError, or nil if none found.
func (m *RemoveEntityResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveEntityResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return RemoveEntityResponseMultiError(errors)
	}

	return nil
}

// RemoveEntityResponseMultiError is an error wrapping multiple validation
// errors returned by RemoveEntityResponse.ValidateAll() if the designated
// constraints aren't met.
type RemoveEntityResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveEntityResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveEntityResponseMultiError) AllErrors() []error { return m }

// RemoveEntityResponseValidationError is the validation error returned by
// RemoveEntityResponse.Validate if the designated constraints aren't met.
type RemoveEntityResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveEntityResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveEntityResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveEntityResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveEntityResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveEntityResponseValidationError) ErrorName() string {
	return "RemoveEntityResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveEntityResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveEntityResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveEntityResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveEntityResponseValidationError{}

// Validate checks the field values on GetEntityRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetEntityRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEntityRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEntityRequestMultiError, or nil if none found.
func (m *GetEntityRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEntityRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetType()) < 1 {
		err := GetEntityRequestValidationError{
			field:  "Type",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := GetEntityRequestValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetEntityRequestMultiError(errors)
	}

	return nil
}

// GetEntityRequestMultiError is an error wrapping multiple validation errors
// returned by GetEntityRequest.ValidateAll() if the designated constraints
// aren't met.
type GetEntityRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEntityRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEntityRequestMultiError) AllErrors() []error { return m }

// GetEntityRequestValidationError is the validation error returned by
// GetEntityRequest.Validate if the designated constraints aren't met.
type GetEntityRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEntityRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEntityRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEntityRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEntityRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEntityRequestValidationError) ErrorName() string { return "GetEntityRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetEntityRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEntityRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEntityRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEntityRequestValidationError{}

// Validate checks the field values on GetEntityResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetEntityResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEntityResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEntityResponseMultiError, or nil if none found.
func (m *GetEntityResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEntityResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetEntityResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetEntityResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEntityResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetEntityResponseMultiError(errors)
	}

	return nil
}

// GetEntityResponseMultiError is an error wrapping multiple validation errors
// returned by GetEntityResponse.ValidateAll() if the designated constraints
// aren't met.
type GetEntityResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEntityResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEntityResponseMultiError) AllErrors() []error { return m }

// GetEntityResponseValidationError is the validation error returned by
// GetEntityResponse.Validate if the designated constraints aren't met.
type GetEntityResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEntityResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEntityResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEntityResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEntityResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEntityResponseValidationError) ErrorName() string {
	return "GetEntityResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEntityResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEntityResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEntityResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEntityResponseValidationError{}

// Validate checks the field values on ListEntitiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListEntitiesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEntitiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListEntitiesRequestMultiError, or nil if none found.
func (m *ListEntitiesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEntitiesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Labels

	// no validation rules for Limit

	// no validation rules for UpdateTimeUnixNanoMin

	// no validation rules for UpdateTimeUnixNanoMax

	// no validation rules for Debug

	if len(errors) > 0 {
		return ListEntitiesRequestMultiError(errors)
	}

	return nil
}

// ListEntitiesRequestMultiError is an error wrapping multiple validation
// errors returned by ListEntitiesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListEntitiesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEntitiesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEntitiesRequestMultiError) AllErrors() []error { return m }

// ListEntitiesRequestValidationError is the validation error returned by
// ListEntitiesRequest.Validate if the designated constraints aren't met.
type ListEntitiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEntitiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEntitiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEntitiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEntitiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEntitiesRequestValidationError) ErrorName() string {
	return "ListEntitiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListEntitiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEntitiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEntitiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEntitiesRequestValidationError{}

// Validate checks the field values on ListEntitiesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListEntitiesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEntitiesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListEntitiesResponseMultiError, or nil if none found.
func (m *ListEntitiesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEntitiesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEntitiesResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEntitiesResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEntitiesResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListEntitiesResponseMultiError(errors)
	}

	return nil
}

// ListEntitiesResponseMultiError is an error wrapping multiple validation
// errors returned by ListEntitiesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListEntitiesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEntitiesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEntitiesResponseMultiError) AllErrors() []error { return m }

// ListEntitiesResponseValidationError is the validation error returned by
// ListEntitiesResponse.Validate if the designated constraints aren't met.
type ListEntitiesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEntitiesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEntitiesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEntitiesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEntitiesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEntitiesResponseValidationError) ErrorName() string {
	return "ListEntitiesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListEntitiesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEntitiesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEntitiesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEntitiesResponseValidationError{}

// Validate checks the field values on EntityList with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EntityList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EntityList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EntityListMultiError, or
// nil if none found.
func (m *EntityList) ValidateAll() error {
	return m.validate(true)
}

func (m *EntityList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntityListValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntityListValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityListValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return EntityListMultiError(errors)
	}

	return nil
}

// EntityListMultiError is an error wrapping multiple validation errors
// returned by EntityList.ValidateAll() if the designated constraints aren't met.
type EntityListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntityListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntityListMultiError) AllErrors() []error { return m }

// EntityListValidationError is the validation error returned by
// EntityList.Validate if the designated constraints aren't met.
type EntityListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityListValidationError) ErrorName() string { return "EntityListValidationError" }

// Error satisfies the builtin error interface
func (e EntityListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntityList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityListValidationError{}