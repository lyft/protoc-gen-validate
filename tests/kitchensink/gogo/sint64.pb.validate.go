// Code generated by protoc-gen-validate
// source: sint64.proto
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

// Validate checks the field values on SInt64 with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SInt64) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for None

	if m.GetLt() >= 5 {
		return SInt64ValidationError{
			Field:  "Lt",
			Reason: "value must be less than 5",
		}
	}

	if m.GetLte() > 5 {
		return SInt64ValidationError{
			Field:  "Lte",
			Reason: "value must be less than or equal to 5",
		}
	}

	if m.GetGt() <= 5 {
		return SInt64ValidationError{
			Field:  "Gt",
			Reason: "value must be greater than 5",
		}
	}

	if m.GetGte() < 5 {
		return SInt64ValidationError{
			Field:  "Gte",
			Reason: "value must be greater than or equal to 5",
		}
	}

	if val := m.GetLtGt(); val <= 10 || val >= 15 {
		return SInt64ValidationError{
			Field:  "LtGt",
			Reason: "value must be inside range (10, 15)",
		}
	}

	if val := m.GetLtGte(); val < 10 || val >= 15 {
		return SInt64ValidationError{
			Field:  "LtGte",
			Reason: "value must be inside range [10, 15)",
		}
	}

	if val := m.GetLteGt(); val <= 10 || val > 15 {
		return SInt64ValidationError{
			Field:  "LteGt",
			Reason: "value must be inside range (10, 15]",
		}
	}

	if val := m.GetLteGte(); val < 10 || val > 15 {
		return SInt64ValidationError{
			Field:  "LteGte",
			Reason: "value must be inside range [10, 15]",
		}
	}

	if val := m.GetLtGtInv(); val >= 20 && val <= 25 {
		return SInt64ValidationError{
			Field:  "LtGtInv",
			Reason: "value must be outside range [20, 25]",
		}
	}

	if val := m.GetLtGteInv(); val >= 20 && val < 25 {
		return SInt64ValidationError{
			Field:  "LtGteInv",
			Reason: "value must be outside range [20, 25)",
		}
	}

	if val := m.GetLteGtInv(); val > 20 && val <= 25 {
		return SInt64ValidationError{
			Field:  "LteGtInv",
			Reason: "value must be outside range (20, 25]",
		}
	}

	if val := m.GetLteGteInv(); val > 20 && val < 25 {
		return SInt64ValidationError{
			Field:  "LteGteInv",
			Reason: "value must be outside range (20, 25)",
		}
	}

	if _, ok := _SInt64_In_InLookup[m.GetIn()]; !ok {
		return SInt64ValidationError{
			Field:  "In",
			Reason: "value must be in list [30 35]",
		}
	}

	if _, ok := _SInt64_NotIn_NotInLookup[m.GetNotIn()]; ok {
		return SInt64ValidationError{
			Field:  "NotIn",
			Reason: "value must not be in list [40 45]",
		}
	}

	if m.GetConst() != 50 {
		return SInt64ValidationError{
			Field:  "Const",
			Reason: "value must equal 50",
		}
	}

	return nil
}

// SInt64ValidationError is the validation error returned by SInt64.Validate if
// the designated constraints aren't met.
type SInt64ValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e SInt64ValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSInt64.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = SInt64ValidationError{}

var _SInt64_In_InLookup = map[int64]struct{}{
	30: {},
	35: {},
}

var _SInt64_NotIn_NotInLookup = map[int64]struct{}{
	40: {},
	45: {},
}
