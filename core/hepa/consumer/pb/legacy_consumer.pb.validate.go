// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: legacy_consumer.proto

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

// Validate checks the field values on GetConsumerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetConsumerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConsumerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConsumerRequestMultiError, or nil if none found.
func (m *GetConsumerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConsumerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrgId

	// no validation rules for ProjectId

	// no validation rules for Env

	if len(errors) > 0 {
		return GetConsumerRequestMultiError(errors)
	}

	return nil
}

// GetConsumerRequestMultiError is an error wrapping multiple validation errors
// returned by GetConsumerRequest.ValidateAll() if the designated constraints
// aren't met.
type GetConsumerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConsumerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConsumerRequestMultiError) AllErrors() []error { return m }

// GetConsumerRequestValidationError is the validation error returned by
// GetConsumerRequest.Validate if the designated constraints aren't met.
type GetConsumerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConsumerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConsumerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConsumerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConsumerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConsumerRequestValidationError) ErrorName() string {
	return "GetConsumerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetConsumerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConsumerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConsumerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConsumerRequestValidationError{}

// Validate checks the field values on Endpoint with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Endpoint) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Endpoint with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EndpointMultiError, or nil
// if none found.
func (m *Endpoint) ValidateAll() error {
	return m.validate(true)
}

func (m *Endpoint) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OuterAddr

	// no validation rules for InnerAddr

	// no validation rules for InnerTips

	if len(errors) > 0 {
		return EndpointMultiError(errors)
	}

	return nil
}

// EndpointMultiError is an error wrapping multiple validation errors returned
// by Endpoint.ValidateAll() if the designated constraints aren't met.
type EndpointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EndpointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EndpointMultiError) AllErrors() []error { return m }

// EndpointValidationError is the validation error returned by
// Endpoint.Validate if the designated constraints aren't met.
type EndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointValidationError) ErrorName() string { return "EndpointValidationError" }

// Error satisfies the builtin error interface
func (e EndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointValidationError{}

// Validate checks the field values on GetConsumerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetConsumerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConsumerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConsumerResponseMultiError, or nil if none found.
func (m *GetConsumerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConsumerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ConsumerName

	// no validation rules for ConsumerId

	if all {
		switch v := interface{}(m.GetEndpoint()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetConsumerResponseValidationError{
					field:  "Endpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetConsumerResponseValidationError{
					field:  "Endpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetConsumerResponseValidationError{
				field:  "Endpoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for GatewayInstance

	// no validation rules for ClusterName

	if len(errors) > 0 {
		return GetConsumerResponseMultiError(errors)
	}

	return nil
}

// GetConsumerResponseMultiError is an error wrapping multiple validation
// errors returned by GetConsumerResponse.ValidateAll() if the designated
// constraints aren't met.
type GetConsumerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConsumerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConsumerResponseMultiError) AllErrors() []error { return m }

// GetConsumerResponseValidationError is the validation error returned by
// GetConsumerResponse.Validate if the designated constraints aren't met.
type GetConsumerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConsumerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConsumerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConsumerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConsumerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConsumerResponseValidationError) ErrorName() string {
	return "GetConsumerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetConsumerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConsumerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConsumerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConsumerResponseValidationError{}