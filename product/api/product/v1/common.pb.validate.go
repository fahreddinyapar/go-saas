// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: product/api/product/v1/common.proto

package v1

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

// Validate checks the field values on Media with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Media) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Media with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MediaMultiError, or nil if none found.
func (m *Media) ValidateAll() error {
	return m.validate(true)
}

func (m *Media) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Type

	// no validation rules for Title

	if len(errors) > 0 {
		return MediaMultiError(errors)
	}

	return nil
}

// MediaMultiError is an error wrapping multiple validation errors returned by
// Media.ValidateAll() if the designated constraints aren't met.
type MediaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MediaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MediaMultiError) AllErrors() []error { return m }

// MediaValidationError is the validation error returned by Media.Validate if
// the designated constraints aren't met.
type MediaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MediaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MediaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MediaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MediaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MediaValidationError) ErrorName() string { return "MediaValidationError" }

// Error satisfies the builtin error interface
func (e MediaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMedia.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MediaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MediaValidationError{}

// Validate checks the field values on Badge with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Badge) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Badge with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BadgeMultiError, or nil if none found.
func (m *Badge) ValidateAll() error {
	return m.validate(true)
}

func (m *Badge) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Code

	// no validation rules for Label

	if len(errors) > 0 {
		return BadgeMultiError(errors)
	}

	return nil
}

// BadgeMultiError is an error wrapping multiple validation errors returned by
// Badge.ValidateAll() if the designated constraints aren't met.
type BadgeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BadgeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BadgeMultiError) AllErrors() []error { return m }

// BadgeValidationError is the validation error returned by Badge.Validate if
// the designated constraints aren't met.
type BadgeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BadgeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BadgeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BadgeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BadgeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BadgeValidationError) ErrorName() string { return "BadgeValidationError" }

// Error satisfies the builtin error interface
func (e BadgeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBadge.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BadgeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BadgeValidationError{}

// Validate checks the field values on Keyword with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Keyword) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Keyword with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in KeywordMultiError, or nil if none found.
func (m *Keyword) ValidateAll() error {
	return m.validate(true)
}

func (m *Keyword) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for Refer

	if len(errors) > 0 {
		return KeywordMultiError(errors)
	}

	return nil
}

// KeywordMultiError is an error wrapping multiple validation errors returned
// by Keyword.ValidateAll() if the designated constraints aren't met.
type KeywordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KeywordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KeywordMultiError) AllErrors() []error { return m }

// KeywordValidationError is the validation error returned by Keyword.Validate
// if the designated constraints aren't met.
type KeywordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeywordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeywordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeywordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeywordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeywordValidationError) ErrorName() string { return "KeywordValidationError" }

// Error satisfies the builtin error interface
func (e KeywordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyword.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeywordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeywordValidationError{}

// Validate checks the field values on CampaignRule with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CampaignRule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CampaignRule with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CampaignRuleMultiError, or
// nil if none found.
func (m *CampaignRule) ValidateAll() error {
	return m.validate(true)
}

func (m *CampaignRule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Rule

	if all {
		switch v := interface{}(m.GetExtra()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CampaignRuleValidationError{
					field:  "Extra",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CampaignRuleValidationError{
					field:  "Extra",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExtra()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CampaignRuleValidationError{
				field:  "Extra",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CampaignRuleMultiError(errors)
	}

	return nil
}

// CampaignRuleMultiError is an error wrapping multiple validation errors
// returned by CampaignRule.ValidateAll() if the designated constraints aren't met.
type CampaignRuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CampaignRuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CampaignRuleMultiError) AllErrors() []error { return m }

// CampaignRuleValidationError is the validation error returned by
// CampaignRule.Validate if the designated constraints aren't met.
type CampaignRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CampaignRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CampaignRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CampaignRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CampaignRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CampaignRuleValidationError) ErrorName() string { return "CampaignRuleValidationError" }

// Error satisfies the builtin error interface
func (e CampaignRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCampaignRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CampaignRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CampaignRuleValidationError{}

// Validate checks the field values on Stock with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Stock) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Stock with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StockMultiError, or nil if none found.
func (m *Stock) ValidateAll() error {
	return m.validate(true)
}

func (m *Stock) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for InStock

	if _, ok := _Stock_Level_InLookup[m.GetLevel()]; !ok {
		err := StockValidationError{
			field:  "Level",
			reason: "value must be in list [out in low]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Amount

	// no validation rules for DeliveryCode

	if len(errors) > 0 {
		return StockMultiError(errors)
	}

	return nil
}

// StockMultiError is an error wrapping multiple validation errors returned by
// Stock.ValidateAll() if the designated constraints aren't met.
type StockMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StockMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StockMultiError) AllErrors() []error { return m }

// StockValidationError is the validation error returned by Stock.Validate if
// the designated constraints aren't met.
type StockValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StockValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StockValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StockValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StockValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StockValidationError) ErrorName() string { return "StockValidationError" }

// Error satisfies the builtin error interface
func (e StockValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStock.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StockValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StockValidationError{}

var _Stock_Level_InLookup = map[string]struct{}{
	"out": {},
	"in":  {},
	"low": {},
}

// Validate checks the field values on Price with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Price) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Price with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PriceMultiError, or nil if none found.
func (m *Price) ValidateAll() error {
	return m.validate(true)
}

func (m *Price) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CurrencyCode

	// no validation rules for DefaultAmount

	// no validation rules for DiscountedAmount

	// no validation rules for DiscountText

	// no validation rules for DenyMoreDiscounts

	if len(errors) > 0 {
		return PriceMultiError(errors)
	}

	return nil
}

// PriceMultiError is an error wrapping multiple validation errors returned by
// Price.ValidateAll() if the designated constraints aren't met.
type PriceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PriceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PriceMultiError) AllErrors() []error { return m }

// PriceValidationError is the validation error returned by Price.Validate if
// the designated constraints aren't met.
type PriceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PriceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PriceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PriceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PriceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PriceValidationError) ErrorName() string { return "PriceValidationError" }

// Error satisfies the builtin error interface
func (e PriceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrice.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PriceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PriceValidationError{}