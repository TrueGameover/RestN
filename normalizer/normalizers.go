package normalizer

import "github.com/TrueGameover/RestN/rest"

func Init() {
	Reset()

	rest.RegisterNormalizer(RestResponseNormalizer{})
	rest.RegisterNormalizer(LocaleNormalizer{})
	rest.RegisterNormalizer(RestErrorNormalizer{})
	rest.RegisterNormalizer(FieldValidationErrorNormalizer{})
	rest.RegisterNormalizer(ValidationNormalizer{})
	rest.RegisterNormalizer(SliceNormalizer{})
	rest.RegisterNormalizer(SyncMapNormalizer{})
}

func Reset() {
	rest.ClearAllNormalizers()
}
