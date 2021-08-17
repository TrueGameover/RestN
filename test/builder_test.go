package test

import (
	"RestN/rest"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOptions(t *testing.T) {
	r := rest.RestResponse{}

	r.SetNormalizationOption("test1", "val1")
	r.SetNormalizationOption("test2", "val2")
	require.Equal(t, "val1", r.GetNormalizationOption("test1"))
	require.Equal(t, "val2", r.GetNormalizationOption("test2"))

	r.SetNormalizationOption("test1", "val5")
	require.Equal(t, "val5", r.GetNormalizationOption("test1"))

	r.RemoveNormalizationOption("test1")
	require.Empty(t, r.GetNormalizationOption("test1"))

	r.SetNormalizationOptionIf(true, "test1", "val1")
	require.Equal(t, "val1", r.GetNormalizationOption("test1"))

	r.RemoveNormalizationOption("test1")
	r.SetNormalizationOptionIf(false, "test1", "val1")
	require.Empty(t, r.GetNormalizationOption("test1"))
}

func TestSetters(t *testing.T) {
	r := rest.RestResponse{}

	r.SetDepth(5)
	require.Equal(t, 5, r.GetDepth())

	r.SetLocale(rest.Locale{
		Code: "test1",
		Name: "val1",
	})
	require.Equal(t, "test1", r.Locale.Code)
	require.Equal(t, "val1", r.Locale.Name)

	r.SetError(rest.RestError{
		Validation: rest.Validation{
			FieldErrors: []rest.FieldValidationError{
				{
					Field:   "test1",
					Message: "val1",
				},
			},
		},
		Message: []string{
			"str1",
			"str2",
		},
	})
	require.NotEmpty(t, r.Error)
	require.Equal(t, "test1", r.Error.Validation.FieldErrors[0].Field)
	require.Equal(t, "val1", r.Error.Validation.FieldErrors[0].Message)
	require.Equal(t, "str1", r.Error.Message[0])
	require.Equal(t, "str2", r.Error.Message[1])

	r.SetBody("test1")
	require.Equal(t, "test1", r.Body)

	r.SetStatus(15)
	require.Equal(t, 15, r.Status)
}
