// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ozonmp/trv_airport_api/v1/trv_airport_api.proto

package trv_airport_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Airport with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Airport) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Location

	return nil
}

// AirportValidationError is the validation error returned by Airport.Validate
// if the designated constraints aren't met.
type AirportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AirportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AirportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AirportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AirportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AirportValidationError) ErrorName() string { return "AirportValidationError" }

// Error satisfies the builtin error interface
func (e AirportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAirport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AirportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AirportValidationError{}

// Validate checks the field values on CreateAirportV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAirportV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) != 3 {
		return CreateAirportV1RequestValidationError{
			field:  "Name",
			reason: "value length must be 3 runes",
		}

	}

	if l := utf8.RuneCountInString(m.GetLocation()); l < 2 || l > 255 {
		return CreateAirportV1RequestValidationError{
			field:  "Location",
			reason: "value length must be between 2 and 255 runes, inclusive",
		}
	}

	return nil
}

// CreateAirportV1RequestValidationError is the validation error returned by
// CreateAirportV1Request.Validate if the designated constraints aren't met.
type CreateAirportV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAirportV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAirportV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAirportV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAirportV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAirportV1RequestValidationError) ErrorName() string {
	return "CreateAirportV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAirportV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAirportV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAirportV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAirportV1RequestValidationError{}

// Validate checks the field values on CreateAirportV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAirportV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateAirportV1ResponseValidationError is the validation error returned by
// CreateAirportV1Response.Validate if the designated constraints aren't met.
type CreateAirportV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAirportV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAirportV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAirportV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAirportV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAirportV1ResponseValidationError) ErrorName() string {
	return "CreateAirportV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAirportV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAirportV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAirportV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAirportV1ResponseValidationError{}

// Validate checks the field values on DescribeAirportV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeAirportV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAirportId() <= 0 {
		return DescribeAirportV1RequestValidationError{
			field:  "AirportId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeAirportV1RequestValidationError is the validation error returned by
// DescribeAirportV1Request.Validate if the designated constraints aren't met.
type DescribeAirportV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeAirportV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeAirportV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeAirportV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeAirportV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeAirportV1RequestValidationError) ErrorName() string {
	return "DescribeAirportV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeAirportV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeAirportV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeAirportV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeAirportV1RequestValidationError{}

// Validate checks the field values on DescribeAirportV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeAirportV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeAirportV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeAirportV1ResponseValidationError is the validation error returned by
// DescribeAirportV1Response.Validate if the designated constraints aren't met.
type DescribeAirportV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeAirportV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeAirportV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeAirportV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeAirportV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeAirportV1ResponseValidationError) ErrorName() string {
	return "DescribeAirportV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeAirportV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeAirportV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeAirportV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeAirportV1ResponseValidationError{}

// Validate checks the field values on ListAirportsV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListAirportsV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Cursor

	// no validation rules for Limit

	return nil
}

// ListAirportsV1RequestValidationError is the validation error returned by
// ListAirportsV1Request.Validate if the designated constraints aren't met.
type ListAirportsV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAirportsV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAirportsV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAirportsV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAirportsV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAirportsV1RequestValidationError) ErrorName() string {
	return "ListAirportsV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAirportsV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAirportsV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAirportsV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAirportsV1RequestValidationError{}

// Validate checks the field values on ListAirportsV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListAirportsV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetAirports() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAirportsV1ResponseValidationError{
					field:  fmt.Sprintf("Airports[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListAirportsV1ResponseValidationError is the validation error returned by
// ListAirportsV1Response.Validate if the designated constraints aren't met.
type ListAirportsV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAirportsV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAirportsV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAirportsV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAirportsV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAirportsV1ResponseValidationError) ErrorName() string {
	return "ListAirportsV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAirportsV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAirportsV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAirportsV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAirportsV1ResponseValidationError{}

// Validate checks the field values on DeleteAirportV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteAirportV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAirportId() <= 0 {
		return DeleteAirportV1RequestValidationError{
			field:  "AirportId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DeleteAirportV1RequestValidationError is the validation error returned by
// DeleteAirportV1Request.Validate if the designated constraints aren't met.
type DeleteAirportV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAirportV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAirportV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAirportV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAirportV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAirportV1RequestValidationError) ErrorName() string {
	return "DeleteAirportV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAirportV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAirportV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAirportV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAirportV1RequestValidationError{}

// Validate checks the field values on DeleteAirportV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteAirportV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AirportId

	return nil
}

// DeleteAirportV1ResponseValidationError is the validation error returned by
// DeleteAirportV1Response.Validate if the designated constraints aren't met.
type DeleteAirportV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAirportV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAirportV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAirportV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAirportV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAirportV1ResponseValidationError) ErrorName() string {
	return "DeleteAirportV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAirportV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAirportV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAirportV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAirportV1ResponseValidationError{}
