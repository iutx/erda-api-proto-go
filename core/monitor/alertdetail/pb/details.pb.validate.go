// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: details.proto

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

// Validate checks the field values on QuerySystemPodMetricsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuerySystemPodMetricsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuerySystemPodMetricsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuerySystemPodMetricsRequestMultiError, or nil if none found.
func (m *QuerySystemPodMetricsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuerySystemPodMetricsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := QuerySystemPodMetricsRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetClusterName()) < 1 {
		err := QuerySystemPodMetricsRequestValidationError{
			field:  "ClusterName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTimestamp() <= 1800000 {
		err := QuerySystemPodMetricsRequestValidationError{
			field:  "Timestamp",
			reason: "value must be greater than 1800000",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return QuerySystemPodMetricsRequestMultiError(errors)
	}

	return nil
}

// QuerySystemPodMetricsRequestMultiError is an error wrapping multiple
// validation errors returned by QuerySystemPodMetricsRequest.ValidateAll() if
// the designated constraints aren't met.
type QuerySystemPodMetricsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuerySystemPodMetricsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuerySystemPodMetricsRequestMultiError) AllErrors() []error { return m }

// QuerySystemPodMetricsRequestValidationError is the validation error returned
// by QuerySystemPodMetricsRequest.Validate if the designated constraints
// aren't met.
type QuerySystemPodMetricsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuerySystemPodMetricsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuerySystemPodMetricsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuerySystemPodMetricsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuerySystemPodMetricsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuerySystemPodMetricsRequestValidationError) ErrorName() string {
	return "QuerySystemPodMetricsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QuerySystemPodMetricsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuerySystemPodMetricsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuerySystemPodMetricsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuerySystemPodMetricsRequestValidationError{}

// Validate checks the field values on QuerySystemPodMetricsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuerySystemPodMetricsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuerySystemPodMetricsResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// QuerySystemPodMetricsResponseMultiError, or nil if none found.
func (m *QuerySystemPodMetricsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QuerySystemPodMetricsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuerySystemPodMetricsResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuerySystemPodMetricsResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuerySystemPodMetricsResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return QuerySystemPodMetricsResponseMultiError(errors)
	}

	return nil
}

// QuerySystemPodMetricsResponseMultiError is an error wrapping multiple
// validation errors returned by QuerySystemPodMetricsResponse.ValidateAll()
// if the designated constraints aren't met.
type QuerySystemPodMetricsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuerySystemPodMetricsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuerySystemPodMetricsResponseMultiError) AllErrors() []error { return m }

// QuerySystemPodMetricsResponseValidationError is the validation error
// returned by QuerySystemPodMetricsResponse.Validate if the designated
// constraints aren't met.
type QuerySystemPodMetricsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuerySystemPodMetricsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuerySystemPodMetricsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuerySystemPodMetricsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuerySystemPodMetricsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuerySystemPodMetricsResponseValidationError) ErrorName() string {
	return "QuerySystemPodMetricsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e QuerySystemPodMetricsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuerySystemPodMetricsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuerySystemPodMetricsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuerySystemPodMetricsResponseValidationError{}

// Validate checks the field values on PodInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PodInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PodInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PodInfoMultiError, or nil if none found.
func (m *PodInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *PodInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSummary()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PodInfoValidationError{
					field:  "Summary",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PodInfoValidationError{
					field:  "Summary",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSummary()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PodInfoValidationError{
				field:  "Summary",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetInstances() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PodInfoValidationError{
						field:  fmt.Sprintf("Instances[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PodInfoValidationError{
						field:  fmt.Sprintf("Instances[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PodInfoValidationError{
					field:  fmt.Sprintf("Instances[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PodInfoMultiError(errors)
	}

	return nil
}

// PodInfoMultiError is an error wrapping multiple validation errors returned
// by PodInfo.ValidateAll() if the designated constraints aren't met.
type PodInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PodInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PodInfoMultiError) AllErrors() []error { return m }

// PodInfoValidationError is the validation error returned by PodInfo.Validate
// if the designated constraints aren't met.
type PodInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PodInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PodInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PodInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PodInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PodInfoValidationError) ErrorName() string { return "PodInfoValidationError" }

// Error satisfies the builtin error interface
func (e PodInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPodInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PodInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PodInfoValidationError{}

// Validate checks the field values on PodInfoSummary with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PodInfoSummary) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PodInfoSummary with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PodInfoSummaryMultiError,
// or nil if none found.
func (m *PodInfoSummary) ValidateAll() error {
	return m.validate(true)
}

func (m *PodInfoSummary) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClusterName

	// no validation rules for NodeName

	// no validation rules for HostIP

	// no validation rules for Namespace

	// no validation rules for PodName

	// no validation rules for RestartTotal

	// no validation rules for StateCode

	// no validation rules for TerminatedReason

	if len(errors) > 0 {
		return PodInfoSummaryMultiError(errors)
	}

	return nil
}

// PodInfoSummaryMultiError is an error wrapping multiple validation errors
// returned by PodInfoSummary.ValidateAll() if the designated constraints
// aren't met.
type PodInfoSummaryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PodInfoSummaryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PodInfoSummaryMultiError) AllErrors() []error { return m }

// PodInfoSummaryValidationError is the validation error returned by
// PodInfoSummary.Validate if the designated constraints aren't met.
type PodInfoSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PodInfoSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PodInfoSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PodInfoSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PodInfoSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PodInfoSummaryValidationError) ErrorName() string { return "PodInfoSummaryValidationError" }

// Error satisfies the builtin error interface
func (e PodInfoSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPodInfoSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PodInfoSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PodInfoSummaryValidationError{}

// Validate checks the field values on PodInfoInstanse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PodInfoInstanse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PodInfoInstanse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PodInfoInstanseMultiError, or nil if none found.
func (m *PodInfoInstanse) ValidateAll() error {
	return m.validate(true)
}

func (m *PodInfoInstanse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ContainerId

	// no validation rules for HostIP

	// no validation rules for StartedAt

	// no validation rules for FinishedAt

	// no validation rules for ExitCode

	// no validation rules for OomKilled

	if len(errors) > 0 {
		return PodInfoInstanseMultiError(errors)
	}

	return nil
}

// PodInfoInstanseMultiError is an error wrapping multiple validation errors
// returned by PodInfoInstanse.ValidateAll() if the designated constraints
// aren't met.
type PodInfoInstanseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PodInfoInstanseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PodInfoInstanseMultiError) AllErrors() []error { return m }

// PodInfoInstanseValidationError is the validation error returned by
// PodInfoInstanse.Validate if the designated constraints aren't met.
type PodInfoInstanseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PodInfoInstanseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PodInfoInstanseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PodInfoInstanseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PodInfoInstanseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PodInfoInstanseValidationError) ErrorName() string { return "PodInfoInstanseValidationError" }

// Error satisfies the builtin error interface
func (e PodInfoInstanseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPodInfoInstanse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PodInfoInstanseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PodInfoInstanseValidationError{}