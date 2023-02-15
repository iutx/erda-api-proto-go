// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: dop_config_manage.proto

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

// Validate checks the field values on CONFIG_MANAGE_CONFIG_DEL_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CONFIG_MANAGE_CONFIG_DEL_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CONFIG_MANAGE_CONFIG_DEL_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CONFIG_MANAGE_CONFIG_DEL_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_CONFIG_DEL_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_CONFIG_DEL_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_CONFIG_DEL_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_CONFIG_DEL_RequestMultiError is an error wrapping multiple
// validation errors returned by
// CONFIG_MANAGE_CONFIG_DEL_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_DEL_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_CONFIG_DEL_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_CONFIG_DEL_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_CONFIG_DEL_RequestValidationError is the validation error
// returned by CONFIG_MANAGE_CONFIG_DEL_Request.Validate if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_DEL_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_CONFIG_DEL_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_CONFIG_DEL_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_CONFIG_DEL_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_CONFIG_DEL_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_CONFIG_DEL_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_CONFIG_DEL_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_CONFIG_DEL_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_CONFIG_DEL_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_CONFIG_DEL_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_CONFIG_DEL_RequestValidationError{}

// Validate checks the field values on CONFIG_MANAGE_CONFIG_EXPORT_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CONFIG_MANAGE_CONFIG_EXPORT_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CONFIG_MANAGE_CONFIG_EXPORT_Request
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CONFIG_MANAGE_CONFIG_EXPORT_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_CONFIG_EXPORT_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_CONFIG_EXPORT_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_CONFIG_EXPORT_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_CONFIG_EXPORT_RequestMultiError is an error wrapping multiple
// validation errors returned by
// CONFIG_MANAGE_CONFIG_EXPORT_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_EXPORT_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_CONFIG_EXPORT_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_CONFIG_EXPORT_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError is the validation error
// returned by CONFIG_MANAGE_CONFIG_EXPORT_Request.Validate if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_CONFIG_EXPORT_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_CONFIG_EXPORT_RequestValidationError{}

// Validate checks the field values on CONFIG_MANAGE_CONFIG_GET_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CONFIG_MANAGE_CONFIG_GET_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CONFIG_MANAGE_CONFIG_GET_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CONFIG_MANAGE_CONFIG_GET_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_CONFIG_GET_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_CONFIG_GET_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_CONFIG_GET_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_CONFIG_GET_RequestMultiError is an error wrapping multiple
// validation errors returned by
// CONFIG_MANAGE_CONFIG_GET_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_GET_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_CONFIG_GET_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_CONFIG_GET_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_CONFIG_GET_RequestValidationError is the validation error
// returned by CONFIG_MANAGE_CONFIG_GET_Request.Validate if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_GET_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_CONFIG_GET_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_CONFIG_GET_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_CONFIG_GET_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_CONFIG_GET_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_CONFIG_GET_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_CONFIG_GET_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_CONFIG_GET_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_CONFIG_GET_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_CONFIG_GET_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_CONFIG_GET_RequestValidationError{}

// Validate checks the field values on CONFIG_MANAGE_CONFIG_IMPORT_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CONFIG_MANAGE_CONFIG_IMPORT_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CONFIG_MANAGE_CONFIG_IMPORT_Request
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CONFIG_MANAGE_CONFIG_IMPORT_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_CONFIG_IMPORT_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_CONFIG_IMPORT_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_CONFIG_IMPORT_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_CONFIG_IMPORT_RequestMultiError is an error wrapping multiple
// validation errors returned by
// CONFIG_MANAGE_CONFIG_IMPORT_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_IMPORT_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_CONFIG_IMPORT_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_CONFIG_IMPORT_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError is the validation error
// returned by CONFIG_MANAGE_CONFIG_IMPORT_Request.Validate if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_CONFIG_IMPORT_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_CONFIG_IMPORT_RequestValidationError{}

// Validate checks the field values on CONFIG_MANAGE_CONFIG_POST_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CONFIG_MANAGE_CONFIG_POST_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CONFIG_MANAGE_CONFIG_POST_Request
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CONFIG_MANAGE_CONFIG_POST_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_CONFIG_POST_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_CONFIG_POST_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_CONFIG_POST_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_CONFIG_POST_RequestMultiError is an error wrapping multiple
// validation errors returned by
// CONFIG_MANAGE_CONFIG_POST_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_POST_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_CONFIG_POST_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_CONFIG_POST_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_CONFIG_POST_RequestValidationError is the validation error
// returned by CONFIG_MANAGE_CONFIG_POST_Request.Validate if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_POST_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_CONFIG_POST_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_CONFIG_POST_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_CONFIG_POST_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_CONFIG_POST_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_CONFIG_POST_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_CONFIG_POST_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_CONFIG_POST_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_CONFIG_POST_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_CONFIG_POST_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_CONFIG_POST_RequestValidationError{}

// Validate checks the field values on CONFIG_MANAGE_CONFIG_PUT_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CONFIG_MANAGE_CONFIG_PUT_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CONFIG_MANAGE_CONFIG_PUT_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CONFIG_MANAGE_CONFIG_PUT_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_CONFIG_PUT_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_CONFIG_PUT_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_CONFIG_PUT_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_CONFIG_PUT_RequestMultiError is an error wrapping multiple
// validation errors returned by
// CONFIG_MANAGE_CONFIG_PUT_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_PUT_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_CONFIG_PUT_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_CONFIG_PUT_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_CONFIG_PUT_RequestValidationError is the validation error
// returned by CONFIG_MANAGE_CONFIG_PUT_Request.Validate if the designated
// constraints aren't met.
type CONFIG_MANAGE_CONFIG_PUT_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_CONFIG_PUT_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_CONFIG_PUT_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_CONFIG_PUT_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_CONFIG_PUT_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_CONFIG_PUT_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_CONFIG_PUT_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_CONFIG_PUT_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_CONFIG_PUT_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_CONFIG_PUT_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_CONFIG_PUT_RequestValidationError{}

// Validate checks the field values on CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestMultiError is an error wrapping
// multiple validation errors returned by
// CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError is the validation
// error returned by CONFIG_MANAGE_DEPLOY_CONFIG_GET_Request.Validate if the
// designated constraints aren't met.
type CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_DEPLOY_CONFIG_GET_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_DEPLOY_CONFIG_GET_RequestValidationError{}

// Validate checks the field values on
// CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestMultiError is an error wrapping
// multiple validation errors returned by
// CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError is the validation
// error returned by CONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request.Validate if the
// designated constraints aren't met.
type CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_MULTI_NS_CONFIG_GET_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_MULTI_NS_CONFIG_GET_RequestValidationError{}

// Validate checks the field values on CONFIG_MANAGE_NAMESPACE_DEL_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CONFIG_MANAGE_NAMESPACE_DEL_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CONFIG_MANAGE_NAMESPACE_DEL_Request
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CONFIG_MANAGE_NAMESPACE_DEL_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_NAMESPACE_DEL_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_NAMESPACE_DEL_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_NAMESPACE_DEL_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_NAMESPACE_DEL_RequestMultiError is an error wrapping multiple
// validation errors returned by
// CONFIG_MANAGE_NAMESPACE_DEL_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_NAMESPACE_DEL_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_NAMESPACE_DEL_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_NAMESPACE_DEL_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError is the validation error
// returned by CONFIG_MANAGE_NAMESPACE_DEL_Request.Validate if the designated
// constraints aren't met.
type CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_NAMESPACE_DEL_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_NAMESPACE_DEL_RequestValidationError{}

// Validate checks the field values on CONFIG_MANAGE_NAMESPACE_FIX_Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CONFIG_MANAGE_NAMESPACE_FIX_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CONFIG_MANAGE_NAMESPACE_FIX_Request
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CONFIG_MANAGE_NAMESPACE_FIX_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_NAMESPACE_FIX_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_NAMESPACE_FIX_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_NAMESPACE_FIX_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_NAMESPACE_FIX_RequestMultiError is an error wrapping multiple
// validation errors returned by
// CONFIG_MANAGE_NAMESPACE_FIX_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_NAMESPACE_FIX_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_NAMESPACE_FIX_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_NAMESPACE_FIX_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError is the validation error
// returned by CONFIG_MANAGE_NAMESPACE_FIX_Request.Validate if the designated
// constraints aren't met.
type CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_NAMESPACE_FIX_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_NAMESPACE_FIX_RequestValidationError{}

// Validate checks the field values on CONFIG_MANAGE_NAMESPACE_POST_Request
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CONFIG_MANAGE_NAMESPACE_POST_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CONFIG_MANAGE_NAMESPACE_POST_Request
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CONFIG_MANAGE_NAMESPACE_POST_RequestMultiError, or nil if none found.
func (m *CONFIG_MANAGE_NAMESPACE_POST_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *CONFIG_MANAGE_NAMESPACE_POST_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CONFIG_MANAGE_NAMESPACE_POST_RequestMultiError(errors)
	}

	return nil
}

// CONFIG_MANAGE_NAMESPACE_POST_RequestMultiError is an error wrapping multiple
// validation errors returned by
// CONFIG_MANAGE_NAMESPACE_POST_Request.ValidateAll() if the designated
// constraints aren't met.
type CONFIG_MANAGE_NAMESPACE_POST_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CONFIG_MANAGE_NAMESPACE_POST_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CONFIG_MANAGE_NAMESPACE_POST_RequestMultiError) AllErrors() []error { return m }

// CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError is the validation error
// returned by CONFIG_MANAGE_NAMESPACE_POST_Request.Validate if the designated
// constraints aren't met.
type CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError) ErrorName() string {
	return "CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCONFIG_MANAGE_NAMESPACE_POST_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CONFIG_MANAGE_NAMESPACE_POST_RequestValidationError{}
