// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: openapi_rule.proto

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

// Validate checks the field values on DeleteLimitResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteLimitResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteLimitResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteLimitResponseMultiError, or nil if none found.
func (m *DeleteLimitResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteLimitResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return DeleteLimitResponseMultiError(errors)
	}

	return nil
}

// DeleteLimitResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteLimitResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteLimitResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteLimitResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteLimitResponseMultiError) AllErrors() []error { return m }

// DeleteLimitResponseValidationError is the validation error returned by
// DeleteLimitResponse.Validate if the designated constraints aren't met.
type DeleteLimitResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteLimitResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteLimitResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteLimitResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteLimitResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteLimitResponseValidationError) ErrorName() string {
	return "DeleteLimitResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteLimitResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteLimitResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteLimitResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteLimitResponseValidationError{}

// Validate checks the field values on DeleteLimitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteLimitRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteLimitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteLimitRequestMultiError, or nil if none found.
func (m *DeleteLimitRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteLimitRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RuleId

	if len(errors) > 0 {
		return DeleteLimitRequestMultiError(errors)
	}

	return nil
}

// DeleteLimitRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteLimitRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteLimitRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteLimitRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteLimitRequestMultiError) AllErrors() []error { return m }

// DeleteLimitRequestValidationError is the validation error returned by
// DeleteLimitRequest.Validate if the designated constraints aren't met.
type DeleteLimitRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteLimitRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteLimitRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteLimitRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteLimitRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteLimitRequestValidationError) ErrorName() string {
	return "DeleteLimitRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteLimitRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteLimitRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteLimitRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteLimitRequestValidationError{}

// Validate checks the field values on LimitRuleInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LimitRuleInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LimitRuleInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LimitRuleInfoMultiError, or
// nil if none found.
func (m *LimitRuleInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *LimitRuleInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ConsumerId

	// no validation rules for PackageId

	// no validation rules for Method

	// no validation rules for ApiPath

	if all {
		switch v := interface{}(m.GetLimit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LimitRuleInfoValidationError{
					field:  "Limit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LimitRuleInfoValidationError{
					field:  "Limit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LimitRuleInfoValidationError{
				field:  "Limit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Id

	// no validation rules for CreateAt

	// no validation rules for ConsumerName

	// no validation rules for PackageName

	if len(errors) > 0 {
		return LimitRuleInfoMultiError(errors)
	}

	return nil
}

// LimitRuleInfoMultiError is an error wrapping multiple validation errors
// returned by LimitRuleInfo.ValidateAll() if the designated constraints
// aren't met.
type LimitRuleInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LimitRuleInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LimitRuleInfoMultiError) AllErrors() []error { return m }

// LimitRuleInfoValidationError is the validation error returned by
// LimitRuleInfo.Validate if the designated constraints aren't met.
type LimitRuleInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LimitRuleInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LimitRuleInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LimitRuleInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LimitRuleInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LimitRuleInfoValidationError) ErrorName() string { return "LimitRuleInfoValidationError" }

// Error satisfies the builtin error interface
func (e LimitRuleInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLimitRuleInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LimitRuleInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LimitRuleInfoValidationError{}

// Validate checks the field values on UpdateLimitResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateLimitResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateLimitResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateLimitResponseMultiError, or nil if none found.
func (m *UpdateLimitResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateLimitResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateLimitResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateLimitResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateLimitResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateLimitResponseMultiError(errors)
	}

	return nil
}

// UpdateLimitResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateLimitResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateLimitResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateLimitResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateLimitResponseMultiError) AllErrors() []error { return m }

// UpdateLimitResponseValidationError is the validation error returned by
// UpdateLimitResponse.Validate if the designated constraints aren't met.
type UpdateLimitResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateLimitResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateLimitResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateLimitResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateLimitResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateLimitResponseValidationError) ErrorName() string {
	return "UpdateLimitResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateLimitResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateLimitResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateLimitResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateLimitResponseValidationError{}

// Validate checks the field values on UpdateLimitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateLimitRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateLimitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateLimitRequestMultiError, or nil if none found.
func (m *UpdateLimitRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateLimitRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RuleId

	if all {
		switch v := interface{}(m.GetLimitRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateLimitRequestValidationError{
					field:  "LimitRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateLimitRequestValidationError{
					field:  "LimitRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLimitRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateLimitRequestValidationError{
				field:  "LimitRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateLimitRequestMultiError(errors)
	}

	return nil
}

// UpdateLimitRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateLimitRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateLimitRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateLimitRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateLimitRequestMultiError) AllErrors() []error { return m }

// UpdateLimitRequestValidationError is the validation error returned by
// UpdateLimitRequest.Validate if the designated constraints aren't met.
type UpdateLimitRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateLimitRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateLimitRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateLimitRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateLimitRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateLimitRequestValidationError) ErrorName() string {
	return "UpdateLimitRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateLimitRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateLimitRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateLimitRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateLimitRequestValidationError{}

// Validate checks the field values on CreateLimitResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateLimitResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLimitResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLimitResponseMultiError, or nil if none found.
func (m *CreateLimitResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLimitResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return CreateLimitResponseMultiError(errors)
	}

	return nil
}

// CreateLimitResponseMultiError is an error wrapping multiple validation
// errors returned by CreateLimitResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateLimitResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLimitResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLimitResponseMultiError) AllErrors() []error { return m }

// CreateLimitResponseValidationError is the validation error returned by
// CreateLimitResponse.Validate if the designated constraints aren't met.
type CreateLimitResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLimitResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLimitResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLimitResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLimitResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLimitResponseValidationError) ErrorName() string {
	return "CreateLimitResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateLimitResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLimitResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLimitResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLimitResponseValidationError{}

// Validate checks the field values on LimitType with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LimitType) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LimitType with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LimitTypeMultiError, or nil
// if none found.
func (m *LimitType) ValidateAll() error {
	return m.validate(true)
}

func (m *LimitType) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Qpd

	// no validation rules for Qph

	// no validation rules for Qpm

	// no validation rules for Qps

	if len(errors) > 0 {
		return LimitTypeMultiError(errors)
	}

	return nil
}

// LimitTypeMultiError is an error wrapping multiple validation errors returned
// by LimitType.ValidateAll() if the designated constraints aren't met.
type LimitTypeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LimitTypeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LimitTypeMultiError) AllErrors() []error { return m }

// LimitTypeValidationError is the validation error returned by
// LimitType.Validate if the designated constraints aren't met.
type LimitTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LimitTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LimitTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LimitTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LimitTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LimitTypeValidationError) ErrorName() string { return "LimitTypeValidationError" }

// Error satisfies the builtin error interface
func (e LimitTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLimitType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LimitTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LimitTypeValidationError{}

// Validate checks the field values on LimitRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LimitRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LimitRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LimitRequestMultiError, or
// nil if none found.
func (m *LimitRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LimitRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ConsumerId

	// no validation rules for PackageId

	// no validation rules for Method

	// no validation rules for ApiPath

	if all {
		switch v := interface{}(m.GetLimit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LimitRequestValidationError{
					field:  "Limit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LimitRequestValidationError{
					field:  "Limit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LimitRequestValidationError{
				field:  "Limit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LimitRequestMultiError(errors)
	}

	return nil
}

// LimitRequestMultiError is an error wrapping multiple validation errors
// returned by LimitRequest.ValidateAll() if the designated constraints aren't met.
type LimitRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LimitRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LimitRequestMultiError) AllErrors() []error { return m }

// LimitRequestValidationError is the validation error returned by
// LimitRequest.Validate if the designated constraints aren't met.
type LimitRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LimitRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LimitRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LimitRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LimitRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LimitRequestValidationError) ErrorName() string { return "LimitRequestValidationError" }

// Error satisfies the builtin error interface
func (e LimitRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLimitRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LimitRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LimitRequestValidationError{}

// Validate checks the field values on CreateLimitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateLimitRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLimitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLimitRequestMultiError, or nil if none found.
func (m *CreateLimitRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLimitRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectId

	// no validation rules for Env

	if all {
		switch v := interface{}(m.GetLimitRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateLimitRequestValidationError{
					field:  "LimitRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateLimitRequestValidationError{
					field:  "LimitRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLimitRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateLimitRequestValidationError{
				field:  "LimitRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateLimitRequestMultiError(errors)
	}

	return nil
}

// CreateLimitRequestMultiError is an error wrapping multiple validation errors
// returned by CreateLimitRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateLimitRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLimitRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLimitRequestMultiError) AllErrors() []error { return m }

// CreateLimitRequestValidationError is the validation error returned by
// CreateLimitRequest.Validate if the designated constraints aren't met.
type CreateLimitRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLimitRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLimitRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLimitRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLimitRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLimitRequestValidationError) ErrorName() string {
	return "CreateLimitRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateLimitRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLimitRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLimitRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLimitRequestValidationError{}

// Validate checks the field values on GetLimitsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetLimitsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLimitsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLimitsRequestMultiError, or nil if none found.
func (m *GetLimitsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLimitsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ConsumerId

	// no validation rules for PackageId

	// no validation rules for PageNo

	// no validation rules for PageSize

	if len(errors) > 0 {
		return GetLimitsRequestMultiError(errors)
	}

	return nil
}

// GetLimitsRequestMultiError is an error wrapping multiple validation errors
// returned by GetLimitsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetLimitsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLimitsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLimitsRequestMultiError) AllErrors() []error { return m }

// GetLimitsRequestValidationError is the validation error returned by
// GetLimitsRequest.Validate if the designated constraints aren't met.
type GetLimitsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLimitsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLimitsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLimitsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLimitsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLimitsRequestValidationError) ErrorName() string { return "GetLimitsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetLimitsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLimitsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLimitsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLimitsRequestValidationError{}

// Validate checks the field values on GetLimitsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetLimitsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLimitsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLimitsResponseMultiError, or nil if none found.
func (m *GetLimitsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLimitsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetLimitsResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetLimitsResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetLimitsResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetLimitsResponseMultiError(errors)
	}

	return nil
}

// GetLimitsResponseMultiError is an error wrapping multiple validation errors
// returned by GetLimitsResponse.ValidateAll() if the designated constraints
// aren't met.
type GetLimitsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLimitsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLimitsResponseMultiError) AllErrors() []error { return m }

// GetLimitsResponseValidationError is the validation error returned by
// GetLimitsResponse.Validate if the designated constraints aren't met.
type GetLimitsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLimitsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLimitsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLimitsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLimitsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLimitsResponseValidationError) ErrorName() string {
	return "GetLimitsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetLimitsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLimitsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLimitsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLimitsResponseValidationError{}
