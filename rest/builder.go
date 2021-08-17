package rest

type Options map[string]string

type RestResponse struct {
	Status  int
	Locale  Locale
	Error   RestError
	Body    interface{}
	options Options
	depth   int
}

type Locale struct {
	Code string
	Name string
}

type RestError struct {
	Validation Validation
	Message    []string
}

type Validation struct {
	FieldErrors []FieldValidationError
}

type FieldValidationError struct {
	Field   string
	Message string
}

func (response *RestResponse) SetStatus(status int) *RestResponse {
	response.Status = status
	return response
}

func (response *RestResponse) SetLocale(locale Locale) *RestResponse {
	response.Locale = locale
	return response
}

func (response *RestResponse) SetError(error RestError) *RestResponse {
	response.Error = error
	return response
}

func (response *RestResponse) SetBody(body interface{}) *RestResponse {
	response.Body = body
	return response
}

func (response *RestResponse) SetNormalizationOption(key string, value string) *RestResponse {
	response.checkOptionsMap()

	response.options[key] = value
	return response
}

func (response *RestResponse) SetNormalizationOptionIf(expr bool, key string, value string) *RestResponse {
	response.checkOptionsMap()

	if expr {
		response.options[key] = value
	}
	return response
}

func (response *RestResponse) GetNormalizationOption(key string) string {
	response.checkOptionsMap()

	return response.options[key]
}

func (response *RestResponse) RemoveNormalizationOption(key string) {
	response.checkOptionsMap()

	delete(response.options, key)
}

func (response *RestResponse) SetDepth(depth int) *RestResponse {
	response.depth = depth
	return response
}

func (response *RestResponse) GetDepth() int {
	return response.depth
}

func (response *RestResponse) checkOptionsMap() {
	if response.options == nil {
		response.options = map[string]string{}
	}
}
