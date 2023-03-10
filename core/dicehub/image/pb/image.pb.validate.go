// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: image.proto

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

// Validate checks the field values on ImageGetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ImageGetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImageGetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImageGetRequestMultiError, or nil if none found.
func (m *ImageGetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ImageGetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ImageIDOrImage

	if len(errors) > 0 {
		return ImageGetRequestMultiError(errors)
	}

	return nil
}

// ImageGetRequestMultiError is an error wrapping multiple validation errors
// returned by ImageGetRequest.ValidateAll() if the designated constraints
// aren't met.
type ImageGetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageGetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageGetRequestMultiError) AllErrors() []error { return m }

// ImageGetRequestValidationError is the validation error returned by
// ImageGetRequest.Validate if the designated constraints aren't met.
type ImageGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageGetRequestValidationError) ErrorName() string { return "ImageGetRequestValidationError" }

// Error satisfies the builtin error interface
func (e ImageGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageGetRequestValidationError{}

// Validate checks the field values on ImageGetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ImageGetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImageGetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImageGetResponseMultiError, or nil if none found.
func (m *ImageGetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ImageGetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ImageGetResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImageGetResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageGetResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ImageGetResponseMultiError(errors)
	}

	return nil
}

// ImageGetResponseMultiError is an error wrapping multiple validation errors
// returned by ImageGetResponse.ValidateAll() if the designated constraints
// aren't met.
type ImageGetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageGetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageGetResponseMultiError) AllErrors() []error { return m }

// ImageGetResponseValidationError is the validation error returned by
// ImageGetResponse.Validate if the designated constraints aren't met.
type ImageGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageGetResponseValidationError) ErrorName() string { return "ImageGetResponseValidationError" }

// Error satisfies the builtin error interface
func (e ImageGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageGetResponseValidationError{}

// Validate checks the field values on ImageListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ImageListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImageListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImageListRequestMultiError, or nil if none found.
func (m *ImageListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ImageListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageSize

	// no validation rules for PageNum

	if len(errors) > 0 {
		return ImageListRequestMultiError(errors)
	}

	return nil
}

// ImageListRequestMultiError is an error wrapping multiple validation errors
// returned by ImageListRequest.ValidateAll() if the designated constraints
// aren't met.
type ImageListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageListRequestMultiError) AllErrors() []error { return m }

// ImageListRequestValidationError is the validation error returned by
// ImageListRequest.Validate if the designated constraints aren't met.
type ImageListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageListRequestValidationError) ErrorName() string { return "ImageListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ImageListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageListRequestValidationError{}

// Validate checks the field values on ImageListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ImageListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImageListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImageListResponseMultiError, or nil if none found.
func (m *ImageListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ImageListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ImageListResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImageListResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageListResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ImageListResponseMultiError(errors)
	}

	return nil
}

// ImageListResponseMultiError is an error wrapping multiple validation errors
// returned by ImageListResponse.ValidateAll() if the designated constraints
// aren't met.
type ImageListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageListResponseMultiError) AllErrors() []error { return m }

// ImageListResponseValidationError is the validation error returned by
// ImageListResponse.Validate if the designated constraints aren't met.
type ImageListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageListResponseValidationError) ErrorName() string {
	return "ImageListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ImageListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageListResponseValidationError{}

// Validate checks the field values on ImageListResponseData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ImageListResponseData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImageListResponseData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImageListResponseDataMultiError, or nil if none found.
func (m *ImageListResponseData) ValidateAll() error {
	return m.validate(true)
}

func (m *ImageListResponseData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ImageListResponseDataValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ImageListResponseDataValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ImageListResponseDataValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ImageListResponseDataMultiError(errors)
	}

	return nil
}

// ImageListResponseDataMultiError is an error wrapping multiple validation
// errors returned by ImageListResponseData.ValidateAll() if the designated
// constraints aren't met.
type ImageListResponseDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageListResponseDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageListResponseDataMultiError) AllErrors() []error { return m }

// ImageListResponseDataValidationError is the validation error returned by
// ImageListResponseData.Validate if the designated constraints aren't met.
type ImageListResponseDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageListResponseDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageListResponseDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageListResponseDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageListResponseDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageListResponseDataValidationError) ErrorName() string {
	return "ImageListResponseDataValidationError"
}

// Error satisfies the builtin error interface
func (e ImageListResponseDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageListResponseData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageListResponseDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageListResponseDataValidationError{}

// Validate checks the field values on ImageGetResponseData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ImageGetResponseData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImageGetResponseData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImageGetResponseDataMultiError, or nil if none found.
func (m *ImageGetResponseData) ValidateAll() error {
	return m.validate(true)
}

func (m *ImageGetResponseData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ID

	// no validation rules for ReleaseID

	// no validation rules for ImageName

	// no validation rules for ImageTag

	// no validation rules for Image

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ImageGetResponseDataValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImageGetResponseDataValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageGetResponseDataValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ImageGetResponseDataValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImageGetResponseDataValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageGetResponseDataValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ImageGetResponseDataMultiError(errors)
	}

	return nil
}

// ImageGetResponseDataMultiError is an error wrapping multiple validation
// errors returned by ImageGetResponseData.ValidateAll() if the designated
// constraints aren't met.
type ImageGetResponseDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageGetResponseDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageGetResponseDataMultiError) AllErrors() []error { return m }

// ImageGetResponseDataValidationError is the validation error returned by
// ImageGetResponseData.Validate if the designated constraints aren't met.
type ImageGetResponseDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageGetResponseDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageGetResponseDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageGetResponseDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageGetResponseDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageGetResponseDataValidationError) ErrorName() string {
	return "ImageGetResponseDataValidationError"
}

// Error satisfies the builtin error interface
func (e ImageGetResponseDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageGetResponseData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageGetResponseDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageGetResponseDataValidationError{}
