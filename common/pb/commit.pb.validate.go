// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commit.proto

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

// Validate checks the field values on CommitDetail with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommitDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommitDetail with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommitDetailMultiError, or
// nil if none found.
func (m *CommitDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *CommitDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CommitID

	// no validation rules for Repo

	// no validation rules for RepoAbbr

	// no validation rules for Author

	// no validation rules for Email

	if all {
		switch v := interface{}(m.GetTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommitDetailValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommitDetailValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommitDetailValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Comment

	if len(errors) > 0 {
		return CommitDetailMultiError(errors)
	}

	return nil
}

// CommitDetailMultiError is an error wrapping multiple validation errors
// returned by CommitDetail.ValidateAll() if the designated constraints aren't met.
type CommitDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommitDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommitDetailMultiError) AllErrors() []error { return m }

// CommitDetailValidationError is the validation error returned by
// CommitDetail.Validate if the designated constraints aren't met.
type CommitDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommitDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommitDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommitDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommitDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommitDetailValidationError) ErrorName() string { return "CommitDetailValidationError" }

// Error satisfies the builtin error interface
func (e CommitDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommitDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommitDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommitDetailValidationError{}
