// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/rec/rec.proto

package rec

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on CreateRecipeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRecipeRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return CreateRecipeRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		return CreateRecipeRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetIngredients()) < 1 {
		return CreateRecipeRequestValidationError{
			field:  "Ingredients",
			reason: "value must contain at least 1 item(s)",
		}
	}

	_CreateRecipeRequest_Ingredients_Unique := make(map[string]struct{}, len(m.GetIngredients()))

	for idx, item := range m.GetIngredients() {
		_, _ = idx, item

		if _, exists := _CreateRecipeRequest_Ingredients_Unique[item]; exists {
			return CreateRecipeRequestValidationError{
				field:  fmt.Sprintf("Ingredients[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_CreateRecipeRequest_Ingredients_Unique[item] = struct{}{}
		}

		// no validation rules for Ingredients[idx]
	}

	if len(m.GetSteps()) < 1 {
		return CreateRecipeRequestValidationError{
			field:  "Steps",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetSteps() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateRecipeRequestValidationError{
					field:  fmt.Sprintf("Steps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateRecipeRequestValidationError is the validation error returned by
// CreateRecipeRequest.Validate if the designated constraints aren't met.
type CreateRecipeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRecipeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRecipeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRecipeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRecipeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRecipeRequestValidationError) ErrorName() string {
	return "CreateRecipeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRecipeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRecipeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRecipeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRecipeRequestValidationError{}

// Validate checks the field values on CreateRecipeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRecipeResponse) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return CreateRecipeResponseValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// CreateRecipeResponseValidationError is the validation error returned by
// CreateRecipeResponse.Validate if the designated constraints aren't met.
type CreateRecipeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRecipeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRecipeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRecipeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRecipeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRecipeResponseValidationError) ErrorName() string {
	return "CreateRecipeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRecipeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRecipeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRecipeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRecipeResponseValidationError{}

// Validate checks the field values on UpdateRecipeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRecipeRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return UpdateRecipeRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		return UpdateRecipeRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetIngredients()) < 1 {
		return UpdateRecipeRequestValidationError{
			field:  "Ingredients",
			reason: "value must contain at least 1 item(s)",
		}
	}

	_UpdateRecipeRequest_Ingredients_Unique := make(map[string]struct{}, len(m.GetIngredients()))

	for idx, item := range m.GetIngredients() {
		_, _ = idx, item

		if _, exists := _UpdateRecipeRequest_Ingredients_Unique[item]; exists {
			return UpdateRecipeRequestValidationError{
				field:  fmt.Sprintf("Ingredients[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_UpdateRecipeRequest_Ingredients_Unique[item] = struct{}{}
		}

		// no validation rules for Ingredients[idx]
	}

	if len(m.GetSteps()) < 1 {
		return UpdateRecipeRequestValidationError{
			field:  "Steps",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetSteps() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateRecipeRequestValidationError{
					field:  fmt.Sprintf("Steps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetId() < 1 {
		return UpdateRecipeRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// UpdateRecipeRequestValidationError is the validation error returned by
// UpdateRecipeRequest.Validate if the designated constraints aren't met.
type UpdateRecipeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRecipeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRecipeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRecipeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRecipeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRecipeRequestValidationError) ErrorName() string {
	return "UpdateRecipeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRecipeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRecipeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRecipeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRecipeRequestValidationError{}

// Validate checks the field values on UpdateRecipeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRecipeResponse) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return UpdateRecipeResponseValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// UpdateRecipeResponseValidationError is the validation error returned by
// UpdateRecipeResponse.Validate if the designated constraints aren't met.
type UpdateRecipeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRecipeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRecipeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRecipeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRecipeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRecipeResponseValidationError) ErrorName() string {
	return "UpdateRecipeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRecipeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRecipeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRecipeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRecipeResponseValidationError{}

// Validate checks the field values on DeleteRecipeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteRecipeRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return DeleteRecipeRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// DeleteRecipeRequestValidationError is the validation error returned by
// DeleteRecipeRequest.Validate if the designated constraints aren't met.
type DeleteRecipeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRecipeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRecipeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRecipeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRecipeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRecipeRequestValidationError) ErrorName() string {
	return "DeleteRecipeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRecipeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRecipeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRecipeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRecipeRequestValidationError{}

// Validate checks the field values on DeleteRecipeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteRecipeResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteRecipeResponseValidationError is the validation error returned by
// DeleteRecipeResponse.Validate if the designated constraints aren't met.
type DeleteRecipeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRecipeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRecipeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRecipeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRecipeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRecipeResponseValidationError) ErrorName() string {
	return "DeleteRecipeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRecipeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRecipeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRecipeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRecipeResponseValidationError{}

// Validate checks the field values on RegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RegisterRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetLogin()) < 1 {
		return RegisterRequestValidationError{
			field:  "Login",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetPassword()) < 1 {
		return RegisterRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// RegisterRequestValidationError is the validation error returned by
// RegisterRequest.Validate if the designated constraints aren't met.
type RegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRequestValidationError) ErrorName() string { return "RegisterRequestValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRequestValidationError{}

// Validate checks the field values on RegisterResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RegisterResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// RegisterResponseValidationError is the validation error returned by
// RegisterResponse.Validate if the designated constraints aren't met.
type RegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterResponseValidationError) ErrorName() string { return "RegisterResponseValidationError" }

// Error satisfies the builtin error interface
func (e RegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterResponseValidationError{}

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LoginRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetLogin()) < 1 {
		return LoginRequestValidationError{
			field:  "Login",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetPassword()) < 1 {
		return LoginRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LoginResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// LoginResponseValidationError is the validation error returned by
// LoginResponse.Validate if the designated constraints aren't met.
type LoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponseValidationError) ErrorName() string { return "LoginResponseValidationError" }

// Error satisfies the builtin error interface
func (e LoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponseValidationError{}

// Validate checks the field values on GetAllRecipesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetAllRecipesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MinTime

	// no validation rules for MaxTime

	// no validation rules for TimeOrder

	return nil
}

// GetAllRecipesRequestValidationError is the validation error returned by
// GetAllRecipesRequest.Validate if the designated constraints aren't met.
type GetAllRecipesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAllRecipesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAllRecipesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAllRecipesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAllRecipesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAllRecipesRequestValidationError) ErrorName() string {
	return "GetAllRecipesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAllRecipesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAllRecipesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAllRecipesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAllRecipesRequestValidationError{}

// Validate checks the field values on GetAllRecipesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetAllRecipesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRecipes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAllRecipesResponseValidationError{
					field:  fmt.Sprintf("Recipes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetAllRecipesResponseValidationError is the validation error returned by
// GetAllRecipesResponse.Validate if the designated constraints aren't met.
type GetAllRecipesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAllRecipesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAllRecipesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAllRecipesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAllRecipesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAllRecipesResponseValidationError) ErrorName() string {
	return "GetAllRecipesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAllRecipesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAllRecipesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAllRecipesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAllRecipesResponseValidationError{}

// Validate checks the field values on GetAllRecipesByUserRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetAllRecipesByUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MinTime

	// no validation rules for MaxTime

	// no validation rules for TimeOrder

	return nil
}

// GetAllRecipesByUserRequestValidationError is the validation error returned
// by GetAllRecipesByUserRequest.Validate if the designated constraints aren't met.
type GetAllRecipesByUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAllRecipesByUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAllRecipesByUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAllRecipesByUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAllRecipesByUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAllRecipesByUserRequestValidationError) ErrorName() string {
	return "GetAllRecipesByUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAllRecipesByUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAllRecipesByUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAllRecipesByUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAllRecipesByUserRequestValidationError{}

// Validate checks the field values on GetAllRecipesByUserResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetAllRecipesByUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRecipes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAllRecipesByUserResponseValidationError{
					field:  fmt.Sprintf("Recipes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetAllRecipesByUserResponseValidationError is the validation error returned
// by GetAllRecipesByUserResponse.Validate if the designated constraints
// aren't met.
type GetAllRecipesByUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAllRecipesByUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAllRecipesByUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAllRecipesByUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAllRecipesByUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAllRecipesByUserResponseValidationError) ErrorName() string {
	return "GetAllRecipesByUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAllRecipesByUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAllRecipesByUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAllRecipesByUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAllRecipesByUserResponseValidationError{}

// Validate checks the field values on GetRecipeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetRecipeRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return GetRecipeRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// GetRecipeRequestValidationError is the validation error returned by
// GetRecipeRequest.Validate if the designated constraints aren't met.
type GetRecipeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRecipeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRecipeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRecipeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRecipeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRecipeRequestValidationError) ErrorName() string { return "GetRecipeRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRecipeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRecipeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRecipeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRecipeRequestValidationError{}

// Validate checks the field values on GetRecipeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetRecipeResponse) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return GetRecipeResponseValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return GetRecipeResponseValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		return GetRecipeResponseValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetIngredients()) < 1 {
		return GetRecipeResponseValidationError{
			field:  "Ingredients",
			reason: "value must contain at least 1 item(s)",
		}
	}

	_GetRecipeResponse_Ingredients_Unique := make(map[string]struct{}, len(m.GetIngredients()))

	for idx, item := range m.GetIngredients() {
		_, _ = idx, item

		if _, exists := _GetRecipeResponse_Ingredients_Unique[item]; exists {
			return GetRecipeResponseValidationError{
				field:  fmt.Sprintf("Ingredients[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_GetRecipeResponse_Ingredients_Unique[item] = struct{}{}
		}

		// no validation rules for Ingredients[idx]
	}

	if len(m.GetSteps()) < 1 {
		return GetRecipeResponseValidationError{
			field:  "Steps",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetSteps() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetRecipeResponseValidationError{
					field:  fmt.Sprintf("Steps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetTotalTime() < 1 {
		return GetRecipeResponseValidationError{
			field:  "TotalTime",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// GetRecipeResponseValidationError is the validation error returned by
// GetRecipeResponse.Validate if the designated constraints aren't met.
type GetRecipeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRecipeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRecipeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRecipeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRecipeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRecipeResponseValidationError) ErrorName() string {
	return "GetRecipeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetRecipeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRecipeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRecipeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRecipeResponseValidationError{}

// Validate checks the field values on Step with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Step) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		return StepValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetDuration() < 1 {
		return StepValidationError{
			field:  "Duration",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// StepValidationError is the validation error returned by Step.Validate if the
// designated constraints aren't met.
type StepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StepValidationError) ErrorName() string { return "StepValidationError" }

// Error satisfies the builtin error interface
func (e StepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StepValidationError{}
