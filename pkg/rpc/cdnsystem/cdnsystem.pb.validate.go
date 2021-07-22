// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/rpc/cdnsystem/cdnsystem.proto

package cdnsystem

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

// Validate checks the field values on SeedRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SeedRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TaskId

	// no validation rules for Url

	// no validation rules for Filter

	if v, ok := interface{}(m.GetUrlMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SeedRequestValidationError{
				field:  "UrlMeta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SeedRequestValidationError is the validation error returned by
// SeedRequest.Validate if the designated constraints aren't met.
type SeedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SeedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SeedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SeedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SeedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SeedRequestValidationError) ErrorName() string { return "SeedRequestValidationError" }

// Error satisfies the builtin error interface
func (e SeedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSeedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SeedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SeedRequestValidationError{}

// Validate checks the field values on PieceSeed with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PieceSeed) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PeerId

	// no validation rules for HostUuid

	if v, ok := interface{}(m.GetPieceInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PieceSeedValidationError{
				field:  "PieceInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Done

	// no validation rules for ContentLength

	// no validation rules for TotalPieceCount

	return nil
}

// PieceSeedValidationError is the validation error returned by
// PieceSeed.Validate if the designated constraints aren't met.
type PieceSeedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PieceSeedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PieceSeedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PieceSeedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PieceSeedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PieceSeedValidationError) ErrorName() string { return "PieceSeedValidationError" }

// Error satisfies the builtin error interface
func (e PieceSeedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPieceSeed.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PieceSeedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PieceSeedValidationError{}