// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: micro_api.proto

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

// Validate checks the field values on DeleteApiRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteApiRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteApiRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteApiRequestMultiError, or nil if none found.
func (m *DeleteApiRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteApiRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ApiId

	if len(errors) > 0 {
		return DeleteApiRequestMultiError(errors)
	}

	return nil
}

// DeleteApiRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteApiRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteApiRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteApiRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteApiRequestMultiError) AllErrors() []error { return m }

// DeleteApiRequestValidationError is the validation error returned by
// DeleteApiRequest.Validate if the designated constraints aren't met.
type DeleteApiRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteApiRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteApiRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteApiRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteApiRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteApiRequestValidationError) ErrorName() string { return "DeleteApiRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteApiRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteApiRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteApiRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteApiRequestValidationError{}

// Validate checks the field values on DeleteApiResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteApiResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteApiResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteApiResponseMultiError, or nil if none found.
func (m *DeleteApiResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteApiResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return DeleteApiResponseMultiError(errors)
	}

	return nil
}

// DeleteApiResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteApiResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteApiResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteApiResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteApiResponseMultiError) AllErrors() []error { return m }

// DeleteApiResponseValidationError is the validation error returned by
// DeleteApiResponse.Validate if the designated constraints aren't met.
type DeleteApiResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteApiResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteApiResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteApiResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteApiResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteApiResponseValidationError) ErrorName() string {
	return "DeleteApiResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteApiResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteApiResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteApiResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteApiResponseValidationError{}

// Validate checks the field values on ApiRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ApiRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApiRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ApiRequestMultiError, or
// nil if none found.
func (m *ApiRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ApiRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Path

	// no validation rules for Method

	// no validation rules for RedirectType

	// no validation rules for RedirectAddr

	// no validation rules for RedirectPath

	// no validation rules for ProjectId

	// no validation rules for Description

	// no validation rules for DiceApp

	// no validation rules for DiceService

	// no validation rules for OuterNetEnable

	// no validation rules for RuntimeId

	// no validation rules for Env

	// no validation rules for ConsumerId

	if len(errors) > 0 {
		return ApiRequestMultiError(errors)
	}

	return nil
}

// ApiRequestMultiError is an error wrapping multiple validation errors
// returned by ApiRequest.ValidateAll() if the designated constraints aren't met.
type ApiRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApiRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApiRequestMultiError) AllErrors() []error { return m }

// ApiRequestValidationError is the validation error returned by
// ApiRequest.Validate if the designated constraints aren't met.
type ApiRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApiRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApiRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApiRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApiRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApiRequestValidationError) ErrorName() string { return "ApiRequestValidationError" }

// Error satisfies the builtin error interface
func (e ApiRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApiRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApiRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApiRequestValidationError{}

// Validate checks the field values on UpdateApiRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateApiRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateApiRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateApiRequestMultiError, or nil if none found.
func (m *UpdateApiRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateApiRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetApiRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateApiRequestValidationError{
					field:  "ApiRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateApiRequestValidationError{
					field:  "ApiRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApiRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateApiRequestValidationError{
				field:  "ApiRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ApiId

	if len(errors) > 0 {
		return UpdateApiRequestMultiError(errors)
	}

	return nil
}

// UpdateApiRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateApiRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateApiRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateApiRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateApiRequestMultiError) AllErrors() []error { return m }

// UpdateApiRequestValidationError is the validation error returned by
// UpdateApiRequest.Validate if the designated constraints aren't met.
type UpdateApiRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateApiRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateApiRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateApiRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateApiRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateApiRequestValidationError) ErrorName() string { return "UpdateApiRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateApiRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateApiRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateApiRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateApiRequestValidationError{}

// Validate checks the field values on Policy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Policy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Policy with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PolicyMultiError, or nil if none found.
func (m *Policy) ValidateAll() error {
	return m.validate(true)
}

func (m *Policy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Category

	// no validation rules for PolicyId

	// no validation rules for PolicyName

	// no validation rules for DisplayName

	{
		sorted_keys := make([]string, len(m.GetConfig()))
		i := 0
		for key := range m.GetConfig() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetConfig()[key]
			_ = val

			// no validation rules for Config[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, PolicyValidationError{
							field:  fmt.Sprintf("Config[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, PolicyValidationError{
							field:  fmt.Sprintf("Config[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return PolicyValidationError{
						field:  fmt.Sprintf("Config[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for CreateAt

	if len(errors) > 0 {
		return PolicyMultiError(errors)
	}

	return nil
}

// PolicyMultiError is an error wrapping multiple validation errors returned by
// Policy.ValidateAll() if the designated constraints aren't met.
type PolicyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PolicyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PolicyMultiError) AllErrors() []error { return m }

// PolicyValidationError is the validation error returned by Policy.Validate if
// the designated constraints aren't met.
type PolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolicyValidationError) ErrorName() string { return "PolicyValidationError" }

// Error satisfies the builtin error interface
func (e PolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolicyValidationError{}

// Validate checks the field values on UpdateApiResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateApiResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateApiResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateApiResponseMultiError, or nil if none found.
func (m *UpdateApiResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateApiResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ApiId

	// no validation rules for Path

	// no validation rules for DisplayPath

	if all {
		switch v := interface{}(m.GetDisplayPathPrefix()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateApiResponseValidationError{
					field:  "DisplayPathPrefix",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateApiResponseValidationError{
					field:  "DisplayPathPrefix",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDisplayPathPrefix()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateApiResponseValidationError{
				field:  "DisplayPathPrefix",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OuterNetEnable

	// no validation rules for RegisterType

	// no validation rules for NeedAuth

	if all {
		switch v := interface{}(m.GetMethod()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateApiResponseValidationError{
					field:  "Method",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateApiResponseValidationError{
					field:  "Method",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMethod()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateApiResponseValidationError{
				field:  "Method",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Description

	// no validation rules for RedirectAddr

	// no validation rules for RedirectPath

	// no validation rules for RedirectType

	// no validation rules for MonitorPath

	// no validation rules for CreateAt

	for idx, item := range m.GetPolicies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpdateApiResponseValidationError{
						field:  fmt.Sprintf("Policies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpdateApiResponseValidationError{
						field:  fmt.Sprintf("Policies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateApiResponseValidationError{
					field:  fmt.Sprintf("Policies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UpdateApiResponseMultiError(errors)
	}

	return nil
}

// UpdateApiResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateApiResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateApiResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateApiResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateApiResponseMultiError) AllErrors() []error { return m }

// UpdateApiResponseValidationError is the validation error returned by
// UpdateApiResponse.Validate if the designated constraints aren't met.
type UpdateApiResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateApiResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateApiResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateApiResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateApiResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateApiResponseValidationError) ErrorName() string {
	return "UpdateApiResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateApiResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateApiResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateApiResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateApiResponseValidationError{}

// Validate checks the field values on GetApisRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetApisRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetApisRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetApisRequestMultiError,
// or nil if none found.
func (m *GetApisRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetApisRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for From

	// no validation rules for Method

	// no validation rules for DiceApp

	// no validation rules for DiceService

	// no validation rules for RuntimeId

	// no validation rules for ApiPath

	// no validation rules for RegisterType

	// no validation rules for NetType

	// no validation rules for NeedAuth

	// no validation rules for SortField

	// no validation rules for SortType

	// no validation rules for OrgId

	// no validation rules for ProjectId

	// no validation rules for Env

	// no validation rules for Page

	// no validation rules for Size

	if len(errors) > 0 {
		return GetApisRequestMultiError(errors)
	}

	return nil
}

// GetApisRequestMultiError is an error wrapping multiple validation errors
// returned by GetApisRequest.ValidateAll() if the designated constraints
// aren't met.
type GetApisRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetApisRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetApisRequestMultiError) AllErrors() []error { return m }

// GetApisRequestValidationError is the validation error returned by
// GetApisRequest.Validate if the designated constraints aren't met.
type GetApisRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApisRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApisRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetApisRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApisRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApisRequestValidationError) ErrorName() string { return "GetApisRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetApisRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApisRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApisRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApisRequestValidationError{}

// Validate checks the field values on GetApisResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetApisResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetApisResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetApisResponseMultiError, or nil if none found.
func (m *GetApisResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetApisResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetApisResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetApisResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetApisResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetApisResponseMultiError(errors)
	}

	return nil
}

// GetApisResponseMultiError is an error wrapping multiple validation errors
// returned by GetApisResponse.ValidateAll() if the designated constraints
// aren't met.
type GetApisResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetApisResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetApisResponseMultiError) AllErrors() []error { return m }

// GetApisResponseValidationError is the validation error returned by
// GetApisResponse.Validate if the designated constraints aren't met.
type GetApisResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApisResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApisResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetApisResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApisResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApisResponseValidationError) ErrorName() string { return "GetApisResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetApisResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApisResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApisResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApisResponseValidationError{}

// Validate checks the field values on CreateApiRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateApiRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateApiRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateApiRequestMultiError, or nil if none found.
func (m *CreateApiRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateApiRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetApiRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateApiRequestValidationError{
					field:  "ApiRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateApiRequestValidationError{
					field:  "ApiRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApiRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateApiRequestValidationError{
				field:  "ApiRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateApiRequestMultiError(errors)
	}

	return nil
}

// CreateApiRequestMultiError is an error wrapping multiple validation errors
// returned by CreateApiRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateApiRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateApiRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateApiRequestMultiError) AllErrors() []error { return m }

// CreateApiRequestValidationError is the validation error returned by
// CreateApiRequest.Validate if the designated constraints aren't met.
type CreateApiRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateApiRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateApiRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateApiRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateApiRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateApiRequestValidationError) ErrorName() string { return "CreateApiRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateApiRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateApiRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateApiRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateApiRequestValidationError{}

// Validate checks the field values on CreateApiResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateApiResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateApiResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateApiResponseMultiError, or nil if none found.
func (m *CreateApiResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateApiResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ApiId

	if len(errors) > 0 {
		return CreateApiResponseMultiError(errors)
	}

	return nil
}

// CreateApiResponseMultiError is an error wrapping multiple validation errors
// returned by CreateApiResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateApiResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateApiResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateApiResponseMultiError) AllErrors() []error { return m }

// CreateApiResponseValidationError is the validation error returned by
// CreateApiResponse.Validate if the designated constraints aren't met.
type CreateApiResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateApiResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateApiResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateApiResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateApiResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateApiResponseValidationError) ErrorName() string {
	return "CreateApiResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateApiResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateApiResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateApiResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateApiResponseValidationError{}