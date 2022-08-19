package api

type FailedResponse struct {
	Code   string          `json:"code,omitempty"`
	Errors []ResponseError `json:"errors,omitempty"`
	Data   string          `json:"data,omitempty"`
}

type ResponseError struct {
	Code    string `json:"code,omitempty"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}

const (
	InternalError = "internal-error"
	Unauthorized  = "unauthorized"
)
