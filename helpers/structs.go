package helpers

type ErrorsGeneral struct {
	Result  string          `json:"result"`
	Message string          `json:"errors_message,omitempty"`
	Errors  []ResponseError `json:"errors,omitempty"`
}

type ResponseError struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}
