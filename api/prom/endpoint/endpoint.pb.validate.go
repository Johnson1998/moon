// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prom/endpoint/endpoint.proto

package endpoint

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

// define the regex for a UUID once up-front
var _endpoint_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AppendEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AppendEndpointRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppendEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppendEndpointRequestMultiError, or nil if none found.
func (m *AppendEndpointRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AppendEndpointRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAgentName()); l < 2 || l > 32 {
		err := AppendEndpointRequestValidationError{
			field:  "AgentName",
			reason: "value length must be between 2 and 32 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetEndpoints()) < 1 {
		err := AppendEndpointRequestValidationError{
			field:  "Endpoints",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		if item == nil {
			err := AppendEndpointRequestValidationError{
				field:  fmt.Sprintf("Endpoints[%v]", idx),
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AppendEndpointRequestValidationError{
						field:  fmt.Sprintf("Endpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AppendEndpointRequestValidationError{
						field:  fmt.Sprintf("Endpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AppendEndpointRequestValidationError{
					field:  fmt.Sprintf("Endpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AppendEndpointRequestMultiError(errors)
	}

	return nil
}

// AppendEndpointRequestMultiError is an error wrapping multiple validation
// errors returned by AppendEndpointRequest.ValidateAll() if the designated
// constraints aren't met.
type AppendEndpointRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppendEndpointRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppendEndpointRequestMultiError) AllErrors() []error { return m }

// AppendEndpointRequestValidationError is the validation error returned by
// AppendEndpointRequest.Validate if the designated constraints aren't met.
type AppendEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppendEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppendEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppendEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppendEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppendEndpointRequestValidationError) ErrorName() string {
	return "AppendEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AppendEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppendEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppendEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppendEndpointRequestValidationError{}

// Validate checks the field values on AppendEndpointReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AppendEndpointReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppendEndpointReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppendEndpointReplyMultiError, or nil if none found.
func (m *AppendEndpointReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AppendEndpointReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AppendEndpointReplyValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AppendEndpointReplyValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AppendEndpointReplyValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AppendEndpointReplyMultiError(errors)
	}

	return nil
}

// AppendEndpointReplyMultiError is an error wrapping multiple validation
// errors returned by AppendEndpointReply.ValidateAll() if the designated
// constraints aren't met.
type AppendEndpointReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppendEndpointReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppendEndpointReplyMultiError) AllErrors() []error { return m }

// AppendEndpointReplyValidationError is the validation error returned by
// AppendEndpointReply.Validate if the designated constraints aren't met.
type AppendEndpointReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppendEndpointReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppendEndpointReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppendEndpointReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppendEndpointReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppendEndpointReplyValidationError) ErrorName() string {
	return "AppendEndpointReplyValidationError"
}

// Error satisfies the builtin error interface
func (e AppendEndpointReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppendEndpointReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppendEndpointReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppendEndpointReplyValidationError{}

// Validate checks the field values on DeleteEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteEndpointRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteEndpointRequestMultiError, or nil if none found.
func (m *DeleteEndpointRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteEndpointRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAgentName()); l < 2 || l > 32 {
		err := DeleteEndpointRequestValidationError{
			field:  "AgentName",
			reason: "value length must be between 2 and 32 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetUuids()) < 1 {
		err := DeleteEndpointRequestValidationError{
			field:  "Uuids",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetUuids() {
		_, _ = idx, item

		if err := m._validateUuid(item); err != nil {
			err = DeleteEndpointRequestValidationError{
				field:  fmt.Sprintf("Uuids[%v]", idx),
				reason: "value must be a valid UUID",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return DeleteEndpointRequestMultiError(errors)
	}

	return nil
}

func (m *DeleteEndpointRequest) _validateUuid(uuid string) error {
	if matched := _endpoint_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteEndpointRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteEndpointRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteEndpointRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteEndpointRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteEndpointRequestMultiError) AllErrors() []error { return m }

// DeleteEndpointRequestValidationError is the validation error returned by
// DeleteEndpointRequest.Validate if the designated constraints aren't met.
type DeleteEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEndpointRequestValidationError) ErrorName() string {
	return "DeleteEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEndpointRequestValidationError{}

// Validate checks the field values on DeleteEndpointReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteEndpointReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteEndpointReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteEndpointReplyMultiError, or nil if none found.
func (m *DeleteEndpointReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteEndpointReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteEndpointReplyValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteEndpointReplyValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteEndpointReplyValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeleteEndpointReplyMultiError(errors)
	}

	return nil
}

// DeleteEndpointReplyMultiError is an error wrapping multiple validation
// errors returned by DeleteEndpointReply.ValidateAll() if the designated
// constraints aren't met.
type DeleteEndpointReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteEndpointReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteEndpointReplyMultiError) AllErrors() []error { return m }

// DeleteEndpointReplyValidationError is the validation error returned by
// DeleteEndpointReply.Validate if the designated constraints aren't met.
type DeleteEndpointReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEndpointReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEndpointReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEndpointReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEndpointReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEndpointReplyValidationError) ErrorName() string {
	return "DeleteEndpointReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEndpointReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEndpointReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEndpointReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEndpointReplyValidationError{}

// Validate checks the field values on ListEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListEndpointRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListEndpointRequestMultiError, or nil if none found.
func (m *ListEndpointRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEndpointRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListEndpointRequestMultiError(errors)
	}

	return nil
}

// ListEndpointRequestMultiError is an error wrapping multiple validation
// errors returned by ListEndpointRequest.ValidateAll() if the designated
// constraints aren't met.
type ListEndpointRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEndpointRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEndpointRequestMultiError) AllErrors() []error { return m }

// ListEndpointRequestValidationError is the validation error returned by
// ListEndpointRequest.Validate if the designated constraints aren't met.
type ListEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEndpointRequestValidationError) ErrorName() string {
	return "ListEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEndpointRequestValidationError{}

// Validate checks the field values on ListEndpointReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListEndpointReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEndpointReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListEndpointReplyMultiError, or nil if none found.
func (m *ListEndpointReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEndpointReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEndpointReplyValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEndpointReplyValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEndpointReplyValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListEndpointReplyValidationError{
						field:  fmt.Sprintf("Endpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListEndpointReplyValidationError{
						field:  fmt.Sprintf("Endpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListEndpointReplyValidationError{
					field:  fmt.Sprintf("Endpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListEndpointReplyMultiError(errors)
	}

	return nil
}

// ListEndpointReplyMultiError is an error wrapping multiple validation errors
// returned by ListEndpointReply.ValidateAll() if the designated constraints
// aren't met.
type ListEndpointReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEndpointReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEndpointReplyMultiError) AllErrors() []error { return m }

// ListEndpointReplyValidationError is the validation error returned by
// ListEndpointReply.Validate if the designated constraints aren't met.
type ListEndpointReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEndpointReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEndpointReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEndpointReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEndpointReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEndpointReplyValidationError) ErrorName() string {
	return "ListEndpointReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListEndpointReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEndpointReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEndpointReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEndpointReplyValidationError{}