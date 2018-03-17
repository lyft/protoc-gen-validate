// Code generated by protoc-gen-validate
// source: wrapper.proto
// DO NOT EDIT!!!

package tests_kitchensink

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

	"github.com/gogo/protobuf/types"
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

	_ = types.DynamicAny{}
)

// Validate checks the field values on Wrappers with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Wrappers) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNone()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WrappersValidationError{
				Field:  "None",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if wrapper := m.GetDoubleValue(); wrapper != nil {

		if wrapper.GetValue() <= 1.23 {
			return WrappersValidationError{
				Field:  "DoubleValue",
				Reason: "value must be greater than 1.23",
			}
		}

	}

	if wrapper := m.GetFloatValue(); wrapper != nil {

		if wrapper.GetValue() > 4.56 {
			return WrappersValidationError{
				Field:  "FloatValue",
				Reason: "value must be less than or equal to 4.56",
			}
		}

	}

	if wrapper := m.GetInt64Value(); wrapper != nil {

		if wrapper.GetValue() != 5 {
			return WrappersValidationError{
				Field:  "Int64Value",
				Reason: "value must equal 5",
			}
		}

	}

	if wrapper := m.GetUint64Value(); wrapper != nil {

		if _, ok := _Wrappers_Uint64Value_InLookup[wrapper.GetValue()]; !ok {
			return WrappersValidationError{
				Field:  "Uint64Value",
				Reason: "value must be in list [1 2 3]",
			}
		}

	}

	if wrapper := m.GetInt32Value(); wrapper != nil {

		if _, ok := _Wrappers_Int32Value_NotInLookup[wrapper.GetValue()]; ok {
			return WrappersValidationError{
				Field:  "Int32Value",
				Reason: "value must not be in list [4 5 6]",
			}
		}

	}

	if wrapper := m.GetUint32Value(); wrapper != nil {

		if wrapper.GetValue() <= 3 {
			return WrappersValidationError{
				Field:  "Uint32Value",
				Reason: "value must be greater than 3",
			}
		}

	}

	if wrapper := m.GetBoolValue(); wrapper != nil {

		if wrapper.GetValue() != true {
			return WrappersValidationError{
				Field:  "BoolValue",
				Reason: "value must equal true",
			}
		}

	}

	if wrapper := m.GetStringValue(); wrapper != nil {

		if !strings.HasPrefix(wrapper.GetValue(), "foo") {
			return WrappersValidationError{
				Field:  "StringValue",
				Reason: "value does not have prefix \"foo\"",
			}
		}

	}

	if wrapper := m.GetBytesValue(); wrapper != nil {

		if len(wrapper.GetValue()) > 8 {
			return WrappersValidationError{
				Field:  "BytesValue",
				Reason: "value length must be at most 8 bytes",
			}
		}

	}

	return nil
}

// WrappersValidationError is the validation error returned by
// Wrappers.Validate if the designated constraints aren't met.
type WrappersValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e WrappersValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWrappers.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = WrappersValidationError{}

var _Wrappers_Uint64Value_InLookup = map[uint64]struct{}{
	1: {},
	2: {},
	3: {},
}

var _Wrappers_Int32Value_NotInLookup = map[int32]struct{}{
	4: {},
	5: {},
	6: {},
}
