package models

type ApiError struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload"`
}
