package normalizer

import "RestN/rest"

func Init() {
	Reset()

	rest.RegisterNormalizer(RestResponseNormalizer{})
	rest.RegisterNormalizer(LocaleNormalizer{})
	rest.RegisterNormalizer(RestErrorNormalizer{})
	rest.RegisterNormalizer(FieldValidationErrorNormalizer{})
	rest.RegisterNormalizer(ValidationNormalizer{})
	rest.RegisterNormalizer(SliceNormalizer{})
}

func Reset() {
	rest.ClearAllNormalizers()
}
