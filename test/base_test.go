package test

import (
	"RestN/normalizer"
	"RestN/rest"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func createResponse() rest.RestResponse {
	return rest.RestResponse{
		Status: 5,
		Locale: rest.Locale{Code: "en", Name: "english"},
		Error: rest.RestError{
			Validation: rest.Validation{
				FieldErrors: []rest.FieldValidationError{
					{
						Field:   "test1",
						Message: "msg1",
					},
					{
						Field:   "test2",
						Message: "msg2",
					},
					{
						Field:   "test3",
						Message: "msg3",
					},
				},
			},
			Message: []string{
				"string1",
				"string2",
				"string3",
			},
		},
		Body: nil,
	}
}

func getResponseAsMap() map[string]interface{} {
	resp := createResponse()
	m, _ := resp.NormalizeResponse().(map[string]interface{})

	return m
}

func TestIsMaps(t *testing.T) {
	normalizer.Init()
	resp := createResponse()

	result, ok := resp.NormalizeResponse().(map[string]interface{})

	// is map
	require.True(t, ok)

	// status not map
	_, ok = result["Status"].(map[string]interface{})
	require.True(t, !ok)

	// locale is map
	_, ok = result["Locale"].(map[string]interface{})
	require.True(t, ok)

	// error is map
	_, ok = result["Error"].(map[string]interface{})
	require.True(t, ok)

	// body is map
	require.Equal(t, nil, result["Body"])
}

func TestLocaleContent(t *testing.T) {
	normalizer.Init()
	locale, ok := getResponseAsMap()["Locale"].(map[string]interface{})

	// locale content
	require.True(t, ok)
	require.Equal(t, "en", locale["Code"])
	require.Equal(t, "english", locale["Name"])
}

func TestStatus(t *testing.T) {
	normalizer.Init()
	resp := getResponseAsMap()

	require.Equal(t, "5", resp["Status"])

	normalizer.Reset()
}

func TestErrorContent(t *testing.T) {
	normalizer.Init()

	errorMap, ok := getResponseAsMap()["Error"].(map[string]interface{})
	require.True(t, ok)
	require.NotEmpty(t, errorMap["Validation"])
	require.NotEmpty(t, errorMap["Message"])

	messageMap, ok := errorMap["Message"].([]string)
	require.True(t, ok)
	require.Len(t, messageMap, 3)
	require.Equal(t, "string1", messageMap[0])
	require.Equal(t, "string2", messageMap[1])
	require.Equal(t, "string3", messageMap[2])

	validationMap, ok := errorMap["Validation"].(map[string]interface{})
	require.True(t, ok)
	require.NotEmpty(t, validationMap["FieldErrors"])

	fieldErrors, ok := validationMap["FieldErrors"].([]interface{})
	require.True(t, ok)
	require.Len(t, fieldErrors, 3)

	error1, ok := reflect.ValueOf(fieldErrors).Index(0).Interface().(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, "test1", error1["Field"])
	require.Equal(t, "msg1", error1["Message"])

	error2, ok := reflect.ValueOf(fieldErrors).Index(1).Interface().(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, "test2", error2["Field"])
	require.Equal(t, "msg2", error2["Message"])

	error3, ok := reflect.ValueOf(fieldErrors).Index(2).Interface().(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, "test3", error3["Field"])
	require.Equal(t, "msg3", error3["Message"])
}

func TestDepth(t *testing.T) {
	normalizer.Init()
	r := createResponse()
	r.SetDepth(1)

	result, ok := r.NormalizeResponse().(map[string]interface{})
	require.True(t, ok)
	require.NotEmpty(t, result["Status"])
	require.Equal(t, "5", result["Status"])
	require.Empty(t, result["Locale"])
	require.Empty(t, result["Error"])
	require.Empty(t, result["Body"])
}
