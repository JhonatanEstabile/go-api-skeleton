package utils

import (
	"reflect"
	"strings"
)

var validationMessages = map[string]string{
	"alpha":                "The :attribute may only contain letters.",
	"alphanum":             "The :attribute may only contain letters and numbers.",
	"alphanumunicode":      "The :attribute may only contain letters, numbers and unicode characters.",
	"alphaunicode":         "The :attribute may only contain letters and unicode characters.",
	"boolean":              "The :attribute field must be true or false.",
	"datetime":             "The :attribute is not a valid datetime.",
	"email":                "The :attribute must be a valid email address.",
	"file":                 "The :attribute must be a file.",
	"gt":                   "The :attribute must be greater than :value.",
	"gte":                  "The :attribute must be greater than or equal :value.",
	"integer":              "The :attribute must be an integer.",
	"ip":                   "The :attribute must be a valid IP address.",
	"ipv4":                 "The :attribute must be a valid IPv4 address.",
	"ipv6":                 "The :attribute must be a valid IPv6 address.",
	"json":                 "The :attribute must be a valid JSON string.",
	"lt":                   "The :attribute must be less than :value.",
	"lte":                  "The :attribute must be less than or equal :value.",
	"max":                  "The :attribute may not be greater than :max.",
	"min":                  "The :attribute must be at least :min.",
	"numeric":              "The :attribute must be a number.",
	"regex":                "The :attribute format is invalid.",
	"required":             "The :attribute field is required.",
	"required_if":          "The :attribute field is required when :other is :value.",
	"required_unless":      "The :attribute field is required unless :other is in :values.",
	"required_with":        "The :attribute field is required when :values is present.",
	"required_with_all":    "The :attribute field is required when :values is present.",
	"required_without":     "The :attribute field is required when :values is not present.",
	"required_without_all": "The :attribute field is required when none of :values are present.",
	"string":               "The :attribute must be a string.",
	"timezone":             "The :attribute must be a valid zone.",
	"unique":               "The :attribute has already been taken.",
	"url":                  "The :attribute format is invalid.",
	"ulid":                 "The :attribute must be a valid ULID string.",
}

var ValUtil *ValidatorUtil

type ValidatorUtil struct {
	validate           *validator.Validate
	validationMessages map[string]string
}

func NewValidatorUtil() *ValidatorUtil {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return &ValidatorUtil{
		validate:           validate,
		validationMessages: validationMessages,
	}
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func (val *ValidatorUtil) ValidateStruct(modelData any) []string {
	var errors []string

	err := val.validate.Struct(modelData)
	if err != nil {
		var message string
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()

			message = val.translateError(&element)
			errors = append(errors, message)
		}
	}

	return errors
}

func (val *ValidatorUtil) translateError(errorData *ErrorResponse) string {
	var message string = "The :attribute with :values is not valid."

	if val, ok := val.validationMessages[errorData.Tag]; ok {
		message = val
	}

	message = strings.Replace(message, ":attribute", errorData.FailedField, 1)
	message = strings.Replace(message, ":values", errorData.Value, 1)
	return message
}
