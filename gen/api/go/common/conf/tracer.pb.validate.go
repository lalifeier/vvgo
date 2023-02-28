// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/conf/tracer.proto

package conf

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

// Validate checks the field values on Trace with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Trace) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Trace with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TraceMultiError, or nil if none found.
func (m *Trace) ValidateAll() error {
	return m.validate(true)
}

func (m *Trace) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Batcher

	// no validation rules for Endpoint

	// no validation rules for Sampler

	// no validation rules for Env

	if len(errors) > 0 {
		return TraceMultiError(errors)
	}

	return nil
}

// TraceMultiError is an error wrapping multiple validation errors returned by
// Trace.ValidateAll() if the designated constraints aren't met.
type TraceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TraceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TraceMultiError) AllErrors() []error { return m }

// TraceValidationError is the validation error returned by Trace.Validate if
// the designated constraints aren't met.
type TraceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TraceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TraceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TraceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TraceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TraceValidationError) ErrorName() string { return "TraceValidationError" }

// Error satisfies the builtin error interface
func (e TraceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TraceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TraceValidationError{}